package organizer

import (
    "log"
    "os"
    "path/filepath"

    "logmaster/database"
)

func MoveLogs(sourceFolder string) {
    files, err := os.ReadDir(sourceFolder)
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        if f.IsDir() {
            continue
        }

        filePath := filepath.Join(sourceFolder, f.Name())
        info, err := os.Stat(filePath)
        if err != nil {
            log.Println("Cannot read file info:", err)
            continue
        }

        // Get modification date
        modTime := info.ModTime()
        dateFolder := filepath.Join(sourceFolder, modTime.Format("2006-01-02"))

        // Create folder if not exists
        if _, err := os.Stat(dateFolder); os.IsNotExist(err) {
            os.MkdirAll(dateFolder, os.ModePerm)
        }

        newPath := filepath.Join(dateFolder, f.Name())

        // Move file
        err = os.Rename(filePath, newPath)
        if err != nil {
            log.Println("Error moving file:", err)
            continue
        }

        // Log to database
        database.InsertLog(f.Name(), filePath, newPath)
        log.Println("Moved:", f.Name(), "to", newPath)
    }
}

