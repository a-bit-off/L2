/*
Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Strategy_pattern
*/

package pattern

import "fmt"

type Payment interface {
	Pay() error
}

// ------------------------------

type cardPayment struct {
	cardNumber string
	cvv        string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	// ... implementation
	return nil
}

// ------------------------------

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

func (p *payPalPayment) Pay() error {
	// ... implementation
	return nil
}

// ------------------------------

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	// ... implementation
	return nil
}

// ------------------------------
func processOrder(product string, payment Payment) {
	// ... implementation
	err := payment.Pay()
	if err != nil {
		return
	}
	fmt.Println(product, payment)
}

// Main
func RunStrategy() {
	product := "ticket"
	payWay := 1

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("1234 5678 9101 1121", "123")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}
