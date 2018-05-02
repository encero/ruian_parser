package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/encero/ruian_parser/ent"
	"github.com/encero/ruian_parser/ent/city"
)

type Storage struct {
	handlers []StorageHandler
	entc     *ent.Client
}

func NewStorage(entc *ent.Client) *Storage {
	return &Storage{
		entc: entc,
	}
}

type StorageHandler func(ctx context.Context, entc *ent.Client, item interface{}) (bool, error)

func (s *Storage) AddHandler(handler StorageHandler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Storage) Store(ctx context.Context, item interface{}) error {
	for _, handler := range s.handlers {
		used, err := handler(ctx, s.entc, item)
		if used {
			return err
		}
	}

	return fmt.Errorf("item not supported")
}

func StoreCity(ctx context.Context, entc *ent.Client, item interface{}) (bool, error) {
	c, ok := item.(City)
	if !ok {
		return false, nil
	}

	update, err := MapCity(entc, c)
	if err != nil {
		return true, fmt.Errorf("city mapping err: %v", err)
	}

	err = update.OnConflictColumns(city.FieldID).UpdateNewValues().Exec(ctx)
	if err != nil {
		return true, fmt.Errorf("city save err: %v city: %+v", err, c)
	}

	return true, nil
}

func StoreStreet(ctx context.Context, entc *ent.Client, item interface{}) (bool, error) {
	s, ok := item.(Street)
	if !ok {
		return false, nil
	}

    code, err := strconv.Atoi(s.Code)
    if err != nil {
        return true, fmt.Errorf("street code is not int")
    }


    street, err := GetOrCreateStreet(ctx, entc, int32(code))

    if err != nil {
        return true, fmt.Errorf("get or create street:%w", err)
    }

	update, err := UpdateStreet(street, s)
	if err != nil {
		return true, fmt.Errorf("street mapping err: %v", err)
	}

    _, err = update.Save(ctx)
    if err != nil {
        return true, fmt.Errorf("street save: %w", err)
    }

	return true, nil
}

func StoreAddressPlace(ctx context.Context, entc *ent.Client, item interface{}) (bool, error) {
	aps, ok := item.(AddressPlace)
	if !ok {
		return false, nil
	}

    code, err := strconv.Atoi(aps.Code)
    if err != nil {
        return true, fmt.Errorf("adress place code is not int")
    }

    // todo number parsing, move to elsewhere
    ap, err := GetOrCreateAddressPlace(ctx, entc, int32(code))
    if err != nil {
        return true, fmt.Errorf("adress place not found or cant be created: %w", err)
    }

    apu, err := UpdateAddressPlace(ap, aps)
    if err != nil {
        return true, fmt.Errorf("address place update: %w", err)
    }

    _, err = apu.Save(ctx)
    if err != nil {
        return true, fmt.Errorf("adress place save: %w", err)
    }

	return true, nil
}
