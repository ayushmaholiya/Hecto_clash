package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const version = "0.0.1"

type dbConfig struct {
	addr 			string
	maxOpenConns 	int
	maxIdleConns 	int
	maxIdleTime 	string
}

type config struct {
	addr 	string
	db 		dbConfig
	env 	string
}

type application struct {
	config 	config
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Minute))

	return r
}