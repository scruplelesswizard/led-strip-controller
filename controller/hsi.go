package controller

import "math"

// Based on http://blog.saikoled.com/post/43693602826/why-every-led-light-should-be-using-hsi

type HSI struct {
	Hue        float64
	Saturation float64
	Intensity  float64
}

const HueRed = 0
const HueOrange = 30
const HueYellow = 60
const HueGreen = 120
const HueCyan = 180
const HueBlue = 240
const HueViolet = 260
const HueMagenta = 300
const HuePink = 320

func (h HSI) Add(addend HSI) HSI {

	const HueMin = 0
	const HueMax = 360

	h.Hue += addend.Hue
	h.Saturation = clamp(addend.Saturation+h.Saturation, OFF, ON)
	h.Intensity = clamp(addend.Intensity+h.Intensity, OFF, ON)

	if h.Hue > HueMax {
		h.Hue -= HueMax
	}
	if h.Hue < HueMin {
		h.Hue += HueMax
	}

	return h
}

func (h HSI) Difference(b HSI) HSI {

	hDiff := b.Hue - h.Hue
	sDiff := b.Saturation - h.Saturation
	iDiff := b.Intensity - h.Intensity

	if hDiff > 180 {
		hDiff -= 360
	}

	return HSI{Hue: hDiff, Saturation: sDiff, Intensity: iDiff}
}

func (h HSI) Off() HSI {
	h.Intensity = 0
	return h
}

// ToRGB Converts an HSI color representation to an RGB color representation
func (h HSI) ToRGB() RGB {
	var r, g, b float64
	var H, S, I float64

	H = h.Hue
	S = h.Saturation
	I = h.Intensity

	H = fmod(H, 360)               // cycle H around to 0-360 degrees
	H = 3.14159 * H / float64(180) // Convert to radians.
	// clamp S and I to interval [0,1]
	S = clamp(S, 0, 1)
	I = clamp(I, 0, 1)

	// Math! Thanks in part to Kyle Miller.
	if H < 2.09439 {
		r = I / 3 * (1 + S*math.Cos(H)/math.Cos(1.047196667-H))
		g = I / 3 * (1 + S*(1-math.Cos(H)/math.Cos(1.047196667-H)))
		b = I / 3 * (1 - S)
	} else if H < 4.188787 {
		H = H - 2.09439
		g = I / 3 * (1 + S*math.Cos(H)/math.Cos(1.047196667-H))
		b = I / 3 * (1 + S*(1-math.Cos(H)/math.Cos(1.047196667-H)))
		r = I / 3 * (1 - S)
	} else {
		H = H - 4.188787
		b = I / 3 * (1 + S*math.Cos(H)/math.Cos(1.047196667-H))
		r = I / 3 * (1 + S*(1-math.Cos(H)/math.Cos(1.047196667-H)))
		g = I / 3 * (1 - S)
	}

	return RGB{Red: r, Green: g, Blue: b}

}

func clamp(val, min, max float64) float64 {

	// S = S>0?(S<1?S:1):0;

	if val > min {
		if val > max {
			val = max
		}
	} else {
		val = min
	}
	return val
}

func fmod(x, y float64) float64 {
	// TODO(rsc): Remove manual inlining of IsNaN, IsInf
	// when compiler does it for us.
	if y == 0 || x > math.MaxFloat64 || x < -math.MaxFloat64 || x != x || y != y { // y == 0 || IsInf(x, 0) || IsNaN(x) || IsNan(y)
		return math.NaN()
	}
	if y < 0 {
		y = -y
	}

	yfr, yexp := math.Frexp(y)
	sign := false
	r := x
	if x < 0 {
		r = -x
		sign = true
	}

	for r >= y {
		rfr, rexp := math.Frexp(r)
		if rfr < yfr {
			rexp = rexp - 1
		}
		r = r - math.Ldexp(y, rexp-yexp)
	}
	if sign {
		r = -r
	}
	return r
}
