// a pointer stores the memory address of a value
// in Go, the type *T is a pointer to a T value, the default value for pointers is nil

package main

import "fmt"

func main() {
	var address *int  // declares an int pointer
	number := 42      // int
	address = &number // address stores the memory address of a number
	value := *address // dereferencing the value

	fmt.Printf("address: %v\n", address) // 0xc000018098
	fmt.Printf("value: %v\n", value)     // 42
}

// The & operator provides the memory address of a value. It is used to bind a pointer to a value.
// The (*) operator prefixing a type denotes a pointer type, while an asterisk prefixing a variable is used to dereference the value the variable points to
