package main

import (
    "fmt"
    "net/http"
    "html/template"
    "github.com/google/uuid"
    "database/sql"
)

type Auth struct {}

func authenticate(username string, password string) (bool, uuid.UUID) {
    var test uuid.UUID

    query := `
    select user_id 
    from users 
    where username = $1 and password = $2`
    row := db.QueryRow(query, username, password)
    switch err := row.Scan(&test); err {
    case sql.ErrNoRows:
        return false, test
    case nil:
        return true, test
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
        ret, uid := authenticate(username, password)
        if ret == true {
            // fmt.Fprintf(res, "Logged in, uid = %s", uid.String())
            sessionManager.Put(req.Context(), "user", uid.String())
            http.Redirect(res, req, "/", 301)
        } else {
            fmt.Fprintf(res, "Authentication failed")
        }

    default:
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (a *Auth) logout(res http.ResponseWriter, req *http.Request) {

}

func (a *Auth) Setup(mux *http.ServeMux) {
    mux.HandleFunc("/login", a.login)
    mux.HandleFunc("/logout", a.logout)
}
