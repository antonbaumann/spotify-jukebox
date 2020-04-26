// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"

	user "github.com/antonbaumann/spotify-jukebox/user"
)

// UserCollection is an autogenerated mock type for the UserCollection type
type UserCollection struct {
	mock.Mock
}

// AddSSEConnection provides a mock function with given fields: ctx, userID
func (_m *UserCollection) AddSSEConnection(ctx context.Context, userID string) (int, error) {
	ret := _m.Called(ctx, userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddUser provides a mock function with given fields: ctx, newUser
func (_m *UserCollection) AddUser(ctx context.Context, newUser *user.Model) error {
	ret := _m.Called(ctx, newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.Model) error); ok {
		r0 = rf(ctx, newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, userID
func (_m *UserCollection) DeleteUser(ctx context.Context, userID string) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUsersBySessionID provides a mock function with given fields: ctx, sessionID
func (_m *UserCollection) DeleteUsersBySessionID(ctx context.Context, sessionID string) error {
	ret := _m.Called(ctx, sessionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdminBySessionID provides a mock function with given fields: ctx, sessionID
func (_m *UserCollection) GetAdminBySessionID(ctx context.Context, sessionID string) (*user.Model, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 *user.Model
	if rf, ok := ret.Get(0).(func(context.Context, string) *user.Model); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Model)
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

// GetSpotifyClients provides a mock function with given fields: ctx, sessionID
func (_m *UserCollection) GetSpotifyClients(ctx context.Context, sessionID string) ([]*user.SpotifyClient, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 []*user.SpotifyClient
	if rf, ok := ret.Get(0).(func(context.Context, string) []*user.SpotifyClient); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*user.SpotifyClient)
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

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *UserCollection) GetUserByID(ctx context.Context, userID string) (*user.Model, error) {
	ret := _m.Called(ctx, userID)

	var r0 *user.Model
	if rf, ok := ret.Get(0).(func(context.Context, string) *user.Model); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByState provides a mock function with given fields: ctx, state
func (_m *UserCollection) GetUserByState(ctx context.Context, state string) (*user.Model, error) {
	ret := _m.Called(ctx, state)

	var r0 *user.Model
	if rf, ok := ret.Get(0).(func(context.Context, string) *user.Model); ok {
		r0 = rf(ctx, state)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, state)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementScore provides a mock function with given fields: ctx, username, amount
func (_m *UserCollection) IncrementScore(ctx context.Context, username string, amount int) error {
	ret := _m.Called(ctx, username, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, username, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListUsers provides a mock function with given fields: ctx, sessionID
func (_m *UserCollection) ListUsers(ctx context.Context, sessionID string) ([]*user.ListElement, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 []*user.ListElement
	if rf, ok := ret.Get(0).(func(context.Context, string) []*user.ListElement); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*user.ListElement)
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

// RemoveSSEConnection provides a mock function with given fields: ctx, userID
func (_m *UserCollection) RemoveSSEConnection(ctx context.Context, userID string) (int, error) {
	ret := _m.Called(ctx, userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSynchronized provides a mock function with given fields: ctx, userID, synchronized
func (_m *UserCollection) SetSynchronized(ctx context.Context, userID string, synchronized bool) (*user.Model, error) {
	ret := _m.Called(ctx, userID, synchronized)

	var r0 *user.Model
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *user.Model); ok {
		r0 = rf(ctx, userID, synchronized)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, userID, synchronized)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetToken provides a mock function with given fields: ctx, userID, token
func (_m *UserCollection) SetToken(ctx context.Context, userID string, token *oauth2.Token) error {
	ret := _m.Called(ctx, userID, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *oauth2.Token) error); ok {
		r0 = rf(ctx, userID, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
