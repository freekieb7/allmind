package database

import (
	"errors"
)

const fileName = "db/sqlite.db"

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// type SQLiteRepository struct {
// 	db *sql.DB
// }

// func New() (*SQLiteRepository, error) {
// 	db, err := sql.Open("sqlite3", fileName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &SQLiteRepository{
// 		db: db,
// 	}, nil
// }

// func (r *SQLiteRepository) Migrate() error {
// 	query := `
//     CREATE TABLE IF NOT EXISTS websites(
//         id INTEGER PRIMARY KEY AUTOINCREMENT,
//         name TEXT NOT NULL UNIQUE,
//         url TEXT NOT NULL,
//         rank INTEGER NOT NULL
//     );
//     `

// 	_, err := r.db.Exec(query)
// 	return err
// }
