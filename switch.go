package main

import "fmt"

func main() {
	fmt.Println("Test")

	// Bentuk standar
	day := "Senin"

	switch day {
	case "Sabtu", "Minggu":
	   	fmt.Println("Akhir pekan!")
	case "Senin":
	   	fmt.Println("Awal minggu kerja.")
		fallthrough
	default:
	   	fmt.Println("Hari kerja biasa.")
	}
}
