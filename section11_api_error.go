package main

import (
    "encoding/json"
    "errors"
    "log"
    "net/http"
)

var ErrUserNotFound = errors.New("user not found")

type HTTPError struct {
    Status  int
    Message string
    Err     error
}

func (e *HTTPError) Error() string { return e.Err.Error() }

func writeJSONError(w http.ResponseWriter, status int, msg string) {
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func handler(w http.ResponseWriter, r *http.Request) {
    err := ErrUserNotFound
    log.Println("failed:", err)
    writeJSONError(w, 404, "user not found")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
