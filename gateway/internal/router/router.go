package router

import (
	"github.com/avran02/decplan/gateway/internal/controllers"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	chi.Router
	ac controllers.AuthController
}

func (router *Router) getAuthRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", router.ac.Register)
		r.Post("/login", router.ac.Login)
		r.Post("/refresh", router.ac.RefreshTokens)
		r.Post("/logout", router.ac.Logout)
	})

	return r
}

func New(controller controllers.AuthController) Router {
	r := Router{
		ac: controller,
	}
	main := chi.NewRouter()

	filesRouter := r.getAuthRoutes()
	main.Mount("/", filesRouter)
	r.Router = main
	return r
}
