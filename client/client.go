package main

import (
  "crypto/tls"
  "fmt"
  "io/ioutil"
  "log"
  "net"
  "net/http"
  "golang.org/x/net/http2"
)

const url = "https://localhost:8000"

func main() {
    client := &http.Client{
        Transport: &http2.Transport{
            AllowHTTP: true,
            DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
                return net.Dial(network, addr)
            },
        },
    }
  // Perform the request
  resp, err := client.Get(url)

  if err != nil {
     log.Fatalf("Failed get: %s", err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
     log.Fatalf("Failed reading response body: %s", err)
  }
  fmt.Printf(
    "Got response %d: %s %s\n",
    resp.StatusCode, resp.Proto, string(body))
