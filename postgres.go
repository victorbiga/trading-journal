package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
)

const (
    host     = "db"
    port     = 5432
    user     = "postgres"
    password = "example"
    dbname   = "trades"
)

func Postgres() {
    // Connection string
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var db *sql.DB
    var err error

    // Sleep for 1 second before starting the connect retry loop
    time.Sleep(2 * time.Second)
    // Number of retries
    maxRetries := 10
    retryInterval := 1 * time.Second

    for i := 1; i <= maxRetries; i++ {
        // Attempt to open a connection to the database
        db, err = sql.Open("postgres", psqlInfo)
        if err != nil {
            log.Printf("Connection attempt %d failed: %v. Retrying in %v...", i, err, retryInterval)
            time.Sleep(retryInterval)
            continue
        }

        // Try to ping the database to check the connection
        err = db.Ping()
        if err == nil {
            fmt.Println("Successfully connected to the database!")
            return
        }

        log.Printf("Ping attempt %d failed: %v. Retrying in %v...", i, err, retryInterval)
        time.Sleep(retryInterval)
    }

    log.Fatal("Max retries reached. Unable to connect to the database.")
}
