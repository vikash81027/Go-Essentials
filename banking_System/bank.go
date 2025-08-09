package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"


func main() {
	fmt.Println(randomdata.City())
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	fmt.Println("Welcome to Go Bank !")

	/*
		// infinite for loop
		for {

		}
	*/
	// ---------------------------------
	/*
		// for loop
		for i := 1; i < 10; i++ {
			fmt.Println("Hello")
		}
	*/
	for {
		getOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositValue float64
			fmt.Scan(&depositValue)
			if depositValue <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}
			accountBalance += depositValue
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 3 {
			var withdrawlValue float64
			fmt.Scan(&withdrawlValue)
			if withdrawlValue > accountBalance {
				fmt.Println("Insufficient Balance : ")
				fmt.Println("Your balance :", accountBalance)
			} else if withdrawlValue <= 0 {
				fmt.Println("Invalid amount!")
				continue
			} else {
				fmt.Printf("%f is withdrawn", withdrawlValue)
				accountBalance -= withdrawlValue
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
				fmt.Println("Your balance updated! :", accountBalance)
			}
		} else if choice == 4 {
			fmt.Println("Thank You for yout visit!")
			break
		} else {
			fmt.Println("Not a valid option!")
		}
	}
	/*
		// switch statement (no break needed)
		choice := 10
		switch choice {
		case 1:
			fmt.Println(choice)
		case 2:
			fmt.Println(choice);
		default:
			fmt.Println("Good Bye !")
		}
	*/
}

