package main

import (
    "fmt"
    "logmaster/database"
    "logmaster/organizer"
)

func main() {
    database.InitDB()

    var folder string
    fmt.Println("Enter the path of log folder to organize:")
    fmt.Scanln(&folder)

    fmt.Println("Organizing logs...")
    organizer.MoveLogs(folder)
    fmt.Println("Log organization complete!")
}
