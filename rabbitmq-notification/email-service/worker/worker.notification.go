package worker

import (
	"encoding/json"
	"log"

	"github.com/snykk/learn-rabbitmq/rabbitmq-notification/email-service/dto"
	"github.com/snykk/learn-rabbitmq/rabbitmq-notification/email-service/mailer"
	"github.com/streadway/amqp"
)

func ConsumeNotifications(ch *amqp.Channel, mailerInstance mailer.Mailer) {
	// Deklarasi queue
	q, err := ch.QueueDeclare(
		"notification_queue",
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Menarik pesan dari queue
	msgs, err := ch.Consume(
		q.Name, // Queue
		"",     // Consumer
		false,  // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Fungsi goroutine untuk memproses pesan
	go func() {
		for d := range msgs {
			var notification dto.NotificationRequest
			err := json.Unmarshal(d.Body, &notification)
			if err != nil {
				log.Printf("Error decoding JSON: %v", err)
				d.Nack(false, true) // Nack untuk mencoba pesan kembali
				continue
			}

			// Logika pemrosesan notifikasi
			log.Println("Processing notification")
			log.Printf("noticiation recipient: %s\n", notification.Recipient)
			log.Printf("noticiation subject: %s\n", notification.Subject)
			log.Printf("noticiation message: %s\n", notification.Message)

			// Contoh: Kirim email atau SMS berdasarkan tipe notifikasi
			err = mailerInstance.SendMessage(notification.Subject, notification.Recipient, notification.Message)
			if err != nil {
				log.Printf("Error when sending message: %v\n", err)
			} else {
				log.Printf("Success sending email to %s\n", notification.Recipient)
			}

			// Acknowledge pesan
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	select {} // block runtime using select{} instead go chan
}
