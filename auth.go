package main

import (
    "fmt"
    "net/http"
)

type Auth struct {}

func (a *Auth) login(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        fmt.Fprintf(res, "Login form")

    case "POST":
        fmt.Fprintf(res, "Login handler")

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
