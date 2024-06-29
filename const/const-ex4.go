package main

import "fmt"

// Skipping Values
const (
	ErrVirtualCardPayTransactionFailed = "Transaction amount is more than the transaction limit of %.2f"
)

func main() {
	msg := fmt.Sprintf(ErrVirtualCardPayTransactionFailed, 4445.8699)
	fmt.Println("RESP:", msg)
}
