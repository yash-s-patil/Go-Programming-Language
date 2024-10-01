package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5 // initial value
	var years float64         // if we dont have any initial value we have to use var keyword along with data type
	const inflationRate = 2.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
