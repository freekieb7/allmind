package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dataFile = "/app/var/allmind/db/data.db"

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Manager struct {
	DB *sql.DB
}

func NewManager() *Manager {
	db, err := sql.Open("sqlite3", dataFile)

	if err != nil {
		log.Fatal(err)
	}

	err = migrate(db)

	if err != nil {
		log.Fatal(err)
	}

	return &Manager{
		db,
	}
}

func migrate(db *sql.DB) error {
	content, err := os.ReadFile("/app/pkg/database/scripts/init.sql")

	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))

	return err
}
