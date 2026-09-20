package pharmacy

import (
	"errors"
)

type Medicine struct {
	BatchNumber string
	BrandName   string
	GenericName string
	StockUnits  int
	ExpiryYear  int
}

var medicines []Medicine

func AddMedicine(m Medicine) error {

	currentYear := 2026

	if m.ExpiryYear < currentYear {
		return errors.New("expiry year is invalid")
	}

	if m.StockUnits < 0 {
		return errors.New("stock units cannot be negative")
	}

	medicines = append(medicines, m)

	return nil
}

func FindMedicine(batch string) (*Medicine, error) {

	for i := range medicines {

		if medicines[i].BatchNumber == batch {
			return &medicines[i], nil
		}
	}

	return nil, errors.New("medicine batch not found")
}

func DispenseMedicine(batch string, qty int) error {

	if qty <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	medicine, err := FindMedicine(batch)

	if err != nil {
		return err
	}

	if medicine.StockUnits < qty {
		return errors.New("insufficient stock")
	}

	medicine.StockUnits = medicine.StockUnits - qty

	return nil
}

func DiscardExpiredBatch(batch string) error {

	medicine, err := FindMedicine(batch)

	if err != nil {
		return err
	}

	currentYear := 2026

	if medicine.ExpiryYear >= currentYear {
		return errors.New("medicine batch has not expired")
	}

	for i := range medicines {

		if medicines[i].BatchNumber == batch {

			medicines = append(
				medicines[:i],
				medicines[i+1:]...,
			)

			return nil
		}
	}

	return errors.New("medicine batch not found")
}