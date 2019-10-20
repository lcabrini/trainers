package main

import (
    "net/http"
    "github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

func main() {
    sessionManager = scs.New()

    mux := http.NewServeMux()

    auth := &Auth{}
    auth.Setup(mux)

    dashboard := &Dashboard{}
    dashboard.Setup(mux)

    fs := http.FileServer(http.Dir("static/"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
    http.ListenAndServe(":8000", sessionManager.LoadAndSave(mux))
}
