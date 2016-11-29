package controller

import (
	"fmt"
	"time"
)

type Strip struct {
	Color HSI
	rPin  pwmPin
	gPin  pwmPin
	bPin  pwmPin
}

func NewStrip(rPinNumber, gPinNumber, bPinNumber string) (*Strip, error) {

	rPin, err := newPWMPin(rPinNumber)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize rPin: %v", err)
	}

	gPin, err := newPWMPin(gPinNumber)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize gPin: %v", err)
	}

	bPin, err := newPWMPin(bPinNumber)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize bPin: %v", err)
	}

	s := Strip{
		rPin: *rPin,
		gPin: *gPin,
		bPin: *bPin,
	}

	s.Off()

	return &s, nil

}

func (s *Strip) SetColor(color HSI) {
	s.Color = color
	s.setPins()
}

func (s *Strip) setPins() {

	color := s.Color.ToRGB()

	if err := s.rPin.Set(color.Red); err != nil {
		panic(err)
	}

	if err := s.gPin.Set(color.Green); err != nil {
		panic(err)
	}

	if err := s.bPin.Set(color.Blue); err != nil {
		panic(err)
	}
}

func (s *Strip) TestStrip() {

	const testSeparationDuration = 250

	println("Testing LED Strip")

	var test TestPatterns
	test.Default()

	for _, v := range test {
		fmt.Printf("Starting Test %s\n", v.Name)
		s.SetColor(v.Color)
		time.Sleep(time.Duration(v.Duration) * time.Millisecond)
		s.Off()
		time.Sleep(time.Duration(testSeparationDuration) * time.Millisecond)

	}

}

func (s *Strip) Off() {
	color := s.Color
	color.Intensity = 0
	s.SetColor(color)
}
