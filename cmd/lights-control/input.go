package main

import (
	"fmt"
	"golang-bicycle-lights/cmd/lights-control/pkg/light"
	"strconv"
)

func readInput() light.EnvLight {
	// Declare a variable to store the user input
	var input string

	// Print a message to ask the user to enter a number from 1 to 5
	fmt.Println("Receiving information about the light... This is a basic concept demonstration, please type out the intensity of the light in the environment. Write out a single number from 1 to 5, with 5 being the brightest and 1 being the dimmest. Type \"exit\" to exit")

	// Use fmt.Scanln to read a line of text from the standard input
	// Store it in the input variable
	// Handle any error that may occur
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	if input == "exit" {
		return -1 // Return -1 to indicate exit
	}

	// Convert the input from a string to an integer
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	// Convert the integer to an EnvLight value
	envLight := light.EnvLight(num - 1) // Subtract 1 because EnvLight values start from zero

	// Return the input as the output of the function
	return envLight
}
