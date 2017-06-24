package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chaosaffe/led-strip-controller/controller"

	"gopkg.in/yaml.v2"
)

type StripDef struct {
	Name     string `yaml:"name"`
	RedPin   string `yaml:"redPin"`
	GreenPin string `yaml:"greenPin"`
	BluePin  string `yaml:"bluePin"`
}

type StripsDef struct {
	Strips []*StripDef `yaml:"strips"`
}

func LoadStripsDefFromFile(path string) (*StripsDef, error) {

	var sD *StripsDef

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("Could not stat %s: %s", path, err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Unable to read %s: %s", path, err)
		return sD, err
	}

	err = yaml.Unmarshal(data, &sD)
	if err != nil {
		err = fmt.Errorf("Unable to unmarshal %s to YAML: %s", path, err)
		return sD, err
	}

	return sD, nil

}

func BuildStrips(path string) controller.Strips {

	var s controller.Strips

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
		log.Printf("Created Strip %s", strip.Name)
	}

	return s

}
