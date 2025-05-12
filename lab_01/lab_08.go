package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string) bool
}

type EmailNotifier struct {
	EmailAddress string
}

func (e EmailNotifier) Send(message string) bool {
	fmt.Printf("Sending Email to %s: %s\n", e.EmailAddress, message)
	return true
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) bool {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return true
}

type NotificationService struct {
	notifiers []Notifier
}

func (ns *NotificationService) AddNotifier(notifier Notifier) {
	ns.notifiers = append(ns.notifiers, notifier)
}

func (ns *NotificationService) NotifyAll(message string) []bool {
	results := make([]bool, len(ns.notifiers))
	for i, notifier := range ns.notifiers {
		results[i] = notifier.Send(message)
	}
	return results
}

func main() {
	service := NotificationService{}

	// Create notifiers
	emailNotifier := EmailNotifier{EmailAddress: "user@example.com"}
	smsNotifier := SMSNotifier{PhoneNumber: "123-456-7890"}

	// Add notifiers to the service
	service.AddNotifier(emailNotifier)
	service.AddNotifier(smsNotifier)

	// Send notifications through all channels
	service.NotifyAll("System maintenance scheduled tonight")

	fmt.Println("All notifications sent")
}
