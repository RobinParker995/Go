package main

import "fmt"

type Notifier interface {
	Send(message string)
}


type Email struct {
	Address string
}

func (e Email) Send(message string) {
	fmt.Printf("Sending Email to %s: %s \n", e.Address, message)
}


type SMS struct {
	Phone string
}


func (s SMS) Send(message string) {
	fmt.Printf("Sending SMS to %s: %s \n", s.Phone, message)
}


type Push struct {
	DeviceID string
}

func (p Push) Send(message string) {
	fmt.Printf("Sending Pysh Notifucation to %s: %s \n", p.DeviceID, message)
}


func Notify(n Notifier, message string) {
	fmt.Println("Preparing notifications...")
	n.Send(message)
	fmt.Println("Notification Sent! \n")
}

func main() {
	
	email := Email{
		Address: "rob@example.com",
	}

	sms := SMS{
		Phone: "+123456789",
	}

	push := Push{
		DeviceID: "DEVICE-001",
	}

	Notify(email, "Welcome!")
	Notify(sms, "Your OTP is 123456")
	Notify(push, "You have a new message")
}
