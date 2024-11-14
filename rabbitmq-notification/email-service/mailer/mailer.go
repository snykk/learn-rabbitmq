package mailer

import (
	"fmt"
	"time"

	gomail "gopkg.in/mail.v2"
)

type Mailer interface {
	SendMessage(subject, receiver, message string) (err error)
}

type mailer struct {
	email    string
	password string
}

func NewMailerInstance(email, password string) Mailer {
	return &mailer{
		email:    email,
		password: password,
	}
}

func (mailer *mailer) SendMessage(subject, receiver, message string) (err error) {
	now := time.Now()
	configMessage := gomail.NewMessage()
	configMessage.SetHeader("From", mailer.email)
	configMessage.SetHeader("To", receiver)
	configMessage.SetHeader("Subject", subject)

	// Template HTML untuk notifikasi
	emailBody := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Notification</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
				color: #333333;
				margin: 0;
				padding: 0;
			}
			.container {
				width: 100%%;
				max-width: 600px;
				margin: 0 auto;
				background-color: #ffffff;
				padding: 20px;
				border-radius: 8px;
				box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
			}
			.header {
				text-align: center;
				padding: 10px 0;
				border-bottom: 1px solid #dddddd;
			}
			.header h1 {
				color: #4CAF50;
				font-size: 24px;
			}
			.content {
				padding: 20px 0;
			}
			.content p {
				line-height: 1.6;
				margin: 15px 0;
			}
			.footer {
				text-align: center;
				padding-top: 20px;
				border-top: 1px solid #dddddd;
				color: #888888;
				font-size: 12px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>Notification</h1>
			</div>
			<div class="content">
				<p>Dear User,</p>
				<p>%s</p> <!-- Isi pesan notifikasi -->
				<p>Thank you for using our service.</p>
			</div>
			<div class="footer">
				<p>&copy; %d Our Service. All rights reserved.</p>
			</div>
		</div>
	</body>
	</html>`, message, now.Year())

	// Set HTML email body
	configMessage.SetBody("text/html", emailBody)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, mailer.email, mailer.password)

	err = dialer.DialAndSend(configMessage)
	return
}
