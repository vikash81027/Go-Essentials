package functions

import "fmt"

func GetUserInput(prompt string) (userInput float64) {
	fmt.Print(prompt);
	fmt.Scan(&userInput);
	return
}