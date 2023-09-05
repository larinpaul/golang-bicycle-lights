package main

import (
	"fmt"
	"golang-bicycle-lights/cmd/lights-control/pkg/light"
	"strconv"
)

func main() {

	for {
		// Call the readInput function and assign its return value to envLight
		envLight := readInput()

		// Check if the user entered 0 which means to exit
		if envLight == "0" {
			fmt.Println("Thanks for using our program! :) Exiting.")
			break // Break out of the loop and end the program
		}

		// Convert the envLight string value to an int value using strconv.Atoi
		// Handle any error that may occur
		envLightInt, err := strconv.Atoi(envLight)
		if err != nil {
			fmt.Println("An error occurred while converting input. Please try again.", err)
			return
		}

		// Call the light.Check function from the light package
		// Pass the envLight variable as an argument
		// Print the returned message to the standard output
		fmt.Println(light.Check(envLightInt))
		//fmt.Println("/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\")
		fmt.Println("🚲")
	}
}
