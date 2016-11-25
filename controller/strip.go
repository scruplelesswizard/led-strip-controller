package controller

import (
	"fmt"
	"time"
)

type Strip struct {
	Color RGB
	rPin  pwmPin
	gPin  pwmPin
	bPin  pwmPin
}

var pinRed = "17"
var pinGreen = "22"
var pinBlue = "24"

func NewStrip() (*Strip, error) {

	rPin, err := newPWMPin(pinRed)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize pinRed: %v", err)
	}

	gPin, err := newPWMPin(pinGreen)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize pinGreen: %v", err)
	}

	bPin, err := newPWMPin(pinBlue)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize pinBlue: %v", err)
	}

	s := Strip{
		rPin: *rPin,
		gPin: *gPin,
		bPin: *bPin,
	}
	s.Color = RGB{Red: OFF, Green: OFF, Blue: OFF}
	err = s.Off()
	if err != nil {
		return nil, err
	}

	return &s, nil

}

func (s Strip) SetColor(color RGB) error {
	s.Color = color
	return s.setPins()
}

func (s Strip) setPins() error {
	if err := s.rPin.Set(s.Color.Red); err != nil {
		return fmt.Errorf("Failed to set redPin: %v", err)
	}

	if err := s.gPin.Set(s.Color.Green); err != nil {
		return fmt.Errorf("Failed to set greenPin: %v", err)
	}

	if err := s.bPin.Set(s.Color.Blue); err != nil {
		return fmt.Errorf("Failed to set bluePin: %v", err)
	}

	return nil
}

func (s Strip) TestStrip() error {

	const testSeparationDuration = 250

	println("Testing LED Strip")

	var test TestPatterns
	test.Default()

	for _, v := range test {
		fmt.Printf("Starting Test %s\n", v.Name)
		err := s.SetColor(v.Color)
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(v.Duration) * time.Millisecond)
		err = s.Off()
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(testSeparationDuration) * time.Millisecond)
	}

	return nil
}
