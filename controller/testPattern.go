package controller

// TestPattern is a struct representing a pattern used to test the strip
type TestPattern struct {
	Name  string
	Color HSI
	// Duration is the duration the pattern should be held in ms
	Duration int
}

// TestPatterns is a collection of TestPattern
type TestPatterns []TestPattern

// Default creates the default test patterns of Full Red, Full Green, Full Blue, Full White, and Half White
func (t *TestPatterns) Default() {

	redFull := TestPattern{
		Name:     "Full Red",
		Color:    HSI{Hue: HueRed, Saturation: 1, Intensity: 1},
		Duration: 1000,
	}

	greenFull := TestPattern{
		Name:     "Full Green",
		Color:    HSI{Hue: HueGreen, Saturation: 1, Intensity: 1},
		Duration: 1000,
	}

	blueFull := TestPattern{
		Name:     "Full Blue",
		Color:    HSI{Hue: HueBlue, Saturation: 1, Intensity: 1},
		Duration: 1000,
	}

	whiteFull := TestPattern{
		Name:     "Full All",
		Color:    HSI{Hue: 0, Saturation: 0, Intensity: 1},
		Duration: 1000,
	}

	whiteHalf := TestPattern{
		Name:     "Half All",
		Color:    HSI{Hue: 0, Saturation: 0, Intensity: .5},
		Duration: 1000,
	}

	*t = append(*t, redFull, greenFull, blueFull, whiteFull, whiteHalf)

}
