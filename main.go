package main

import (
    "net/http"
)

func main() {
    auth := &Auth{}
    auth.Setup()

    http.ListenAndServe(":8000", nil)
}
