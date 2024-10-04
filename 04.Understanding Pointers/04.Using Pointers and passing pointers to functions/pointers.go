package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", agePointer)
	// de referencing - get value of the address
	fmt.Println("Age Pointer:", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

// no copy is getting created - still not recommended
func getAdultYears(age *int) int {
	return *age - 18
}
