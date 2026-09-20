package main

import (
	"fmt"
	"learngo/pharmacy"
)

func main() {

	for {

		fmt.Println("\n==============================")
		fmt.Println("   PHARMACY MEDICINE INVENTORY")
		fmt.Println("==============================")
		fmt.Println("1. Add Medicine")
		fmt.Println("2. Find Medicine")
		fmt.Println("3. Dispense Medicine")
		fmt.Println("4. Discard Expired Batch")
		fmt.Println("==============================")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			addMedicine()

		case 2:
			findMedicine()

		case 3:
			dispenseMedicine()

		case 4:
			discardExpiredBatch()


		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func addMedicine() {

	var medicine pharmacy.Medicine

	fmt.Print("Enter Batch Number: ")
	fmt.Scanln(&medicine.BatchNumber)

	fmt.Print("Enter Brand Name: ")
	fmt.Scanln(&medicine.BrandName)

	fmt.Print("Enter Generic Name: ")
	fmt.Scanln(&medicine.GenericName)

	fmt.Print("Enter Stock Units: ")
	fmt.Scanln(&medicine.StockUnits)

	fmt.Print("Enter Expiry Year: ")
	fmt.Scanln(&medicine.ExpiryYear)

	pharmacy.AddMedicine(medicine)

	fmt.Println("Medicine added successfully!")
}

func findMedicine() {

	var batch string

	fmt.Print("Enter Batch Number: ")
	fmt.Scanln(&batch)

	medicine, err := pharmacy.FindMedicine(batch)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nMedicine Found!")
	fmt.Println("Batch Number:", medicine.BatchNumber)
	fmt.Println("Brand Name:", medicine.BrandName)
	fmt.Println("Generic Name:", medicine.GenericName)
	fmt.Println("Stock Units:", medicine.StockUnits)
	fmt.Println("Expiry Year:", medicine.ExpiryYear)
}

func dispenseMedicine() {

	var batch string
	var qty int

	fmt.Print("Enter Batch Number: ")
	fmt.Scanln(&batch)

	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&qty)

	err := pharmacy.DispenseMedicine(batch, qty)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Medicine dispensed successfully!")
}

func discardExpiredBatch() {

	var batch string

	fmt.Print("Enter Batch Number: ")
	fmt.Scanln(&batch)

	err := pharmacy.DiscardExpiredBatch(batch)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Expired batch discarded successfully!")
}