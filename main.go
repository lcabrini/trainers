package main

import (
    "net/http"
)

func main() {
    http.ListenAndServer(":8000", nil)
}
