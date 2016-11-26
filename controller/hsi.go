package controller

import "math"

// Based on http://blog.saikoled.com/post/43693602826/why-every-led-light-should-be-using-hsi

type HSI struct {
	Hue        float64
	Saturation float64
	Intensity  float64
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

	return RGB{Red: float32(r), Green: float32(g), Blue: float32(b)}

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
