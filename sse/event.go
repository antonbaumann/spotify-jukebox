package sse

import (
	"time"

	"github.com/antonbaumann/spotify-jukebox/user"

	"github.com/antonbaumann/spotify-jukebox/events"
	"github.com/antonbaumann/spotify-jukebox/song"
)

const (
	PlaylistChange         events.EventType = "sse:playlist_change"
	PlayerStateChange      events.EventType = "sse:player_state_change"
	UserListChange         events.EventType = "sse:user_list_change"
	UserSynchronizedChange events.EventType = "sse:user_synchronized_change"
)

type PlaylistChangePayload []*song.Model

type PlayerStateChangePayload struct {
	CurrentSong *song.Model `json:"current_song"`
	IsPlaying   bool        `json:"is_playing"`
	ProgressMs  int64       `json:"progress"`
	Timestamp   time.Time   `json:"timestamp"`
}

type UserListChangePayload []*user.Model

type UserSynchronizedChangePayload struct {
	UserID       string `json:"user_id"`
	Synchronized bool   `json:"synchronized"`
}
