package controllers

import "github.com/dallaspadilla/silverline/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/pings", middlewares.SetMiddlewareJSON(s.CreatePing)).Methods("POST")
	s.Router.HandleFunc("/pings", middlewares.SetMiddlewareJSON(s.GetPings)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetPing)).Methods("GET")
}
