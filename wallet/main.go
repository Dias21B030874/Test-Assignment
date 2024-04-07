package main

import (
	"fmt"

	"wallet"
)

func main() {
	w := wallet.Wallet{}
	w.Deposit(10)
	fmt.Println("Balance:", w.Balance())
}
