package main

// built in go packages and internal
import (
	"fmt"
	"math"

	"example.com/investment_calculator/profitcalculator"
)

const inflationRate = 2.5
// main function is executed by Go
func main() {
	getInvestment()
	ebt, profit, ratio, err := profitcalculator.CalculateProfit()
	if err != nil {
		fmt.Println(err)
		panic("oops")
	} else {
		fmt.Println(ebt, profit, ratio)
	}

	// profitcalculator.CalculateProfit()
	// profitcalculator.Add()
	
}

func getInvestment() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formatedFV := fmt.Sprintf("Future value is %.1f\n", futureValue);
	fmt.Println(formatedFV)
	fmt.Println("future value is", futureValue)
	fmt.Printf(`Future value is %.1f
							Future inflated value is %.1f\n`, futureValue, futureRealValue)
}

/*
func outputText(text1, text2 string) {
	fmt.Print(text1, text2);
}
*/

/*
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv := fv / math.Pow(1 + inflationRate / 100, years)

	return fv, rfv
}
*/

// alernate return syntax
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return 
}