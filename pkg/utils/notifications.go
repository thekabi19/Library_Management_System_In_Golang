package utils

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

// Function to simulate sending a notification email in a Goroutine (Malek, 2024)
func SendNotification(bookName string, done chan bool) {
	smtpHost := "smtp.gmail.com"      // Gmail SMTP host
	smtpPort := 587                   // SMTP port
	senderEmail := "sender@gmail.com" //sender Gmail email
	senderKey := "ubqffxj"            // replace the SMTP pass

	// Simulate some delay before sending email
	fmt.Printf("Starting email notification process for the book: %s...\n", bookName)
	time.Sleep(2 * time.Second)

	// Construct the email message.
	to := "recipient@gmail.com" // recipient's email
	subject := "New Book Added: " + bookName
	body := fmt.Sprintf("A new book titled '%s' has been added to the bookstore.", bookName)

	// Create the new email message using gomail
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body) // Sending a plain text email

	// Create the SMTP dialer using gmail smtp server credentials
	dialer := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderKey)

	// Send the email asynchronously.
	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Printf("Failed to send notification for book: %s. Error: %s\n", bookName, err.Error())
	} else {
		fmt.Printf("Notification email sent for the book: %s\n", bookName)
	}

	// Signal that the Goroutine is done.
	done <- true
}

/*
// Function to simulate sending a notification in a Goroutine

	func SendNotification(bookName string, done chan bool) {
		fmt.Printf("Starting notification for book: %s...\n", bookName)
		// Simulating delay in sending notification
		time.Sleep(3 * time.Second)
		fmt.Printf("Notification sent for the book: %s!\n", bookName)

		// Signal completion of notification task via channel
		done <- true
	}
*/
