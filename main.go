package main

import (
    "net/http"
    "github.com/alexedwards/scs/v2"
)

var SessionManager *scs.SessionManager

func main() {
    mux := http.NewServeMux()

    auth := &Auth{}
    auth.Setup(mux)

    dashboard := &Dashboard{}
    dashboard.Setup(mux)

    fs := http.FileServer(http.Dir("static/"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    http.ListenAndServe(":8000", SessionManager.LoadAndSave(mux))
}
