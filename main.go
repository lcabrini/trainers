package main

import (
    "net/http"
)

func main() {
    auth := &Auth{}
    auth.Setup()

    fs := http.FileServer(http.Dir("/resources/"))
    http.Handle("/resources/", fs)

    http.ListenAndServe(":8000", nil)
}
