package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
type application struct{
	config config
}

type config struct{
	addr string
}
func(app *application) mount() * chi.Mux {
		r := chi.NewRouter()
		// Add middleware
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
		
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	return r
}
func (app *application) run(mux * chi.Mux) error{
	srv:=&http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
	}
	log.Printf("server is running %s",app.config.addr)
	return srv.ListenAndServe()
}