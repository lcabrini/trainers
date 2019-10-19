package main

import (
    "fmt"
    "net/http"
    "html/template"
    "database/sql"
    _ "github.com/lib/pq"
)

type Auth struct {}

func (a *Auth) login(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        files := []string{
            "templates/login_form.tmpl",
            "templates/base.tmpl",
        }
        ts, err := template.ParseFiles(files...)
        if err != nil {
            fmt.Println(err.Error())
        }
        ts.Execute(res, nil)

    case "POST":
        fmt.Fprintf(res, "Login handler")
        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
            "password=%s dbname=%s sslmode=disable", 
            host, port, user, password, dbname)
        db, err := sql.Open("postgres", psqlInfo)
        if err != nil {
            panic(err)
        }
        defer db.Close()
        err = db.Ping()
        if err != nil {
            panic(err)
        }

    default:
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (a *Auth) logout(res http.ResponseWriter, req *http.Request) {

}

func (a *Auth) Setup() {
    http.HandleFunc("/login", a.login)
    http.HandleFunc("/logout", a.logout)
}
