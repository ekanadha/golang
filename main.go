package main

import (
        "fmt"
        "log"
        "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, " Welcome to the DevOps World !!!")

}

func handleRequest() {
        http.HandleFunc("/", homePage)
        log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}

func main() {
        handleRequest()
}
