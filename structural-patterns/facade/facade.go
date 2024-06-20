package main

import "fmt"

// Waiter struct represents a waiter in a restaurant
type Waiter struct{}

// TakeOrder method simulates a waiter taking an order
func (w *Waiter) TakeOrder() {
	fmt.Println("Waiter is taking the order...")
}

// Kitchen struct represents a kitchen in a restaurant
type Kitchen struct{}

// PrepareOrder method simulates a kitchen preparing an order
func (k *Kitchen) PrepareOrder() {
	fmt.Println("Kitchen is preparing the order...")
}

// Billing struct represents the billing process in a restaurant
type Billing struct{}

// ProcessPayment method simulates the billing process
func (b *Billing) ProcessPayment() {
	fmt.Println("Billing is processing the payment...")
}

// Order struct represents an order in a restaurant
// It uses the Facade design pattern to simplify the process of placing an order
type Order struct {
	waiter  Waiter
	kitchen Kitchen
	billing Billing
}

// PlaceOrder method uses the Facade design pattern to simplify the process of placing an order
func (o *Order) PlaceOrder() {
	o.waiter.TakeOrder()
	o.kitchen.PrepareOrder()
	o.billing.ProcessPayment()
}

// main function
func main() {
	// create an order
	order := Order{}
	// place the order
	order.PlaceOrder()
}
