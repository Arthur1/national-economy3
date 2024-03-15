package main

import (
	"log/slog"
	"net/http"

	"github.com/Arthur1/national-economy3/internal/logger"
	"github.com/Arthur1/national-economy3/internal/server"
)

func main() {
	logger := logger.New(&logger.Config{
		UseConsoleHandler: true,
	})
	slog.SetDefault(logger)

	router := server.NewRouter()
	http.ListenAndServe(":8000", router)
}
