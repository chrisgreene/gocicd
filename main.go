package main

import (
  "fmt"
  "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "home")
} 

func main() {
  http.HandleFunc("/", homeHandler)
  http.ListenAndServe(":8080", nil)
}
