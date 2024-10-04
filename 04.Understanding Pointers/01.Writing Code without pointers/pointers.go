package main

import "fmt"

func main() {
	age := 32 // regular variable
	fmt.Println("Age:", age)
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

// when passing value to a func, a copy of that value is created and stored somewhere in memory. Once function execution is completed garbage collector deletes the copied value.
func getAdultYears(age int) int {
	return age - 18
}
