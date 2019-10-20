package main

import (
    "fmt"
    "net/http"
)

func dashboard(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "here")
}

type Dashboard struct {}

func (d *Dashboard) Setup(mux *http.ServeMux) {
    mux.HandleFunc("/", dashboard)
}
