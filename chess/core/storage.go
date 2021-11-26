package core

import (
	"database/sql"
	"encoding/json"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitGameStorage() error {
	_db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		return err
	}

	createStatement := `CREATE TABLE IF NOT EXISTS games (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created TEXT NOT NULL,
	updated TEXT NOT NULL,
	info TEXT NOT NULL
)`
	_, err = _db.Exec(createStatement)
	if err != nil {
		return err
	}

	db = _db
	return nil
}

func insertGame(g *Game) error {
	gameJson, err := json.Marshal(g)
	if err != nil {
		return err
	}

	now := time.Now().String()
	result, err := db.Exec("INSERT INTO games (created, updated, info) VALUES (?, ?, ?)", now, now, string(gameJson))
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	g.ID = id

	return nil
}

func updateGame(g *Game) error {
	gameJson, err := json.Marshal(g)
	if err != nil {
		return err
	}

	now := time.Now().String()
	_, err = db.Exec("UPDATE games SET created = ?, updated = ?, info = ? WHERE id = ?", now, now, string(gameJson), g.ID)
	if err != nil {
		return err
	}

	return nil
}
