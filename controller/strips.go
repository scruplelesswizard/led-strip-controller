package controller

import "fmt"

type Strips []*Strip

func (sCol *Strips) AddStrip(s *Strip) {
	*sCol = append(*sCol, s)
}

func (sCol *Strips) GetStrip(name string) (*Strip, error) {

	for _, s := range *sCol {
		if s.Name == name {
			return s, nil
		}
	}

	return nil, fmt.Errorf("Strip with name %s was not found", name)
}

func (sCol *Strips) AllOff() {

	for _, s := range *sCol {
		s.Off()
	}

}
