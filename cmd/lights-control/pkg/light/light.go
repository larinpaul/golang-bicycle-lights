package light

type EnvLight int // Define an enumeration type for environment light intensity

const (
	ExtremelyDark EnvLight = iota // Assign values starting from zero
	VeryDark
	Dark
	ModerateBrightness
	Bright
)

// Check if a function that takes an int representing the environment light intensity
// and returns a string with a message indicating the energy consumption level of the bicycle lights
func Check(envLight EnvLight) string {
	var msg string

	switch envLight {
	case Bright:
		msg = "Bright environment ☀️ \nNo energy consumption 🔋"
	case ModerateBrightness:
		msg = "Moderate brightness environment 🌥️ \nMedium energy consumption 💡⚡"
	case Dark:
		msg = "Dark environment 🌃 \nHigh energy consumption 🔦⚡⚡"
	case VeryDark, ExtremelyDark:
		msg = "Extremely dark environment 🦇\nMaximum energy consumption 🔦🔦⚡⚡⚡"
	default:
		msg = "Invalid light environment information ⚠️"
	}

	return msg
}
