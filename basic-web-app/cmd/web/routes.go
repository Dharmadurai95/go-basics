package main

import (
	"net/http"

	"github.com/Dharmadurai95/go-basic-web-application/pkg/config"
	"github.com/Dharmadurai95/go-basic-web-application/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	return mux
}
