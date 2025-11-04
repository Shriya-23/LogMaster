package database

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3" // stable version v1.14.16
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./log_movement.db")
    if err != nil {
        log.Fatal(err)
    }

    query := `CREATE TABLE IF NOT EXISTS log_files(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        filename TEXT,
        old_path TEXT,
        new_path TEXT,
        moved_at DATETIME
    );`

    _, err = DB.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}

func InsertLog(filename, oldPath, newPath string) {
    _, err := DB.Exec(
        "INSERT INTO log_files(filename, old_path, new_path, moved_at) VALUES (?, ?, ?, datetime('now'))",
        filename, oldPath, newPath,
    )
    if err != nil {
        log.Println("Error inserting log:", err)
    }
}


