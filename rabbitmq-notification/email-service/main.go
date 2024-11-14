package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/snykk/learn-rabbitmq/rabbitmq-notification/email-service/mailer"
	"github.com/snykk/learn-rabbitmq/rabbitmq-notification/email-service/worker"
	"github.com/streadway/amqp"
)

func main() {
	// Menghubungkan ke RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Memuat file .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mailerInstance := mailer.NewMailerInstance(os.Getenv("MAILER_EMAIL"), os.Getenv("MAILER_PASS"))

	log.Println("Mailer instance established successfully")

	worker.ConsumeNotifications(ch, mailerInstance)
}
