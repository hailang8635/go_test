package chapter_http

import (
    "fmt"
    "log"
    "net/http"
)

func Test_go_http2() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintln(w, "hello world")
        fmt.Fprintf(w, "hello world you've requested %s\n", r.URL.Path)

        log.Println(r.URL.Query().Get("token"))
        log.Println(r.FormValue("field1"))
    })

    log.Println("Starting server...")

    http.ListenAndServe("127.0.0.1:8080", nil)

    fmt.Println("dfsd")
}

func Test_go_http3() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintln(w, "hello world")
        fmt.Fprintf(w, "hello world you've requested %s\n", r.URL.Path)

        log.Println(r.URL.Query().Get("token"))
        log.Println(r.FormValue("field1"))
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("Starting server...")


    http.ListenAndServe("127.0.0.1:8080", nil)

    fmt.Println("started")
}