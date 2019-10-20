package main

import (
    "fmt"
    "net/http"
)

type Dashboard struct {}

func (d *Dashboard) dashboard(res http.ResponseWriter, req *http.Request) {
    user := sessionManager.GetInt(req.Context(), "user")

    if user == 0 {
        fmt.Fprintf(res, "no user")
    } else {
        fmt.Fprintf(res, "a user with ID: %d", user)
    }
}

func (d *Dashboard) Setup(mux *http.ServeMux) {
    mux.HandleFunc("/", d.dashboard)
}
