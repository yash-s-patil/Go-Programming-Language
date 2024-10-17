package main

import "fmt"

func main() {
	var productNames [3]string = [3]string{"A Book"}
	prices := [4]float64{4.3, 3.2, 6.5, 4.8}
	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[2])
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

}
