package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/mattn/go-sqlite3"
    "appointments/sender"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Please provide date argument: ./main YYYY-MM-DD")
    }
    summaryDate := os.Args[1]

    db, err := sql.Open("sqlite3", "./appointments.db")
    if err != nil {
        log.Fatal(err)
    }

    err = sender.GenerateDoctorSummary(db, summaryDate)
    if err != nil {
        log.Fatal(err)
    }

    err = sender.GenerateCenterSummary(db, summaryDate)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Summary messages generated successfully.")
}
