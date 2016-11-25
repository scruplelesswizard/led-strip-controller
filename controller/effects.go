package controller

import "time"

// Off sets all LED's on the strip to off
func (s *Strip) Off() error {
	return s.SetColor(s.Color.Off())
}

func (s *Strip) Fade(color RGB, duration time.Duration) error {

	// calculate step duration and # of steps
	stepDuration := time.Duration(20) * time.Millisecond
	steps := float32(duration / stepDuration)

	// calculate differences

	// TODO: Fades should be done in the HSV space for more appeasing gradients

	rDiff, gDiff, bDiff := color.difference(s.Color)

	// calculate change per steps

	rStep := rDiff / steps
	gStep := gDiff / steps
	bStep := bDiff / steps

	for step := 0; step <= int(steps); step++ {
		err := s.SetColor(s.Color.add(rStep, gStep, bStep))
		if err != nil {
			return err
		}
		time.Sleep(stepDuration)
	}

	return nil
}

func (s *Strip) FadeBetween(a, b RGB, duration time.Duration) error {

	err := s.Fade(a, duration/2)
	if err != nil {
		return err
	}
	// HACK: This will block. Consider using channel to break when required
	for {
		err = s.Fade(a, duration/2)
		if err != nil {
			return err
		}
		err = s.Fade(b, duration/2)
		if err != nil {
			return err
		}
	}

}

func (s *Strip) FadeOut(duration time.Duration) error {
	return s.Fade(s.Color.Off(), duration)
}

func (s *Strip) FlashBetween(c []RGB, d time.Duration) error {

	// HACK: This will block. Consider using channel to break when required
	for {
		for _, color := range c {
			err := s.SetColor(color)
			if err != nil {
				return err
			}
			time.Sleep(d)
		}
	}
}

func (s *Strip) Flash(c RGB, d time.Duration) error {
	return s.FlashBetween([]RGB{c, c.Off()}, d)
}
