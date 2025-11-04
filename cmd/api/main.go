package main

import (
	"Todo-List-Go/internal/db"
	"context"
	"Todo-List-Go/internal/env"
	"Todo-List-Go/internal/store"
	"log"
	"time"
)

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr string
	db   dbConfig
}

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":4000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"), // string, not int
		},
	}

	dbConn, err := db.New(cfg.db.addr, cfg.db.maxIdleConns, cfg.db.maxIdleTime, cfg.db.maxOpenConns)
	if err != nil {
		log.Panic(err)
	}

	store := store.NewSQLStorage(dbConn)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
