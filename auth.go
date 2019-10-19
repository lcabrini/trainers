package main

import (
    "fmt"
    "net/http"
    "html/template"
    "database/sql"
)

type Auth struct {}

func authenticate(username string, password string) bool {
    var test string

    query := `
    select username 
    from users 
    where username = $1 and password = $2`
    row := DB.QueryRow(query, username, password)
    switch err := row.Scan(&test); err {
    case sql.ErrNoRows:
        return false
    case nil:
        return true
    default:
        panic(err)
    }
}


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
        req.ParseForm()
        username := req.FormValue("username")
        password := req.FormValue("password")
        if authenticate(username, password) {
            fmt.Fprintf(res, "Logged in")
        } else {
            fmt.Fprintf(res, "Authentication failed")
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
