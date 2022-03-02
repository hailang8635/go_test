package chapter_http

import (
    "fmt"

    "log"
    "net"
    "net/http"

)

func Test_go_http1() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, playground")
        log.Println("request..")
    })

    log.Println("Starting server...")
    l, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    //    go func() {
    log.Fatal(http.Serve(l, nil))
    //    }()


}
