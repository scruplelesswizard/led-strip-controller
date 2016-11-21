package controller

import "math"

// RGB is a struct representing the Red, Blue and Green byte values
type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

// Add will add the values to the provided RGB struct and return a new RGB struct
func (rgb RGB) add(r, g, b int8) (result RGB) {
	result = rgb
	result.Red += r
	result.Green += g
	result.Blue += b
}

func (a RGB) Off() (black RGB) {
	return RGB{Red: OFF, Green: OFF, Blue: OFF}
}

func (a RGB) maxDifference(b RGB) (maxDiff int8) {
	r, g, b := a.difference(b)
	maxDiff = math.MaxInt8(math.Abs(r), math.Abs(g))
	maxDiff = math.MaxInt8(maxDiff, math.Abs(b))
}

func (a RGB) difference(b RGB) (rDiff, gDiff, bDiff int8) {

	rDiff := a.Red - b.Red
	gDiff := a.Green - b.Green
	bDiff := a.Blue - b.Blue

}
