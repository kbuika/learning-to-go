// a lazier way to write for loops

package main

import "fmt"

func main() {
	// our array
	DeploymentOptions := [...]string{"AWS", "Azure", "GCP", "Heroku"}

	// the loop
	for i, option := range DeploymentOptions {
		fmt.Println(i, option)
	}
}
