package main

import (
    "fmt"
    "net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello-world")
}

func main() {
    http.HandleFunc("/", helloWorldHandler) // Register the handler for the root URL

    fmt.Println("Server starting at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed to start:", err)
    }
}
