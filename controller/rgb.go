package controller

import "math"

// RGB is a struct representing the Red, Blue and Green byte values
type RGB struct {
	Red   float32
	Green float32
	Blue  float32
}

// Add will add the values to the provided RGB struct and return a new RGB struct
func (rgb RGB) add(r, g, b float32) (result RGB) {
	result = rgb
	result.Red = percentAdd(result.Red, r)
	result.Green = percentAdd(result.Green, g)
	result.Blue = percentAdd(result.Blue, b)
	return
}

func percentAdd(a, b float32) float32 {

	result := a + b
	if result < OFF {
		result = OFF
	}
	if result > ON {
		result = ON
	}

	return result

}

// Off returns an RGB object that represents the off state for all LED's
func (a RGB) Off() (black RGB) {
	return RGB{Red: OFF, Green: OFF, Blue: OFF}
}

func (a RGB) maxDifference(other RGB) (maxDiff float32) {
	r, g, b := a.difference(other)
	maxDiff = max(r, g, b)
	return
}

func max(a, b, c float32) float32 {

	aAbs := math.Abs(float64(a))
	bAbs := math.Abs(float64(b))
	cAbs := math.Abs(float64(c))

	var retVal float64

	if aAbs > bAbs {
		retVal = aAbs
	} else {
		retVal = bAbs
	}

	if retVal < cAbs {
		retVal = cAbs
	}

	return float32(retVal)

}

func (a RGB) difference(b RGB) (rDiff, gDiff, bDiff float32) {

	rDiff = a.Red - b.Red
	gDiff = a.Green - b.Green
	bDiff = a.Blue - b.Blue
	return

}
