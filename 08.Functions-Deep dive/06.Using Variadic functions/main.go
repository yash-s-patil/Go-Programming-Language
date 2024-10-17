package main

import "fmt"

func main() {
	// numbers := []int{1, 10, 5}
	sum := sumup(1, 5, 4, 3, 5)

	fmt.Println(sum)
}

// you can have as many single parameter value as you want
func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
