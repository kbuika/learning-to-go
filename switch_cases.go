package main

import "fmt"

func average(x []float64) (avg float64) {
	total := 0.0

	switch len(x) {
	case 0:
		avg = 0
	default:
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}

func main() {
	x := []float64{5.3, 45.2, 21.09, 33.42, 72.7}
	fmt.Println(average(x))
}

// Go runs the selected case only, removing the need for 'break'
