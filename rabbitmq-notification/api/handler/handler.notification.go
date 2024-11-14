package handler

import (
	"encoding/json"
	"net/http"

	"github.com/snykk/learn-rabbitmq/rabbitmq-notification/api/dto"
	"github.com/snykk/learn-rabbitmq/rabbitmq-notification/api/producer"
)

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	var notification dto.NotificationRequest

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Kirim notifikasi ke queue
	err := producer.SendToQueue(notification)
	if err != nil {
		http.Error(w, "Failed to send notification", http.StatusInternalServerError)
		return
	}

	// Return response
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "accepted",
		"message": "Notification will be processed",
	})
}
