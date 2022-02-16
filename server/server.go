package main

import (
    "log"
    "net/http"
    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"
    "github.com/micjerry/nglib/sbi"
)

func main() {
    h2s := &http2.Server{}

    handler := http.HandlerFunc(handle)
    srv := &http.Server{Addr:":8000", Handler: h2c.NewHandler(handler, h2s)}

    log.Printf("Serving on https://0.0.0.0:8000")
    log.Printf("SBI port %d", sbi.OGS_SBI_HTTP_PORT)
    log.Fatal(srv.ListenAndServe())
}

func handle(w http.ResponseWriter, r *http.Request) {
    log.Printf("Got connection: %s", r.Proto)
    w.Write([]byte("Hello"))
}
