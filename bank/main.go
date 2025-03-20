package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())
	fmt.Println("====")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance: %.2f\n", accountBalance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance += amount
			fmt.Printf("New balance: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		case 3:
			fmt.Print("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if amount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}

			accountBalance -= amount
			fmt.Printf("New balance: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}
