package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

func Deposit(acc *Account, sum int) {
	if sum < 0 {
		acc.Balance = 100
	} else {
		acc.Balance += sum
	}
}

func main() {
	acc := Account{Owner: "Alice", Balance: 100}
	fmt.Println(acc)
	Deposit(&acc, -20)
	fmt.Println(acc)
	fmt.Println(acc.Balance) // 150
}
