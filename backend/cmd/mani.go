package main

import (
    "log"
    "github.com/student-nirajkumar/shopping-challenge/backend/internal/server"
)

func main() {
    if err := server.Start(); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
