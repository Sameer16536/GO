package main

import (
	"errors"
	"fmt"
)

var ErrOutOfStock = errors.New("Item is out of Stock")
var ErrPaymentFailed = errors.New("Payment Failed")

// Func to Check if the item is in stock
func checkAvailability(item string) error {
	if item == "Pizza" {
		return ErrOutOfStock
	}
	return nil
}

// FUnc to check if the payment is successful
func makePayment(amount int) error {
	if amount == 0 {
		return ErrPaymentFailed
	}
	return nil
}

// Func to order an item
func placeOrder(item string, amount int) error {
	if err := checkAvailability(item); err != nil {
		return fmt.Errorf("placing order : %w", err)
	}

	if err := makePayment(amount); err != nil {
		return fmt.Errorf("placing order : %w", err)
	}
	return nil
}

func main() {
	// Sample orders
	orders := []struct {
		item   string
		amount int
	}{
		{"Burger", 100}, // Valid order
		{"Pizza", 150},  // Out-of-stock item
		{"Pasta", 0},    // Payment failed
		{"Salad", 80},   // Valid order
	}
	for _, order := range orders {
		fmt.Println("Placing order for:", order.item)

		err := placeOrder(order.item, order.amount)
		if err != nil {
			// Handle specific errors
			if errors.Is(err, ErrOutOfStock) {
				fmt.Println("Error: Item is not available. Please choose another item.")
			} else if errors.Is(err, ErrPaymentFailed) {
				fmt.Println("Error: Payment failed. Please check your payment method.")
			} else {
				fmt.Printf("Error: An unknown issue occurred: %s\n", err)
			}
		} else {
			fmt.Println("Order placed successfully!")
		}
		fmt.Println()
	}

}
