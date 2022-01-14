package main

import (
	"context"
	"github.com/encero/ruian_parser/ent"
	"github.com/encero/ruian_parser/ent/enttest"
	"github.com/encero/ruian_parser/ent/migrate"
	is_ "github.com/matryer/is"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func testDb(t *testing.T) *ent.Client {
	return enttest.Open(
		t,
		"sqlite3",
		"file:ent?mode=memory&_fk=1&cache=shared",
		enttest.WithMigrateOptions(
			migrate.WithForeignKeys(false),
		),
	)
}

func testCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond) //nolint:govet
	return ctx
}

func TestGetOrCreateAddressPlace(t *testing.T) {
    is := is_.New(t)

    ec := testDb(t)
    defer ec.Close()

    _, err := ec.
    AddressPlace.
    Create().
    SetID(1).
    SetNumber(12).
    SetZip(1).
    Save(testCtx())
    is.NoErr(err) // create test address place for retrieval

    ap, err := GetOrCreateAddressPlace(testCtx(), ec, 1)
    is.NoErr(err) // retrieve address place by id

    is.Equal(ap.Number, int32(12))

    ap, err = GetOrCreateAddressPlace(testCtx(), ec, 2)
    is.NoErr(err) // should create new address place without error

    is.Equal(ap.ID, int32(2))
}

func TestUpdateAddressPlace(t *testing.T) {
	is := is_.New(t)

	ec := testDb(t)
	defer ec.Close()

	_, err := ec.Street.Create().SetID(11).SetName("").Save(testCtx())
	is.NoErr(err)

    ap, err := GetOrCreateAddressPlace(testCtx(), ec, 1)
    is.NoErr(err) // prepare address place for update

	aps := AddressPlace{
		Code:              "1",
		Number:            "2",
		OrientationNumber: "3",
		OrientationLetter: "4",
		Zip:               "5",
		Streets: []StreetRef{
			{
				Code: "11",
			},
		},
	}

	apu, err := UpdateAddressPlace(ap, aps)
	is.NoErr(err)

	_, err = apu.Save(testCtx())
	is.NoErr(err)

	apDb, err := ec.AddressPlace.Get(testCtx(), 1)
	is.NoErr(err)

	is.Equal(apDb.ID, int32(1))
	is.Equal(apDb.Number, int32(2))
	is.Equal(apDb.OrientationNumber, int32(3))
	is.Equal(apDb.OrientationNumberLetter, "4")
	is.Equal(apDb.Zip, int32(5))
	is.Equal(apDb.QueryStreets().OnlyX(testCtx()).ID, int32(11))
}

func TestMapCity(t *testing.T) {
    is := is_.New(t)

    entc := testDb(t)
    defer entc.Close()

    c := City{
        Code: "1",
        Name: "2",
    }

    cu, err := MapCity(entc, c)
    is.NoErr(err)

    _, err = cu.Save(testCtx())
    is.NoErr(err)

    cDb, err := entc.City.Get(testCtx(), 1)
    is.NoErr(err)

    is.Equal(cDb.ID, int32(1))
    is.Equal(cDb.Name, "2")
}

func TestGetOrCreateStreet(t *testing.T) {
	is := is_.New(t)

    ec := testDb(t)
    defer ec.Close()

    ec.City.Create().SetID(11).SetName("").SaveX(testCtx())

    ss := Street{
        Code: "1",
        Name: "2",
        Cities: []CityRef{
            {
                Code: "11",
            },
        },
    }

    s, err := GetOrCreateStreet(testCtx(), ec, 1)
    is.NoErr(err)

    su, err := UpdateStreet(s, ss)
    is.NoErr(err)

    _, err = su.Save(testCtx())
    is.NoErr(err)

    sDb, err := ec.Street.Get(testCtx(), 1)
    is.NoErr(err)

    is.Equal(sDb.ID, int32(1))
    is.Equal(sDb.Name, "2")
    is.Equal(sDb.QueryCities().OnlyX(testCtx()).ID, int32(11))
}


