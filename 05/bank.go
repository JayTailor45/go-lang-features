package main

import (
	"fmt"

	"jaytailor.com/bank/file_operations"
)

const accountBalanceFileName = "balance.txt"

func main() {
	fmt.Println("Welcome to the bank")

	var accountBalance, err = file_operations.GetFloatFromFile(accountBalanceFileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
	}

	for {
		presentOptions()

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
			file_operations.WriteBalanceToFile(accountBalance, accountBalanceFileName)
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
				file_operations.WriteBalanceToFile(accountBalance, accountBalanceFileName)
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
