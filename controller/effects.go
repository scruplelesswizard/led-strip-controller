package controller

import "time"

func (s *Strip) Fade(color HSI, duration time.Duration) {

	// calculate step duration and # of steps
	stepDuration := time.Duration(20) * time.Millisecond
	steps := float64(duration / stepDuration)

	// calculate differences

	hsiDiff := s.Color.Difference(color)

	// calculate change per steps

	hsiStep := HSI{
		Hue:        hsiDiff.Hue / steps,
		Saturation: hsiDiff.Saturation / steps,
		Intensity:  hsiDiff.Intensity / steps,
	}

	for step := 0; step <= int(steps); step++ {
		s.SetColor(s.Color.Add(hsiStep))
		time.Sleep(stepDuration)
	}

	// clean up floats
	s.SetColor(color)

}

func (s *Strip) FadeBetween(a, b HSI, duration time.Duration) {

	s.Fade(a, duration/2)
	// HACK: This will block. Use channel to break when required
	for {
		s.Fade(b, duration/2)
		s.Fade(a, duration/2)
	}

}

func (s *Strip) FadeOut(duration time.Duration) {
	s.Fade(s.Color.Off(), duration)
}

func (s *Strip) FlashBetween(c []HSI, d time.Duration) {

	// HACK: This will block. Use channel to break when required
	for {
		for _, color := range c {
			s.SetColor(color)
			time.Sleep(d)
		}
	}
}

func (s *Strip) Flash(c HSI, d time.Duration) {
	s.FlashBetween([]HSI{c, s.Color.Off()}, d)
}
