package utils

import (
	"fmt"
	"time"
)

// Function to simulate sending a notification in a Goroutine
func SendNotification(bookName string, done chan bool) {
	fmt.Printf("Starting notification for book: %s...\n", bookName)
	// Simulating delay in sending notification
	time.Sleep(3 * time.Second)
	fmt.Printf("Notification sent for the book: %s!\n", bookName)

	// Signal completion of notification task via channel
	done <- true
}
