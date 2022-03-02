package chapter_http

import (

    "io"
    "log"

    "net/http"
    "os"
)

func Test_go_client() {

    log.Println("Sending request...")
    res, err := http.Get("http://localhost:8080/hello")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Reading response...")
    if _, err := io.Copy(os.Stdout, res.Body); err != nil {
        log.Fatal(err)
    }
}
