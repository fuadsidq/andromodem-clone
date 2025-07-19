package main

import (
	"fmt"
	"log"
	"net/http"

	"andromodem/backend/handlers"
	"andromodem/backend/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/devices", handlers.DevicesHandler)
	mux.HandleFunc("/api/restart", handlers.RestartHandler)
	mux.HandleFunc("/api/logs", handlers.LogsHandler)

	// Bungkus dengan middleware
	handlerWithMiddleware := middleware.CORS(middleware.Logging(mux))

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithMiddleware))
}
