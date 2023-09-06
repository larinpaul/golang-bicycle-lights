package main

import (
	"fmt"
	"golang-bicycle-lights/cmd/lights-control/pkg/light"
)

func main() {

	for {
		// Call the readInput function and assign its return value to envLight
		envLight := readInput()

		if envLight == "exit" {
			fmt.Println("Thanks for using our program! :) Exiting.")
			break
		}

		fmt.Println(light.Check(envLight))
		//fmt.Println("/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\")
		fmt.Println("🚲")
	}
}
