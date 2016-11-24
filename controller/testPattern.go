package controller

// TestPattern is a struct representing a pattern used to test the strip
type TestPattern struct {
	Name  string
	Color RGB
	// Duration is the duration the pattern should be held in ms
	Duration int
}

// TestPatterns is a collection of TestPattern
type TestPatterns []TestPattern

// ON represents the fully on state for the LED's
const ON = 1

// OFF represents the fully off state for the LED's
const OFF = 0

// Default creates the default test patterns of Full Red, Full Green, Full Blue, Full White, and Half White
func (t *TestPatterns) Default() {

	redFull := TestPattern{
		Name:     "Full Red",
		Color:    RGB{Red: ON, Green: OFF, Blue: OFF},
		Duration: 1000,
	}

	greenFull := TestPattern{
		Name:     "Full Green",
		Color:    RGB{Red: OFF, Green: ON, Blue: OFF},
		Duration: 1000,
	}

	blueFull := TestPattern{
		Name:     "Full Blue",
		Color:    RGB{Red: OFF, Green: OFF, Blue: ON},
		Duration: 1000,
	}

	whiteFull := TestPattern{
		Name:     "Full All",
		Color:    RGB{Red: ON, Green: ON, Blue: ON},
		Duration: 1000,
	}

	whiteHalf := TestPattern{
		Name:     "Half All",
		Color:    RGB{Red: ON / 2, Green: ON / 2, Blue: ON / 2},
		Duration: 1000,
	}

	*t = append(*t, redFull, greenFull, blueFull, whiteFull, whiteHalf)

}
