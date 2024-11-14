package main

import (
	"log"
	"net/http"

	"github.com/snykk/learn-rabbitmq/rabbitmq-notification/api/handler"
)

func main() {
	http.HandleFunc("/api/notifications", handler.NotificationHandler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
