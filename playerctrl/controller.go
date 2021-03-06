package playerctrl

import (
	"context"
	"errors"
	"time"

	"github.com/encore-fm/backend/db"
	"github.com/encore-fm/backend/events"
	"github.com/encore-fm/backend/player"
	"github.com/encore-fm/backend/song"
	"github.com/encore-fm/backend/sse"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
)

var (
	ErrEventPayloadMalformed = errors.New("event payload malformed")
)

type Controller struct {
	sessionCollection db.SessionCollection
	songCollection    db.SongCollection
	userCollection    db.UserCollection
	playerCollection  db.PlayerCollection

	authenticator spotify.Authenticator

	eventBus events.EventBus

	// maps sessions to timers
	// timer fires when current song ended and new song must be fetched from db
	timers map[string]*time.Timer
}

func NewController(
	eventBus events.EventBus,
	sessionCollection db.SessionCollection,
	songCollection db.SongCollection,
	userCollection db.UserCollection,
	playerCollection db.PlayerCollection,
	authenticator spotify.Authenticator,
) *Controller {
	controller := &Controller{
		sessionCollection: sessionCollection,
		userCollection:    userCollection,
		songCollection:    songCollection,
		playerCollection:  playerCollection,
		authenticator:     authenticator,
		eventBus:          eventBus,
		timers:            make(map[string]*time.Timer),
	}
	return controller
}

func (ctrl *Controller) Start() error {
	msg := "[playerctrl] start"
	ctx := context.Background()
	sessionIDs, err := ctrl.sessionCollection.ListSessionIDs(ctx)
	if err != nil {
		return err
	}

	// initialize timer for every known sessionID
	for _, sessionID := range sessionIDs {
		playerState, err := ctrl.playerCollection.GetPlayer(ctx, sessionID)
		if err != nil {
			log.Errorf(msg, err)
		}
		if playerState == nil || playerState.IsEmpty() {
			ctrl.setTimer(sessionID, 0, func() { ctrl.getNextSong(sessionID) })
		} else {
			timerDuration := time.Duration(playerState.CurrentSong.Duration)*time.Millisecond - playerState.Progress()
			if timerDuration < 0 {
				timerDuration = 0
			}
			ctrl.setTimer(
				sessionID,
				timerDuration,
				func() { ctrl.getNextSong(sessionID) },
			)
			ctrl.notifyClientsBySessionID(
				sessionID,
				ctrl.setPlayerStateAction(playerState.CurrentSong.ID, playerState.Progress(), playerState.Paused),
			)
		}
	}

	go ctrl.eventLoop()

	return nil
}

func (ctrl *Controller) eventLoop() {
	songAdded := ctrl.eventBus.Subscribe([]events.EventType{SongAdded}, []events.GroupID{events.GroupIDAny})
	playPause := ctrl.eventBus.Subscribe([]events.EventType{PlayPauseEvent}, []events.GroupID{events.GroupIDAny})
	skip := ctrl.eventBus.Subscribe([]events.EventType{SkipEvent}, []events.GroupID{events.GroupIDAny})
	seek := ctrl.eventBus.Subscribe([]events.EventType{SeekEvent}, []events.GroupID{events.GroupIDAny})
	setSynchronized := ctrl.eventBus.Subscribe([]events.EventType{SetSynchronizedEvent}, []events.GroupID{events.GroupIDAny})
	sseConnection := ctrl.eventBus.Subscribe([]events.EventType{SSEConnectionEvent}, []events.GroupID{events.GroupIDAny})
	reset := ctrl.eventBus.Subscribe([]events.EventType{ResetEvent}, []events.GroupID{events.GroupIDAny})

	for {
		select {
		case ev := <-songAdded.Channel:
			ctrl.handleSongAdded(ev)

		case ev := <-playPause.Channel:
			ctrl.handlePlayPause(ev)

		case ev := <-skip.Channel:
			ctrl.handleSkip(ev)

		case ev := <-seek.Channel:
			ctrl.handleSeek(ev)

		case ev := <-setSynchronized.Channel:
			ctrl.handleSetSynchronized(ev)

		case ev := <-sseConnection.Channel:
			ctrl.handleSSEConnection(ev)

		case ev := <-reset.Channel:
			ctrl.handleReset(ev)
		}
	}
}

