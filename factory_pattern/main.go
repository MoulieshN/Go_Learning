package main

import (
	"fmt"
)

// Step 1: Define an Interface
type PaymentMethod interface {
	Pay(amount int) string
}

// Step 2: Create Concrete Types
type CreditCard struct{}

func (c *CreditCard) Pay(amount int) string {
	return fmt.Sprintf("%d paid using Credit Card", amount)
}

type DebitCard struct{}

func (d *DebitCard) Pay(amount int) string {
	return fmt.Sprintf("%d paid using Debit Card", amount)
}

type UPI struct{}

func (u *UPI) Pay(amount int) string {
	return fmt.Sprintf("%d paid using UPI", amount)
}

// Step 3: Implement the Factory
func GetPaymentMethod(method string) PaymentMethod {
	switch method {
	case "credit":
		return &CreditCard{}
	case "debit":
		return &DebitCard{}
	case "upi":
		return &UPI{}
	default:
		return nil
	}
}

func main() {
	// Example Usage
	payment := GetPaymentMethod("credit")
	if payment != nil {
		fmt.Println(payment.Pay(100))
	} else {
		fmt.Println("Invalid payment method")
	}

	payment = GetPaymentMethod("upi")
	if payment != nil {
		fmt.Println(payment.Pay(150))
	} else {
		fmt.Println("Invalid payment method")
	}
}
