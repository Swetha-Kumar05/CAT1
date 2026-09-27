package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"learngo/pharmacy"
)

func main() {

	// Load medicines from CSV
	err := pharmacy.LoadMedicines()
	if err != nil {
		fmt.Println("Error loading medicines:", err)
	}

	// Goroutine
	var wg sync.WaitGroup
	wg.Add(1)

	go welcomeMessage(&wg)

	wg.Wait()

	// Command-line arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("\nNo command given.")
		fmt.Println("\nAvailable commands:")
		fmt.Println("add      - Add medicine")
		fmt.Println("find     - Find medicine")
		fmt.Println("dispense - Dispense medicine")
		fmt.Println("discard  - Discard expired batch")
		fmt.Println("list     - Display all medicines")
		return
	}

	command := args[0]

	switch command {

	case "add":

		if len(args) != 6 {
			fmt.Println("Usage:")
			fmt.Println("go run . add <batch> <brand> <generic> <stock> <expiry>")
			return
		}

		stock, err1 := strconv.Atoi(args[4])
		expiry, err2 := strconv.Atoi(args[5])

		if err1 != nil || err2 != nil {
			fmt.Println("Stock and expiry must be numbers.")
			return
		}

		medicine := pharmacy.Medicine{
			BatchNumber: args[1],
			BrandName:   args[2],
			GenericName: args[3],
			StockUnits:  stock,
			ExpiryYear:  expiry,
		}

		err := pharmacy.AddMedicine(medicine)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Medicine added successfully!")

	case "find":

		if len(args) != 2 {
			fmt.Println("Usage:")
			fmt.Println("go run . find <batch>")
			return
		}

		medicine, err := pharmacy.FindMedicine(args[1])

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("\nMedicine Found!")
		fmt.Println("-------------------------")
		fmt.Println("Batch Number:", medicine.BatchNumber)
		fmt.Println("Brand Name:", medicine.BrandName)
		fmt.Println("Generic Name:", medicine.GenericName)
		fmt.Println("Stock Units:", medicine.StockUnits)
		fmt.Println("Expiry Year:", medicine.ExpiryYear)

	case "dispense":

		if len(args) != 3 {
			fmt.Println("Usage:")
			fmt.Println("go run . dispense <batch> <quantity>")
			return
		}

		qty, err := strconv.Atoi(args[2])

		if err != nil {
			fmt.Println("Quantity must be a number.")
			return
		}

		err = pharmacy.DispenseMedicine(args[1], qty)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Medicine dispensed successfully!")

	case "discard":

		if len(args) != 2 {
			fmt.Println("Usage:")
			fmt.Println("go run . discard <batch>")
			return
		}

		err := pharmacy.DiscardExpiredBatch(args[1])

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Expired batch discarded successfully!")

	case "list":

		pharmacy.DisplayMedicines()

	default:

		fmt.Println("Unknown command:", command)
		fmt.Println("\nAvailable commands:")
		fmt.Println("add")
		fmt.Println("find")
		fmt.Println("dispense")
		fmt.Println("discard")
		fmt.Println("list")
	}
}

func welcomeMessage(wg *sync.WaitGroup) {

	defer wg.Done()

	time.Sleep(1 * time.Second)

	fmt.Println("Welcome to Pharmacy Medicine Inventory System")
}
