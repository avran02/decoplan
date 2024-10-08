package router

import (
	"github.com/avran02/decoplan/chat/internal/hub"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	chi.Router
	hub hub.WebsocketHub
}

func (r *Router) getConnectionRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.HandleFunc("/connect", r.hub.RegisterWebsocket)
	router.HandleFunc("/disconnect", r.hub.CloseWebsocket)

	return router
}

func New(hub hub.WebsocketHub) *Router {
	r := &Router{
		hub: hub,
	}

	routes := r.getConnectionRoutes()
	r.Router = routes

	return r
}
