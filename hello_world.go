package main // required for a standalone executable

import "fmt" // implements formatted I/O

// when this program is executed the first function that runs is main.main()
func main() {
	fmt.Println("Hello, world") // call Println from the fmt package
}

// A package is a collection of source files in the same directory that allows variables, types and functions to be visible among other source
// files within the same package. For standalone files like this one, the package is called main.