func (ctrl *Controller) setTimer(sessionID string, duration time.Duration, f func()) {
	t, ok := ctrl.timers[sessionID]
	if !ok {
		ctrl.timers[sessionID] = time.AfterFunc(duration, f)
	} else {
		t.Reset(duration)
	}
}

func (ctrl *Controller) stopTimer(sessionID string) {
	t, ok := ctrl.timers[sessionID]
	if ok {
		if !t.Stop() {
			<-t.C
		}
		delete(ctrl.timers, sessionID)
	}
}

// gets next song from db and deletes it
// sends event
func (ctrl *Controller) getNextSong(sessionID string) {
	msg := "[playerctrl] get next song from db"
	ctx := context.Background()
	songList, err := ctrl.songCollection.ListSongs(ctx, sessionID)
	if err != nil {
		// if error occurs while fetching list
		// log error and try again in 500ms
		log.Errorf("%v: %v", msg, err)
		ctrl.setTimer(
			sessionID,
			time.Duration(500)*time.Millisecond,
			func() { ctrl.getNextSong(sessionID) },
		)
		return
	}
	if len(songList) == 0 {
		// if songList is empty
		// reset player, log error and wait for songAdded
		err = ctrl.playerCollection.SetPlayer(ctx, sessionID, player.New())
		if err != nil {
			log.Errorf("%v: %v", msg, err)
		}
		// explicitly publish a skip event when playlist is empty, or else last song (in player) cannot get skipped
		ctrl.notifyClientsBySessionID(sessionID, ctrl.playerSkipAction())
		log.Warnf("%v: %v", msg, "songlist empty")
		return
	}

	nextSong := songList[0]

	// remove nextSong from db
	if err := ctrl.songCollection.RemoveSong(ctx, sessionID, nextSong.ID); err != nil {
		log.Errorf("%v: %v", msg, err)
	}

	ctrl.eventBus.Publish(sse.PlaylistChange, events.GroupID(sessionID), songList[1:])

	// fetch next song after song has ended
	ctrl.setTimer(
		sessionID,
		time.Duration(nextSong.Duration)*time.Millisecond,
		func() { ctrl.getNextSong(sessionID) },
	)

	// update session player
	newPlayer := &player.Player{
		CurrentSong: nextSong,
		SongStart:   time.Now(),
		PauseStart:  time.Now(),
		Paused:      false,
	}
	if err := ctrl.playerCollection.SetPlayer(ctx, sessionID, newPlayer); err != nil {
		log.Errorf("%v: %v", msg, err)
	}

	// send out a player state change event
	ctrl.notifyPlayerStateChange(sessionID)

	ctrl.notifyClientsBySessionID(
		sessionID,
		ctrl.setPlayerStateAction(
			newPlayer.CurrentSong.ID,
			0,
			false,
		),
	)
}

// sends out a player state change event with relevant data about the current player state
func (ctrl *Controller) notifyPlayerStateChange(sessionID string) {
	msg := "[playerctrl] notify player state change"
	ctx := context.Background()

	var currentSong *song.Model
	var isPlaying bool
	var progress int64

	playr, err := ctrl.playerCollection.GetPlayer(ctx, sessionID)
	if err != nil {
		log.Errorf("%v: %v", msg, err)
		return
	}

	if playr != nil {
		currentSong = playr.CurrentSong
		isPlaying = !playr.Paused
		progress = playr.Progress().Milliseconds()
	}

	payload := &sse.PlayerStateChangePayload{
		CurrentSong: currentSong,
		IsPlaying:   isPlaying,
		ProgressMs:  progress,
		Timestamp:   time.Now(),
	}

	ctrl.eventBus.Publish(
		sse.PlayerStateChange,
		events.GroupID(sessionID),
		payload,
	)
}
