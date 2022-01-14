package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"sync"
)

type Decoder struct {
	r io.Reader

	init  sync.Once
	depth int
	d     *xml.Decoder
	path  []string
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

const maxDepth = 32

var ErrMaxDepth = fmt.Errorf("max depth reached")

const (
	xmlExchangeFormat = "VymennyFormat"
	xmlData           = "Data"
	xmlAddressPlaces  = "AdresniMista"
	xmlAddressPlace   = "AdresniMisto"
	xmlCities         = "Obce"
	xmlCity           = "Obec"
	xmlStreet         = "Ulice"
)

func (d *Decoder) Next() (interface{}, error) {
	d.init.Do(func() {
		d.d = xml.NewDecoder(d.r)
		d.path = make([]string, maxDepth)
	})

	for {
		token, err := d.d.Token()
		if err != nil {
			return nil, err
		}

		if element, ok := token.(xml.StartElement); ok {
			d.path[d.depth] = element.Name.Local
			d.depth++

			if d.depth == maxDepth {
				return nil, ErrMaxDepth
			}

			if comparePath(d.path, []string{xmlExchangeFormat, xmlData, xmlAddressPlaces, xmlAddressPlace}) {
				ap := AddressPlace{}
				err = d.decodeOne(&ap, &element)

				return ap, err
			}

			if comparePath(d.path, []string{xmlExchangeFormat, xmlData, xmlCities, xmlCity}) {
				c := City{}
				err = d.decodeOne(&c, &element)

				return c, err
			}

			if comparePath(d.path, []string{xmlExchangeFormat, xmlData, xmlStreet, xmlStreet}) {
				c := Street{}
				err = d.decodeOne(&c, &element)

				return c, err
			}
		}

		if _, ok := token.(xml.EndElement); ok {
			d.popPath()
		}
	}
}

// decode one element
func (d *Decoder) decodeOne(target interface{}, element *xml.StartElement) error {
	d.popPath()

	err := d.d.DecodeElement(&target, element)
	if err != nil {
		return err
	}

	return nil
}

func (d *Decoder) popPath() {
	d.depth--
	d.path[d.depth] = ""
}

func comparePath(path []string, target []string) bool {
	if len(target) > len(path) {
		return false
	}

	for i, v := range target {
		if v != path[i] {
			return false
		}
	}

	return true
}
