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

	// [0,1,2]
	featuredPrices := prices[:3]
	fmt.Println(featuredPrices)

	// [1,2,3]
	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)

	// if we modify the sliced array, orginal array gets changed
	featuredPrices[0] = 322.2
	fmt.Println(prices)

	highlightedPrices := featuredPrices[:2]
	fmt.Println("Highlighted prices:", highlightedPrices)

	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// we can still select more values from slice, but only towards right
	highlightedPrices = highlightedPrices[:3]
	fmt.Println("Highlighted prices:", highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

}
