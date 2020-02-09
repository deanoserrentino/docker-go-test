package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", strings.Replace(r.URL.Path,"/","",1))
    })

    http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "bar")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
