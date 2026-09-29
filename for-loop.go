package main

import "fmt"

func main() {
	// Gaya C/Java
    // for i := 1; i <= 5; i++ {
    //     fmt.Println("Perulangan ke-", i)
    // }

    // For sebagai while
	//    n := 1
	//
	// for n <= 32 {
	//    	fmt.Println("Nilai n:", n)
	//    	n *= 2 // Pengali 2 di setiap iterasi
	// }

	// Loop tanpa kondisi
	counter := 1

	for {
    	fmt.Println("Iterasi ke-", counter)
    	if counter == 3 {
        	fmt.Println("Stop perulangan!")
        	break // Keluar dari loop
    	}
    	counter++
	}
}
