package main

import (
	"os"
	"testing"

	"github.com/matryer/is"
)

var (
	sampleAddressPlace = AddressPlace{
		Code:   "21690278",
		Number: "1",
		Zip:    "11900",
		Streets: []StreetRef{{
			Code: "482536",
		}},
	}

	sampleStreet = Street{
		Code: "782335",
		Name: "1. máje",
		Cities: []CityRef{
			{
				Code: "538728",
			},
		},
	}

	sampleCity = City{
		Code:          "554782",
		Name:          "Praha",
		Nonvalid:      false,
		StatutaryCode: 5,
		Precinct: PrecinctRef{
			Code: "3100",
		},
	}
)

func TestDecoder(t *testing.T) {
	var is = is.New(t)

	f, err := os.Open("testdata/sample.xml")
	is.NoErr(err) // os.Open should always succeed

	defer f.Close()

	decoder := &Decoder{
		r: f,
	}

	decodeAndAssert(is, decoder, sampleAddressPlace)

	decodeAndAssert(is, decoder, sampleCity)

	decodeAndAssert(is, decoder, sampleStreet)
}

func decodeAndAssert(is *is.I, d *Decoder, expected interface{}) {
	is.Helper()

	next, err := d.Next()
	is.NoErr(err) // Next should not return error

	is.Equal(next, expected)
}
