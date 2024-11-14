package dto

// NotificationRequest struct
type NotificationRequest struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}
