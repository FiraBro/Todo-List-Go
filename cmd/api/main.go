package main

import (
	"Todo-List-Go/internal/env"
	"Todo-List-Go/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":4000"),
	}

	store := store.NewSQLStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
