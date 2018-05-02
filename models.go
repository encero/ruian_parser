package main

type AddressPlace struct {
	Code              string      `xml:"Kod"`
	Number            string      `xml:"CisloDomovni"`
	OrientationNumber string      `xml:"CisloOrientacni"`
	OrientationLetter string      `xml:"CisloOrientacniPismeno"`
	Zip               string      `xml:"Psc"`
	Streets           []StreetRef `xml:"Ulice"`
}

type City struct {
	Code          string      `xml:"Kod"`
	Name          string      `xml:"Nazev"`
	Nonvalid      bool        `xml:"Nespravny"`
	StatutaryCode int         `xml:"StatusKod"`
	Precinct      PrecinctRef `xml:"Okres"`
}

type CityRef struct {
	Code string `xml:"Kod"`
}

type Street struct {
	Code   string    `xml:"Kod"`
	Name   string    `xml:"Nazev"`
	Cities []CityRef `xml:"Obec"`
}

type StreetRef struct {
	Code string `xml:"Kod"`
}

type PrecinctRef struct {
	Code string `xml:"Kod"`
}
