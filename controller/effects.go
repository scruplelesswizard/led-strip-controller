package controller

import "time"

func (s *Strip) Rotate(stop chan bool) error {

	d, _ := time.ParseDuration("50ms")
	err := s.SetColor(HSI{Hue: 0, Saturation: 1, Intensity: 1})
	if err != nil {
		return err
	}
	for {
		select {
		case <-stop:
			return nil
		default:
			err = s.SetColor(s.Color.Add(HSI{Hue: 1, Saturation: 0, Intensity: 0}))
			if err != nil {
				return err
			}
			time.Sleep(d)
		}

	}
}

func (s *Strip) Fade(color HSI, duration time.Duration, stop chan bool) error {

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
			return nil
		default:
			err := s.SetColor(s.Color.Add(hsiStep))
			if err != nil {
				return err
			}
			time.Sleep(stepDuration)
		}

	}

	// clean up floats
	err := s.SetColor(color)
	if err != nil {
		return err
	}

	return nil

}

func (s *Strip) FadeBetween(a, b HSI, duration time.Duration, stop chan bool) error {

	err := s.Fade(a, duration/2, stop)
	if err != nil {
		return err
	}

	if <-stop {
		return nil
	}

	for {
		switch {
		case <-stop:
			return nil
		default:
			err = s.Fade(b, duration/2, stop)
			if err != nil {
				return err
			}
			err = s.Fade(a, duration/2, stop)
			if err != nil {
				return err
			}
		}
	}
}

func (s *Strip) FadeOut(duration time.Duration, stop chan bool) error {
	err := s.Fade(s.Color.Off(), duration, stop)
	if err != nil {
		return err
	}
	return nil
}

func (s *Strip) FlashBetween(c []HSI, d time.Duration) error {

	// HACK: This will block. Use channel to break when required
	for {
		for _, color := range c {
			err := s.SetColor(color)
			if err != nil {
				return err
			}
			time.Sleep(d)
		}
	}
	return nil
}

func (s *Strip) Flash(c HSI, d time.Duration) error {
	err := s.FlashBetween([]HSI{c, s.Color.Off()}, d)
	if err != nil {
		return err
	}
	return nil
}

func (s *Strip) Pulse(c HSI, d time.Duration) {
	var intensity float64
	color := HSI{Hue: 0, Saturation: 1, Intensity: .5}
	for {
		for intensity = .3; intensity < .4; intensity = intensity + 0.001 {
			color.Intensity = intensity
			s.SetColor(color)
			time.Sleep(d)
		}
		for intensity = .4; intensity > .3; intensity = intensity - 0.001 {
			color.Intensity = intensity
			s.SetColor(color)
			time.Sleep(d * 2)
		}
	}
}
