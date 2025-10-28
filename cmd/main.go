package main

import (
	"log"
	"os"

	"github.com/SlowPRO-Alex/final-sprint6/tree/first-iteration/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	srv := server.NewServer(logger)
	logger.Printf("Сервер запущен на порту :8080")
	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
