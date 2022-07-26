package main

import (
	"fmt"
	"reflect"
)

func main() {
	languages := [9]string{
		"Go", "Python", "Ruby", "Java", "C++", "C", "JavaScript", "PHP", "C#",
	}

	// define slices
	classics := languages[:5]   // first 5 elements, could also be languages[0:5]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[5:9]     // last 4 elements,
	new := languages[7:]        // all elements after 7, could also be languages[7:9]

	fmt.Printf("classic languages: %v\n", classics)
	fmt.Printf("modern languages: %v\n", modern)
	fmt.Printf("new languages: %v\n", new)

	allLangs := languages[:]                     // copy of the array
	fmt.Println(reflect.TypeOf(allLangs).Kind()) // slice

	// web frameworks
	// leaving the brackets blank creates a slice
	// adding a value in the brackets creates an array
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte", "Laravel", "Django",
		"Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]          // length 4 capacity 4
	frameworks = append(frameworks, "Meteors") // this is not possible with arrays

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrameworks)
}
