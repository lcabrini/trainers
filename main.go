package main

import (
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    auth := &Auth{}
    auth.Setup(mux)

    fs := http.FileServer(http.Dir("static/"))
    mux.Handle("/", fs)

    http.ListenAndServe(":8000", mux)
}
