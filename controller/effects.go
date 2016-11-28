package controller

import "time"

func (s *Strip) Rotate(stop chan bool) {

	d, _ := time.ParseDuration("50ms")
	s.SetColor(HSI{Hue: 0, Saturation: 1, Intensity: 1})
	for {
		select {
		case <-stop:
			return
		default:
			s.SetColor(s.Color.Add(HSI{Hue: 1, Saturation: 0, Intensity: 0}))
			time.Sleep(d)
		}

	}
}

func (s *Strip) Fade(color HSI, duration time.Duration, stop chan bool) {

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
		select {
		case <-stop:
			return
		default:
			s.SetColor(s.Color.Add(hsiStep))
			time.Sleep(stepDuration)
		}

	}

	// clean up floats
	s.SetColor(color)

}

func (s *Strip) FadeBetween(a, b HSI, duration time.Duration, stop chan bool) {

	s.Fade(a, duration/2, stop)
	if <-stop {
		return
	}
	for {
		switch {
		case <-stop:
			return
		default:
			s.Fade(b, duration/2, stop)
			s.Fade(a, duration/2, stop)
		}

	}

}

func (s *Strip) FadeOut(duration time.Duration, stop chan bool) {
	s.Fade(s.Color.Off(), duration, stop)
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
