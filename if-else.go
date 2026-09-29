package main

import "fmt"

func main() {
    score := 85

    // Basic if else
    if score >= 90 {
        fmt.Println("Nilai: A")
    } else if score >= 75 {
        fmt.Println("Nilai: B")
    } else {
        fmt.Println("Nilai: C")
    }

    // Variabel `length` diinisialisasi sebelum kondisi `;`
	if length := len("Go Programming"); length > 10 {
    	fmt.Println("Teks terlalu panjang:", length)
	} else {
    	fmt.Println("Teks aman:", length)
	}

	// fmt.Println(length) // ERROR: length tidak dikenal di luar blok if
}
