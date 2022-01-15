package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/encero/ruian_parser/ent"
)

func GetOrCreateAddressPlace(ctx context.Context, etnc *ent.Client, ID int32) (*ent.AddressPlace, error) {
	ap, err := etnc.AddressPlace.Get(ctx, ID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("loading addressPlace: %w", err)
	}

	if ent.IsNotFound(err) {
		apc := etnc.AddressPlace.Create()
		apc.SetID(ID)
		apc.SetNumber(-1)
		apc.SetZip(-1)

		ap, err := apc.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("addressPlace Create:%w", err)
		}

		return ap, nil
	}

	return ap, nil
}

func GetOrCreateStreet(ctx context.Context, entc *ent.Client, ID int32) (*ent.Street, error) {
	s, err := entc.Street.Get(ctx, ID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("loading street: %w", err)
	}

	if ent.IsNotFound(err) {
		sc := entc.Street.Create()
		sc.SetID(ID)
		sc.SetName("")

		s, err := sc.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("street create: %w", err)
		}

		return s, nil
	}

	return s, nil
}

func UpdateAddressPlace(ap *ent.AddressPlace, source AddressPlace) (*ent.AddressPlaceUpdateOne, error) {
	apu := ap.Update()

	number, err := strconv.ParseInt(source.Number, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("addressPlace number not int, %w", err)
	}

	apu.SetNumber(int32(number))

	zip, err := strconv.ParseInt(source.Zip, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("ZIP not int, %w", err)
	}

	apu.SetZip(int32(zip))

	if source.OrientationNumber != "" {
		orientation, err := strconv.ParseInt(source.OrientationNumber, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("orientation number not int, %w", err)
		}

		apu.SetOrientationNumber(int32(orientation))
	}

	apu.SetOrientationNumberLetter(source.OrientationLetter)

	apu.ClearStreets()

	for _, street := range source.Streets {
		streetID, err := strconv.ParseInt(street.Code, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("street code not int, %w", err)
		}

		apu.AddStreetIDs(int32(streetID))
	}

	return apu, nil
}

func MapCity(c *ent.Client, source City) (*ent.CityCreate, error) {
	city := c.City.Create()

	code, err := strconv.ParseInt(source.Code, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("city code not int, %w", err)
	}

	city.SetID(int32(code))
	city.SetName(source.Name)

	return city, nil
}

func UpdateStreet(street *ent.Street, source Street) (*ent.StreetUpdateOne, error) {
	su := street.Update()

	su.SetName(source.Name)
	su.ClearCities()

	for _, city := range source.Cities {
		cityID, err := strconv.ParseInt(city.Code, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("city code not int, %w", err)
		}

		su.AddCityIDs(int32(cityID))
	}

	return su, nil
}
