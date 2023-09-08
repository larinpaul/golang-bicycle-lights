package main

import (
	"fmt"
	"golang-bicycle-lights/cmd/lights-control/pkg/light"
)

func main() {

	defer func() { // Declare a deferred function that will run after main returtns
		if r := recover(); r != nil { // check if there was a panic
			fmt.Println("An error occurred while reading input. Please try again.", r) // Print a friendly error message
			main()                                                                     // Call main again to restart the program
		}
	}()

	running := true

	for running {
		// Call the readInput function and assign its return value to envLight
		envLight := readInput()

		if envLight == "exit" {
			fmt.Println("Thanks for using our program! :) Exiting.")
			running = false
		}

		fmt.Println(light.Check(envLight))
		fmt.Println("🚲")
	}
}
