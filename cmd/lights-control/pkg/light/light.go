package light

// Check is a function that takes a string representing the environment light intensity
// and returns a string with a message indicating the energy consumption level of the bicycle lights
func Check(envLight string) string {
	// Declare a variable to store the message
	var msg string

	// Use a switch statement to check the value of envLight
	switch envLight {
	// If envLight is "intensive", set msg to "Bright environment ☀️ \n No energy consumption"
	case "intensive":
		msg = "Bright environment ☀️ \nNo energy consumption 🔋"
	case "medium":
		msg = "Moderate brightness environment 🌥️ \nMedium energy consumption 💡⚡"
	case "dark":
		msg = "Dark environment 🌃 \nHigh energy consumption 🔦⚡⚡"
	case "ultra-dark":
		msg = "Extremely dark environment 🦇\nMaximum energy consumption 🔦🔦⚡⚡⚡"
	// If envLight is anything else, set msg to "Invalid input"
	default:
		msg = "Invalid light environment info"
	}

	// Return msg as the output of the function
	return msg
}
