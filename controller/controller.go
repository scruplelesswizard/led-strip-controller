package controller

import (
	"fmt"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

type Strip struct {
	Color RGB
	rPin  embd.PWMPin
	gPin  embd.PWMPin
	bPin  embd.PWMPin
}

var pinRed = "P1_36"
var pinGreen = "P1_38"
var pinBlue = "P1_40"

func NewStrip() (strip Strip) {

	err := embd.InitGPIO()
	if err != nil {
		panic(err)
	}

	rPin, err := embd.NewPWMPin(pinRed)
	if err != nil {
		panic(err)
	}

	gPin, err := embd.NewPWMPin(pinGreen)
	if err != nil {
		panic(err)
	}

	bPin, err := embd.NewPWMPin(pinBlue)
	if err != nil {
		panic(err)
	}

	off := RGB{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	s := Strip{
		Color: off,
		rPin:  rPin,
		gPin:  gPin,
		bPin:  bPin,
	}

	s.setPins()

	return s

}

func (s *Strip) SetColor(color RGB) {
	s.Color = color
	s.setPins()
}

func (s *Strip) setPins() {
	if err := s.rPin.SetAnalog(s.Color.Red); err != nil {
		panic(err)
	}

	if err := s.gPin.SetAnalog(s.Color.Green); err != nil {
		panic(err)
	}

	if err := s.bPin.SetAnalog(s.Color.Blue); err != nil {
		panic(err)
	}
}

func (s *Strip) Close() {
	s.rPin.Close()
	s.gPin.Close()
	s.bPin.Close()
	embd.CloseGPIO()
}

func (s *Strip) TestStrip() {

	const testSeparationDuration = 250

	println("Starting Test")

	var test TestPatterns
	test.Default()

	for _, v := range test {
		fmt.Printf("Starting Test %s", v.Name)
		s.SetColor(v.Color)
		time.Sleep(time.Duration(v.Duration) * time.Millisecond)
		s.Off()
		time.Sleep(time.Duration(testSeparationDuration) * time.Millisecond)

	}

}

func (s *Strip) Off() {
	s.SetColor(RGB{Red: 0, Green: 0, Blue: 0})
}
