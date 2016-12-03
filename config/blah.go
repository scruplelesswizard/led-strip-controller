package config

import (
	"log"

	"github.com/chaosaffe/led-strip-controller/controller"
)

func BuildStrips() controller.Strips {

	var s controller.Strips

	path := "./strips-example.yaml"
	sD, err := LoadStripsDefFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, stripDef := range sD.Strips {
		strip, err := controller.NewStrip(stripDef.Name, stripDef.RedPin, stripDef.GreenPin, stripDef.BluePin)
		if err != nil {
			log.Fatal(err)
		}
		s.AddStrip(strip)
	}

	return s

}
