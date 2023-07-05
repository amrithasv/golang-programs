package main


import (
    "fmt"
    "log"
    "net/http"
)


func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello,World!")
    }
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
    log.Println("Server is listening on the port 8080...")
    log.Fatal(server.ListenAndServe())
}


