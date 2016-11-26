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

var pinRed = "17"
var pinGreen = "22"
var pinBlue = "24"

func NewStrip() (strip Strip) {

	rPin, err := newPWMPin(pinRed)
	if err != nil {
		panic(err)
	}

	gPin, err := newPWMPin(pinGreen)
	if err != nil {
		panic(err)
	}

	bPin, err := newPWMPin(pinBlue)
	if err != nil {
		panic(err)
	}

	s := Strip{
		rPin: rPin,
		gPin: gPin,
		bPin: bPin,
	}

	s.Off()

	return s

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
