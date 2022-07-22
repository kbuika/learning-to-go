package main

import "fmt"

func main() {
	// user specified types
	const a int32 = 123         // int32 is a 32-bit signed integer
	const b float32 = 1.23      // float32 is a 32-bit floating-point number
	const c complex128 = 1 + 2i // complex128 is a complex number with floating-point real and imaginary parts
	const d uint16 = 14         // uint16 is a 16-bit unsigned integer

	// type inference or default types
	n := 123            // int is inferred
	pi := 3.14          // float64 is inferred
	x, y := true, false // bool is inferred
	z := "Go"           // string is inferred

	fmt.Printf("user specified types: \n%T %T %T %T \n", a, b, c, d)
	fmt.Printf("type inference or default types:\n %T %T %T %T %T \n", n, pi, x, y, z)
}

// %T is a verb and it stands for the type of the passed variable
// %v is a verb that stands for the value of the passed variable
// %d ... decimal integers
// %s ... strings
// %f ... floating-point numbers
// %t ... booleans
// %c ... runes (unicode code points)

// int is an alias for either int32 or int64
