package main

import (
	"context"
	"testing"

	"github.com/encero/ruian_parser/ent"
	is_ "github.com/matryer/is"
)

func TestStorageStore(t *testing.T) {
	is := is_.New(t)
	storage := NewStorage(&ent.Client{})

	in := &City{
		Code: "xxx",
		Name: "yyy",
	}

	called := false

	storage.AddHandler(func(_ context.Context, _ *ent.Client, item interface{}) (bool, error) {
		called = true

		is.Equal(item, in) // should get same item

		return true, nil
	})

	err := storage.Store(context.Background(), in)
	is.NoErr(err) // store should not return error

	is.True(called) // handler should be called
}

func TestStorageStore_CallsAllHandlers(t *testing.T) {
	is := is_.New(t)
	storage := NewStorage(&ent.Client{})

	in := &City{
		Code: "xxx",
		Name: "yyy",
	}

	callCount := 0
	handler := func(_ context.Context, _ *ent.Client, item interface{}) (bool, error) {
		callCount++

		is.Equal(item, in) // should get same item

		return false, nil
	}

	for i := 0; i < 10; i++ {
		storage.AddHandler(handler)
	}

	err := storage.Store(context.Background(), in)
	is.Equal(err.Error(), "item not supported") // store should return error

	is.Equal(callCount, 10) // handler should be called 10 times
}

func TestStoreCity(t *testing.T) {
	entc := testDb(t)

	is := is_.New(t)
	storage := NewStorage(entc)
	storage.AddHandler(StoreCity)

	in := City{
		Code: "1234",
		Name: "Test City",
	}

	err := storage.Store(testCtx(), in)
	is.NoErr(err) // store should not return error

	stored, err := entc.City.Get(testCtx(), 1234)
	is.NoErr(err)                      // get should not return error
	is.Equal(stored.ID, int32(1234))   // should get same item
	is.Equal(stored.Name, "Test City") // should get same item
}

func TestStoreStreet(t *testing.T) {
	entc := testDb(t)

	is := is_.New(t)
	storage := NewStorage(entc)
	storage.AddHandler(StoreStreet)

	entc.City.Create().SetID(11).SetName("").SaveX(testCtx())

	in := Street{
		Code: "1234",
		Name: "Test Street",
		Cities: []CityRef{
			{
				Code: "11",
			},
		},
	}

	err := storage.Store(testCtx(), in)
	is.NoErr(err) // store should not return error

	stored, err := entc.Street.Get(testCtx(), 1234)
	is.NoErr(err)                                                 // get should not return error
	is.Equal(stored.ID, int32(1234))                              // should get same item
	is.Equal(stored.Name, "Test Street")                          // should get same item
	is.Equal(stored.QueryCities().OnlyX(testCtx()).ID, int32(11)) // should get same item
}

func TestStoreAddressPlace(t *testing.T) {
	entc := testDb(t)

	is := is_.New(t)
	storage := NewStorage(entc)
	storage.AddHandler(StoreAddressPlace)

	entc.Street.Create().SetID(12).SetName("").SaveX(testCtx())

	in := AddressPlace{
		Code:              "1",
		Number:            "2",
		OrientationNumber: "3",
		OrientationLetter: "A",
		Zip:               "12345",
		Streets: []StreetRef{
			{
				Code: "12",
			},
		},
	}

	asserts := func() {
		stored, err := entc.AddressPlace.Get(testCtx(), 1)
		is.NoErr(err)                                                  // get should not return error
		is.Equal(stored.ID, int32(1))                                  // should get same item
		is.Equal(stored.Number, int32(2))                              // should get same item
		is.Equal(stored.OrientationNumber, int32(3))                   // should get same item
		is.Equal(stored.OrientationNumberLetter, "A")                  // should get same item
		is.Equal(stored.Zip, int32(12345))                             // should get same item
		is.Equal(stored.QueryStreets().OnlyX(testCtx()).ID, int32(12)) // should get same item
	}

	// create
	err := storage.Store(testCtx(), in)
	is.NoErr(err) // store should not return error

	asserts()

	// update
	err = storage.Store(testCtx(), in)
	is.NoErr(err) // store should not return error

	asserts()
}
