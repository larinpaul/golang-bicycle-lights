package main

import (
	"fmt"
	"golang-bicycle-lights/cmd/lights-control/pkg/light"
)

type Outputter interface {
	// Display takes a string and displays it to the user
	Display(string)
}

// ConsoleOutputter is a type that implements the Outputter interface using fmt.Println
type ConsoleOutputter struct{}

// Display takes a string and displays it to the standard output using fmt.Println
func (c *ConsoleOutputter) Display(s string) {
	fmt.Println(s)
}

func run(output Outputter) {

	defer func() {                    // Declare a deferred function that will run after main returtns
		if r := recover(); r != nil { // check if there was a panic
			fmt.Println("An error occurred while reading input. Please try again.", r) // Print a friendly error message
			main()                                                                     // Call main again to restart the program                                                                     // Call main again to restart the program
		}
	}()

	running := true

	for running {
		// Call the readInput function and assign its return value to envLight
		envLight := readInput()

		if envLight == -1 {
			fmt.Println("Thanks for using our program! :) Exiting.")
			running = false
		} else {
			output.Display(light.Check(envLight)) // Display the result of the light.Check function
			output.Display("🚲")                   // Display the bicycle emoji
		}
	}

}

func main() {
	// Create an outputter instance
	output := &ConsoleOutputter{}

	// Call the run function with the outputter instance
	run(output)
}
