package main

import (
	"fmt"
	"math"
)

func main() {
	// when using explicit type assignment we cannot use := and we need to use var keyword
	var investmentAmount float64 = 1000
	// when we are not using explicit type assign, we can use := and no need to use var keyword
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
