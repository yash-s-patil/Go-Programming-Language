package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	// extract value using key
	fmt.Println(websites["Amazon Web Services"])

	// add new key value pair to the map
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	// overwrite the value for a specific key
	websites["Google"] = "https://gmail.com"
	fmt.Println(websites)

	// delete a key value pair
	delete(websites, "Google")
	fmt.Println(websites)
}
