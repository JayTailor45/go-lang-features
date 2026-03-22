package main

import "fmt"

func main() {
	fmt.Println("Welcome to the bank")

	var accountBalance float64 = 1000

	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Println("How much do you want to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Withdraw balance can not exceed account balance")
				continue
			} else if withdrawAmount <= 0 {
				fmt.Println("Invalid ammount. Must be greater than 0")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
			}
		case 4:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid action!")
			return
		}
	}
}
