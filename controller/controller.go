package controller

import (
	"fmt"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

type Strip struct {
	Color RGB
	rPin  embd.PWMPin
	gPin  embd.PWMPin
	bPin  embd.PWMPin
}

const pinRed = "17"
const pinGreen = "27"
const pinBlue = "22"

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

	return s, nil

}

func (s *Strip) SetColor(color RGB) {
	s.Color = color
	s.setPins()
}

func (s *Strip) setPins() {
	if err := rPin.SetAnalog(s.Color.Red); err != nil {
		panic(err)
	}

	if err := gPin.SetAnalog(s.Color.Green); err != nil {
		panic(err)
	}

	if err := bPin.SetAnalog(s.Color.Blue); err != nil {
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

	test := TestPatterns
	test.Default()

	for _, v := range test {
		fmt.Printf("Starting Test %s", v.Name)
		s.SetColor(v.Color)
		time.sleep(v.Duration * time.Millisecond)
		s.Off()
		time.sleep(testSeparationDuration * time.Millisecond)

	}

}

func (s *Strip) Off() {
	s.SetColor(RGB{Red: 0, Green: 0, Blue: 0})
}
