/*
Storing a number of elements in a list can be achieved using arrays,
slices and maps(Go's version of hash-maps).
*/

/*
Arrays are defined by their fixed size and a common data type for all elements.
Interestingly, the size of the array is part of the type, meaning arrays cannot grow or shrink, otherwise, they would have a different type.
Array elements are accessed using square brackets
*/

package main

import "fmt"

func main() {
	// array of size 4 that stores deplloyment options
	var DeploymentOptions = [4]string{"AWS", "Azure", "GCP", "Heroku"}

	// loop through the array and print each element
	for i := 0; i < len(DeploymentOptions); i++ {
		fmt.Println(i, DeploymentOptions[i])
	}
}

// mental check: notice missing parenthesis around the loop condition
