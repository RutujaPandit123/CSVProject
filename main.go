package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
)

func main() {

	currentData := [][]string{
		{"Mattress ID", "Date Rotated", "Room No", "Mattress Status", "Employee Name", "Employee ID"},
		{"12345", "2024-03-01", "101", "flip", "mohit sharma", "001"},
		{"54321", "2024-03-02", "102", "rotate", "rahul sharma", "002"},
		{"67890", "2024-03-03", "103", "rotate", "poonam dixit", "003"},
	}

	projectedData := [][]string{
		{"Mattress ID", "Date Rotated", "Room No", "Mattress Status", "Employee Name", "Employee ID"},
		{"13579", "2024-03-04", "104", "flip", "moni pandit", "004"},
		{"24680", "2024-03-05", "105", "rotate", "sanika raut", "005"},
		{"98765", "2024-03-06", "106", "flip", "swapnali biradar", "006"},
	}

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"Mattress ID", "Date Rotated", "Room No", "Mattress Status", "Employee Name", "Employee ID"})
	if err != nil {
		log.Fatal("Error writing headers for current data:", err)
	}

	err = writer.WriteAll(currentData)
	if err != nil {
		log.Fatal("Error writing current data to CSV:", err)
	}

	err = writer.Write([]string{"Mattress ID", "Date Rotated", "Room No", "Mattress Status", "Employee Name", "Employee ID"})
	if err != nil {
		log.Fatal("Error writing headers for projected data:", err)
	}

	err = writer.WriteAll(projectedData)
	if err != nil {
		log.Fatal("Error writing projected data to CSV:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "data.csv")
	})

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
