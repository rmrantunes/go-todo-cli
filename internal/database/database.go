package database

import (
	"database/sql"
	"log"
	"path/filepath"
	"todo-cli/util"

	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	DB *sql.DB
}

var file = filepath.Join("internal", "database", "main.db")

func New() *Service {
	log.Printf("Connecting to database...")
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		util.DieOnError(err)
	}

	return &Service{
		DB: db,
	}
}

func (s *Service) Close() error {
	log.Printf("Disconnected from database")
	return s.DB.Close()
}
