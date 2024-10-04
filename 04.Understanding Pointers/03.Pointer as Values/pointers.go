package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)
	// de referencing - get value of the address
	fmt.Println("Age Pointer:", *agePointer)
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
