package main

import "fmt"

func main() {
	// use make when we know how much element can be added to the slice, which saves creation of array again and again
	// 2 - create 2 empty slot in array
	// 5 - capacity
	userNames := make([]string, 2, 5)

	userNames[0] = "Kevin"
	userNames[1] = "Albert"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)
}
