package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64
	const inflationRate = 2.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value(Adjusted for Inflation):", futureRealValue)

	// print upto one decimal value
	// fmt.Printf("Future Value: %.1f\nFuture Value(Adjusted for Inflation): %.1f \n", futureValue, futureRealValue)

	// if we want to format, change and then store text in a variable and return string use Sprint()
	//formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	//formattedRFV := fmt.Sprintf("Future Value(Adjusted for Inflation): %.1f \n", futureRealValue)
	//fmt.Print(formattedFV, formattedRFV)

	// multiline strings using Backticks(``)
	fmt.Printf(`Future Value: %.1f
		Future Value(Adjusted for Inflation): %.1f`, futureValue, futureRealValue)
}
