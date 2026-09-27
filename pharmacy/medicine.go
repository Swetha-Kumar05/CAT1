package pharmacy

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Medicine struct {
	BatchNumber string
	BrandName   string
	GenericName string
	StockUnits  int
	ExpiryYear  int
}

var medicines []Medicine

const fileName = "medicines.csv"

func AddMedicine(m Medicine) error {

	currentYear := 2026

	if m.BatchNumber == "" {
		return errors.New("batch number cannot be empty")
	}

	if m.ExpiryYear < currentYear {
		return errors.New("expiry year is invalid")
	}

	if m.StockUnits < 0 {
		return errors.New("stock units cannot be negative")
	}

	// Check duplicate batch number
	for _, medicine := range medicines {

		if medicine.BatchNumber == m.BatchNumber {
			return errors.New("batch number already exists")
		}
	}

	medicines = append(medicines, m)

	return SaveMedicines()
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

	// Reduce stock
	medicine.StockUnits -= qty

	// Save updated data
	return SaveMedicines()
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

			// Remove medicine from slice
			medicines = append(
				medicines[:i],
				medicines[i+1:]...,
			)

			// Save updated data
			return SaveMedicines()
		}
	}

	return errors.New("medicine batch not found")
}

func SaveMedicines() error {

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	// CSV header
	header := []string{
		"BatchNumber",
		"BrandName",
		"GenericName",
		"StockUnits",
		"ExpiryYear",
	}

	err = writer.Write(header)

	if err != nil {
		return err
	}

	// Write medicine records
	for _, medicine := range medicines {

		record := []string{
			medicine.BatchNumber,
			medicine.BrandName,
			medicine.GenericName,
			strconv.Itoa(medicine.StockUnits),
			strconv.Itoa(medicine.ExpiryYear),
		}

		err = writer.Write(record)

		if err != nil {
			return err
		}
	}

	// Make sure all data is written
	writer.Flush()

	return writer.Error()
}

func LoadMedicines() error {

	file, err := os.Open(fileName)

	if err != nil {

		if os.IsNotExist(err) {
			return SaveMedicines()
		}

		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	// Read CSV header
	_, err = reader.Read()

	if err != nil {
		return err
	}

	medicines = nil

	for {

		record, err := reader.Read()

		if err != nil {

			if err.Error() == "EOF" {
				break
			}

			return err
		}

		if len(record) != 5 {
			continue
		}

		stock, err1 := strconv.Atoi(record[3])

		if err1 != nil {
			continue
		}

		expiry, err2 := strconv.Atoi(record[4])

		if err2 != nil {
			continue
		}

		medicine := Medicine{
			BatchNumber: record[0],
			BrandName:   record[1],
			GenericName: record[2],
			StockUnits:  stock,
			ExpiryYear:  expiry,
		}

		medicines = append(medicines, medicine)
	}

	return nil
}

func DisplayMedicines() {

	if len(medicines) == 0 {

		fmt.Println("No medicines available.")

		return
	}

	for _, medicine := range medicines {

		fmt.Println("------------------------------")
		fmt.Println("Batch Number:", medicine.BatchNumber)
		fmt.Println("Brand Name:", medicine.BrandName)
		fmt.Println("Generic Name:", medicine.GenericName)
		fmt.Println("Stock Units:", medicine.StockUnits)
		fmt.Println("Expiry Year:", medicine.ExpiryYear)
	}
}
