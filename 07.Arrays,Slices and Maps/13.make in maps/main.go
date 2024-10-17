package main

import "fmt"

func main() {
	// use make when we know how much element can be added to the slice, which saves creation of array again and again
	// 2 - create 2 empty slot in array
	// 5 - length
	userNames := make([]string, 2, 5)

	userNames[0] = "Kevin"
	userNames[1] = "Albert"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// make in map
	// 3 - length, it can store 3 elementes without Go reallocating memory
	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.6
	courseRatings["react"] = 4.3
	courseRatings["python"] = 4.1
	fmt.Println(courseRatings)
}
