package router

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/avran02/decoplan/files/internal/controller"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	chi.Router
	controller controller.FilesController
}

func (r Router) getFilesRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/upload", r.controller.UploadFile)
	router.Post("/download/{id}", r.controller.DownloadFile)
	router.Delete("/delete/{id}", r.controller.DeleteFile)

	return router
}

func New(controller controller.FilesController) Router {
	r := Router{
		controller: controller,
	}
	main := chi.NewRouter()

	filesRouter := r.getFilesRoutes()
	main.Mount("/", filesRouter)
	r.Router = main
	printRoutes(r.Router)
	return r
}

func printRoutes(router chi.Routes) {
	slog.Info("Routes:")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		loggingStr := fmt.Sprintf("Method: %s, Route: %s", method, route)
		slog.Info(loggingStr)
		return nil
	}
	err := chi.Walk(router, walkFunc)
	if err != nil {
		slog.Error(err.Error())
	}
}
