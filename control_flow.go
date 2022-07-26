/*
suppose we’re given a slice containing floats, and we’re interested in computing their average value.
We’ll proceed by creating a function called average that takes a slice as a parameter and returns a float called avg
*/

package main

import "fmt"

// the retuen parameter ``avg`` is defined immediately at the end of the function declaration
func average(x []float64) (avg float64) {
	total := 0.0

	if len(x) == 0 {
		avg = 0
	} else {
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}

	return
}

func main() {
	x := []float64{5.75, 4.34, 23.65, 90.21}
	fmt.Println(average(x))
}
