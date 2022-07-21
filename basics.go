package main

import "fmt" // implements formatted I/O

// declare a single variable
var a int

// declare a group of variable
var (
	b bool
	c float32
	d string
)

func main() {
	a = 40           // assign sigle value
	b, c = true, 4.2 // assign multiple values
	d = "Hello, world"
	fmt.Println(a, b, c, d)
}

// var can be replaced with const when we need constants.
