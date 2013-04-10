package main

import (
        "flag"
        "fmt"
        "log"
        "net/http"
)

var (address string
port string)

func init () {
        fAddr := flag.String("a", "127.0.0.1", "address to listen on")
        fPort := flag.Int("p", 8080, "port to listen on")
        flag.Parse()
        port = fmt.Sprintf("%d", *fPort)
        address = *fAddr
}

func main() {
        listenAddr := fmt.Sprintf("%s:%s", address, port)
        http.Handle("/", http.FileServer(http.Dir("static")))
        log.Printf("listening on http://%s", listenAddr)
        log.Fatal(http.ListenAndServe(listenAddr, nil))
}
