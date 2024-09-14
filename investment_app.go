package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// scan the terminal for user input
	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	// save formatted data and text in variables
	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	// output information
	// fmt.Println("Future value:", futureValue)
	// fmt.Printf("Future value: %.2f\n", futureValue)
	// fmt.Printf("Future value (adjusted for inflation): %.2f", futureRealValue)

	// fmt.Print(formattedFV, formattedRFV)
	outputText(formattedFV, formattedRFV)
}

func outputText(text, text2 string){
	fmt.Print(text, text2)
}

// function that returns 2 float64 values
func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)

	return fv, rfv
	// return
}
