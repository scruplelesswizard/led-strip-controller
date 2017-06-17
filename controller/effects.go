package controller

import (
	"time"

	"github.com/cskr/pubsub"
)

var ps *pubsub.PubSub

func init() {
	ps = pubsub.New(1)
}

func (s *Strip) Rotate() error {

	s.Stop()
	stop := s.StopChan()
	defer s.Unsub(stop)

	d, _ := time.ParseDuration("50ms")
	err := s.SetColor(HSI{Hue: s.Color.Hue, Saturation: 1, Intensity: 1})
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

func (s *Strip) Fade(color HSI, effectDuration time.Duration) error {

	s.Stop()
	return s.fade(color, effectDuration)
}

func (s *Strip) fade(color HSI, effectDuration time.Duration) error {

	stop := s.StopChan()
	defer s.Unsub(stop)

	// calculate step duration and # of steps
	stepDuration := time.Duration(20) * time.Millisecond
	steps := float64(effectDuration / stepDuration)

	// calculate differences

	s.OverrideOff(color)

	diff := s.Color.Difference(color)

	// calculate change per steps

	changeStep := HSI{
		Hue:        diff.Hue / steps,
		Saturation: diff.Saturation / steps,
		Intensity:  diff.Intensity / steps,
	}

	for step := 0; step <= int(steps); step++ {
		select {
		case <-stop:
			return nil
		default:
			err := s.SetColor(s.Color.Add(changeStep))
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

func (s *Strip) FadeBetween(colors []HSI, duration time.Duration) error {

	s.Stop()
	stop := s.StopChan()
	defer s.Unsub(stop)

	for {
		for _, color := range colors {
			switch {
			case <-stop:
				return nil
			default:
				err := s.Fade(color, duration/2)
				if err != nil {
					return err
				}
			}
		}
	}
}

func (s *Strip) FadeOut(duration time.Duration) error {

	err := s.Fade(s.Color.Off(), duration)
	if err != nil {
		return err
	}
	return nil
}

func (s *Strip) FlashBetween(c []HSI, d time.Duration) error {

	s.Stop()
	stop := s.StopChan()
	defer s.Unsub(stop)

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

func (s *Strip) Pulse(c HSI, d time.Duration) error {
	var intensity float64
	color := HSI{Hue: 0, Saturation: 1, Intensity: .5}
	for {
		for intensity = .3; intensity < .4; intensity = intensity + 0.001 {
			color.Intensity = intensity
			err := s.SetColor(color)
			if err != nil {
				return err
			}
			time.Sleep(d)
		}
		for intensity = .4; intensity > .3; intensity = intensity - 0.001 {
			color.Intensity = intensity
			err := s.SetColor(color)
			if err != nil {
				return err
			}
			time.Sleep(d * 2)
		}
	}
	return nil
}

func (s *Strip) Off() error {
	s.Stop()
	color := s.Color
	color.Intensity = 0
	return s.SetColor(color)
}

func (s *Strip) OverrideOff(color HSI) {
	if s.Color.Intensity == 0 {
		s.Color.Hue = color.Hue
		s.Color.Saturation = color.Saturation
	}
}
