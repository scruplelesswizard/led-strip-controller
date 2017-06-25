package controller

import (
	"math"
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

func (s *Strip) fade(color HSI, d time.Duration) error {

	stop := s.StopChan()
	defer s.Unsub(stop)

	// calculate step duration and # of steps
	s.OverrideOff(color)

	if s.Color == color {
		return nil
	}

	diff := s.Color.Difference(color)

	steps := math.Max(diff.Hue, math.Max(diff.Intensity, diff.Saturation))

	stepDuration := time.Duration((d.Nanoseconds() / int64(steps))) * time.Nanosecond

	// calculate change per steps

	changeStep := HSI{
		Hue:        diff.Hue / steps,
		Saturation: diff.Saturation / steps,
		Intensity:  diff.Intensity / steps,
	}

	stepCount := int(steps)

	for step := 0; step <= stepCount; step++ {
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

func (s *Strip) FadeBetween(colors []HSI, effectDuration time.Duration) error {

	s.Stop()
	stop := s.StopChan()
	defer s.Unsub(stop)

	for {
		for _, color := range colors {
			select {
			case <-stop:
				return nil
			default:
				err := s.fade(color, effectDuration/2)
				if err != nil {
					return err
				}
			}
		}
	}
}

func (s *Strip) FadeOut(effectDuration time.Duration) error {

	s.Stop()

	err := s.fade(s.Color.Off(), effectDuration)
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
			select {
			case <-stop:
				return nil
			default:
				err := s.SetColor(color)
				if err != nil {
					return err
				}
				time.Sleep(d)
			}
		}
	}

}

func (s *Strip) Flash(c HSI, d time.Duration) error {
	err := s.FlashBetween([]HSI{c, s.Color.Off()}, d)
	if err != nil {
		return err
	}
	return nil
}

func (s *Strip) Pulse(c HSI, d time.Duration) error {

	s.Stop()
	stop := s.StopChan()
	defer s.Unsub(stop)

	maxIntensity := .4
	minIntensity := .3
	stepDistance := 0.001
	sleepCycles := int64((maxIntensity - minIntensity) / stepDistance)
	sleepDuration := time.Duration((d.Nanoseconds() / sleepCycles)) * time.Nanosecond

	var i float64

	for {
		select {
		case <-stop:
			return nil
		default:
			for i = minIntensity; i < maxIntensity; i += stepDistance {
				c.Intensity = i
				err := s.SetColor(c)
				if err != nil {
					return err
				}
				time.Sleep(sleepDuration)
			}
			for i = maxIntensity; i > minIntensity; i -= stepDistance {
				c.Intensity = i
				err := s.SetColor(c)
				if err != nil {
					return err
				}
				time.Sleep(sleepDuration * 2)
			}
		}
	}
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
