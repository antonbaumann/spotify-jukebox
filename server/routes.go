package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Model) setupServerRoutes(r *mux.Router) {
	r.Handle(
		"/ping",
		http.HandlerFunc(s.ServerHandler.Ping),
	)
}

func (s *Model) setupUserRoutes(r *mux.Router, auth authFunc) {
	r.Handle(
		"/users/join/{username}",
		http.HandlerFunc(s.UserHandler.Join),
	).Methods(http.MethodGet)

	r.Handle(
		"/users/{username}/list",
		auth(http.HandlerFunc(s.UserHandler.ListUsers)),
	).Methods(http.MethodGet)

	r.Handle(
		"/users/{username}/suggest/{song_id}",
		auth(http.HandlerFunc(s.UserHandler.SuggestSong)),
	).Methods(http.MethodGet)

	r.Handle(
		"/users/{username}/listSongs",
		auth(http.HandlerFunc(s.UserHandler.ListSongs)),
	).Methods(http.MethodGet)

	r.Handle(
		"/users/{username}/vote/{song_id}/{vote_action}",
		auth(http.HandlerFunc(s.UserHandler.Vote)),
	).Methods(http.MethodGet)
}

func (s *Model) setupAdminRoutes(r *mux.Router, auth authFunc) {
	r.Handle(
		"/admin/login",
		http.HandlerFunc(s.AdminHandler.Login),
	).Methods(http.MethodPost)

	r.Handle(
		"/users/{username}/removeSong/{song_id}",
		auth(http.HandlerFunc(s.AdminHandler.RemoveSong)),
	).Methods(http.MethodGet)
}

func (s *Model) setupSpotifyRoutes(r *mux.Router) {
	r.Handle(
		"/callback",
		http.HandlerFunc(s.SpotifyHandler.Redirect),
	)
}

func (s *Model) setupEventRoutes(r *mux.Router) {
	r.Handle(
		"/events",
		http.HandlerFunc(s.Broker.ServeHTTP),
	)
}
