package main

import "fmt"

func main() {
	fruits := []string{"Apel", "Pisang", "Jeruk"}

	for index, fruit := range fruits {
    	fmt.Printf("Indeks %d: %s\n", index, fruit)
	}
}
