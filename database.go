package main

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB

func init() {
    var err error

    psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s
        password=%s dbname=%s sslmode=disable`,
        host, port, user, password, dbname)
        db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    //defer DB.Close()
    err = db.Ping()
    if err != nil {
        panic(err)
    }
}
