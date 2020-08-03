package main

import (
  "net/http"
  "fmt"
)
func title(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "This is my server")
}

func sayHello (w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "hello world\n")
}

func main () {
  http.HandleFunc("/hello", sayHello)
  http.HandleFunc("/", title)

  http.ListenAndServe(":8080", nil)
}
