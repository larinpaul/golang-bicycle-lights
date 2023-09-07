package main

import (
	"fmt"
)

func readInput() string {
	// Declare a variable to store the user input
	var input string

	// Print a message to ask the user to enter a number from 1 to 5
	fmt.Println("Receiving information about the light... This is a basic concept demonstration, please type out the intensity of the light in the environment. Write out a single number from 1 to 5, with 5 being the brightest and 1 being the dimmest. Type \"exit\" to exit")

	// Use fmt.Scanln to read a line of text from the standard input
	// Store it in the input variable
	// Handle any error that may occur
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.", err)
		panic(err)
	}

	// Return the input as the output of the function
	return input
}
