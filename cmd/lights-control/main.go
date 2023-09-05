package main

import (
	"fmt"
	"golang-bicycle-lights/cmd/lights-control/pkg/light"
)

func main() {
	// Assume that we have some way to measure the environment light intensity
	// For simplicity, we will use a variable with a fixed value
	// We will later replace this without own logic to get the actual value
	envLight := "intensive"
	envLight = "medium"
	envLight = "dark"
	envLight = "ultra-dark"

	// Call the light.Check function from the light package
	// Pass the envLight variable as an argument
	// Print the returned message to the standard output
	fmt.Println(light.Check(envLight))
}
