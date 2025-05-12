package main

import "fmt"

// NotificationService defines the interface for sending notifications
type NotificationService interface {
	Send(recipient string, message string) error
	GetServiceName() string
}

// EmailNotification implements the NotificationService interface
type EmailNotification struct {
	SenderEmail string
}

func (e *EmailNotification) Send(recipient string, message string) error {
	fmt.Printf("Sending email to %s from %s: %s\n", recipient, e.SenderEmail, message)
	return nil
}

func (e *EmailNotification) GetServiceName() string {
	return "Email Service"
}

// SMSNotification implements the NotificationService interface
type SMSNotification struct {
	PhoneNumber string
}

func (s *SMSNotification) Send(recipient string, message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", recipient, message)
	return nil
}

func (s *SMSNotification) GetServiceName() string {
	return "SMS Service"
}

type NotificationManager struct {
	services []NotificationService
}

func (nm *NotificationManager) SendAll(recipient string, message string) {
	for _, service := range nm.services {
		fmt.Printf("Using %s:\n", service.GetServiceName())
		service.Send(recipient, message)
	}
}

func main() {
	services := []NotificationService{
		EmailNotification{SenderEmail: "noreply@example.com"},
		SMSNotification{PhoneNumber: "+123456789"},
	}

	// Create notification services
	emailService := &EmailNotification{SenderEmail: "noreply@example.com"}
	smsService := &SMSNotification{PhoneNumber: "+123456789"}

	// Create notification manager
	manager := &NotificationManager{}

	// Add services to manager
	manager.services = append(manager.services, emailService)
	manager.services = append(manager.services, smsService)

	// Send notifications through all channels
	manager.SendAll("user@example.com", "Hello from notification service!")
}
