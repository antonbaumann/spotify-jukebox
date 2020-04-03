// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import session "github.com/antonbaumann/spotify-jukebox/session"
import song "github.com/antonbaumann/spotify-jukebox/song"

// SessionCollection is an autogenerated mock type for the SessionCollection type
type SessionCollection struct {
	mock.Mock
}

// AddSession provides a mock function with given fields: ctx, sess
func (_m *SessionCollection) AddSession(ctx context.Context, sess *session.Session) error {
	ret := _m.Called(ctx, sess)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *session.Session) error); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddSong provides a mock function with given fields: ctx, sessionID, newSong
func (_m *SessionCollection) AddSong(ctx context.Context, sessionID string, newSong *song.Model) error {
	ret := _m.Called(ctx, sessionID, newSong)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *song.Model) error); ok {
		r0 = rf(ctx, sessionID, newSong)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSessionByID provides a mock function with given fields: ctx, sessionID
func (_m *SessionCollection) GetSessionByID(ctx context.Context, sessionID string) (*session.Session, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 *session.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) *session.Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSongByID provides a mock function with given fields: ctx, sessionID, songID
func (_m *SessionCollection) GetSongByID(ctx context.Context, sessionID string, songID string) (*song.Model, error) {
	ret := _m.Called(ctx, sessionID, songID)

	var r0 *song.Model
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *song.Model); ok {
		r0 = rf(ctx, sessionID, songID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*song.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sessionID, songID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSongs provides a mock function with given fields: ctx, sessionID
func (_m *SessionCollection) ListSongs(ctx context.Context, sessionID string) ([]*song.Model, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 []*song.Model
	if rf, ok := ret.Get(0).(func(context.Context, string) []*song.Model); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*song.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveSong provides a mock function with given fields: ctx, sessionID, songID
func (_m *SessionCollection) RemoveSong(ctx context.Context, sessionID string, songID string) error {
	ret := _m.Called(ctx, sessionID, songID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, sessionID, songID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VoteDown provides a mock function with given fields: ctx, sessionID, songID, username
func (_m *SessionCollection) VoteDown(ctx context.Context, sessionID string, songID string, username string) (int, error) {
	ret := _m.Called(ctx, sessionID, songID, username)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) int); ok {
		r0 = rf(ctx, sessionID, songID, username)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, sessionID, songID, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VoteUp provides a mock function with given fields: ctx, sessionID, songID, username
func (_m *SessionCollection) VoteUp(ctx context.Context, sessionID string, songID string, username string) (int, error) {
	ret := _m.Called(ctx, sessionID, songID, username)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) int); ok {
		r0 = rf(ctx, sessionID, songID, username)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, sessionID, songID, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
