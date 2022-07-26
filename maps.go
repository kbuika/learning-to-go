// a map is a data structure that stores key-value pairs with a constant look-up line

package main

import "fmt"

func main() {

	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996, "Python": 1991,
		"Javascript": 1996, "Go": 2012,
	}

	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}
}

// the output is in random order because the efficiency of maps comes at the expense
//   of randomizing the order of the keys and associated values
