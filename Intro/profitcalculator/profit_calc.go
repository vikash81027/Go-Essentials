package profitcalculator

import (
	"errors"
	"fmt"
	"os"

	"example.com/investment_calculator/functions"
)


func CalculateProfit() (float64, float64, float64, error){
	var revenue, expenses, taxRate float64

	revenue = functions.GetUserInput("Enter your revenue: ");
	expenses = functions.GetUserInput("Enter your expenses: ");
	taxRate = functions.GetUserInput("Enter your tax-rate: ");
	if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		return 0, 0, 0, errors.New("values can not be less than 0")
	}
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	text := fmt.Sprintf("ebt %f\nprofit %f\nratio %f\n",ebt, profit, ratio)
	os.WriteFile("profitCalc.txt", []byte(text), 0644)
	return ebt, profit, ratio, nil
}