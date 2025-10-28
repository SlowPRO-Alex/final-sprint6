package server

import (
    "log"
    "net/http"
    "time"
    "github.com/SlowPRO-Alex/final-sprint6/tree/main/internal/handlers" 
)

type Server struct {
    Logger *log.Logger
    HTTPServer *http.Server
}

// Функция создания сервера
func NewServer(logger *log.Logger) *Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handlers.IndexHandler)
    mux.HandleFunc("/upload", handlers.UploadHandler)

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ErrorLog:     logger,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    return &Server{
        Logger: logger,
        HTTPServer: srv,
    }
}