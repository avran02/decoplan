package router

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/decplan/gateway/internal/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	chi.Router
	uc controllers.UsersController
	cc controllers.ChatsController
}

func (router *Router) getUsersRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/users", func(r chi.Router) {
		r.Post("/", router.uc.CreateUserHandler)
		r.Get("/{id}", router.uc.GetUserHandler)
		r.Put("/{id}", router.uc.UpdateUserHandler)
		r.Delete("/{id}", router.uc.DeleteUserHandler)
	})

	r.Route("/chats", func(r chi.Router) {
		r.Post("/", router.cc.CreateChatHandler)
		r.Get("/{id}", router.cc.GetChatHandler)
		r.Delete("/{id}", router.cc.DeleteChatHandler)

		r.Route("/{chatID}/users", func(r chi.Router) {
			r.Post("/", router.cc.AddUserToChatHandler)
			r.Delete("/", router.cc.RemoveUserFromChatHandler)
		})
	})

	return r
}

func New(controller *controllers.Controller) Router {
	r := Router{
		cc: controller.ChatsController(),
		uc: controller.UsersController(),
	}
	main := chi.NewRouter()
	main.Use(middleware.Logger)
	main.Use(cors.Handler(allowAllCORS()))

	usersRoutes := r.getUsersRoutes()
	main.Mount("/api/v1", usersRoutes)
	r.Router = main
	printRoutes(r.Router)
	return r
}

func printRoutes(router chi.Routes) {
	slog.Debug("Serving routes:")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		loggingStr := fmt.Sprintf("Method: %s, Route: %s", method, route)
		slog.Debug(loggingStr)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Fatal(err)
	}
}
