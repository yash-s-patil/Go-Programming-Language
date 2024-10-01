// we use the package name as "main" because it tells go that this package will be the main entry point of the application we are building
package main

// inbuilt package provided by Go
import "fmt"

// main function tell which code should be executed first.
// there will be only one main function inside main package
func main() {
	// Print is a function and "Hello World" is the value
	fmt.Print("Hello World")
}

// You can have multiple files that make up one package
// You can multiple packages in One Go project

// packages are used to organize your code
// we use packages so we can export and import features across our files

// One module consist of multiple packages, in many cases a module is simply a go project

// go mod init example.com/first-app -> to create a module
// go build - to build go project and create executable file
// first-app.exe -> run the executable file
