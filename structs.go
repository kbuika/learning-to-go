// A struct is just a collection of fields

package main

import "fmt"

// define a stack type using a struct
type stack struct {
	index int
	data  [5]int
}

// define push and pop methods
func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

// notice the stack pointer s passed as an argument
func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

func main() {
	// create a pointer to the new stack and push 2 values
	s := new(stack)
	s.push(23)
	s.push(14)

	fmt.Printf("stack: %v\n", *s) //   stack: {2 [23 14 0 0 0]}
}

/*
First, we define our custom type that represents a stack. To achieve the stack functionality, we need an array to store the stack elements,
and an index to point to the last item in the stack. For the sake of the example, let’s fix our stack size to 5 elements. Inside the body of the struct,
 we specify an index field which is of type int, and a field called data, which is an array of 5 int elements.

Next we define the push and pop methods. A method is a special kind of function that takes a receiver argument between the func keyword and the method name.
 Notice the type of the parameter s. In this case, it is a stack pointer instead of a stack. By default, Go doesn’t pass values by reference.
 Instead, if we were to omit the asterisk, Go would pass a copy of our stack, meaning the original stack would not be modified by our methods.

In the body of our stack methods, we access the stack fields using the dot notation.
In the push method, we write a given integer k to the first available index (recall the default value of a declared int is 0), and increment the index by 1.
In the pop method we decrement the index by 1, and return the last item in the stack. In the body of the main function, we use new() to create a pointer to a newly allocated stack.
 We then push 2 items and write the result to standard output.
*/
