package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "lukap",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "gin",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
