package main

import (
    // "errors"
    "fmt"
)

// Fungsi tanpa parameter
// func greet() {
//     fmt.Println("Selamat datang di pemrogramman Go!")
// }
//
// func main() {
//     greet() // Memanggil fungsi
// }

// Fungsi dengan parameter
// Tipe data `a` dan `b` sama-sama `int`
// func add(a, b int) int {
//     return a + b
// }
//
// func main() {
//     result := add(10, 20)
//     fmt.Println("Hasil penjumlahan:", result) // 30
// }

// Multiple retrun value
// Mengembalikan dua nilai: float64 dan error
// func divide(a, b float64) (float64, error) {
//     if b == 0 {
//         return 0, errors.New("tidak bisa membagi dengan angka nol")
//     }
//     return a / b, nil
// }
//
// func main() {
//     res, err := divide(10, 2)
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Hasil pembagian:", res) // 5
//     }
//
//     // Mengabaikan salah satu return value dengan blank identifier (_)
//     resOnly, _ := divide(20, 4)
//     fmt.Println("Hasil saja:", resOnly)
// }

// Fungsi parameter tak hingga
// func sumAll(numbers ...int) int {
//     total := 0
//     for _, num := range numbers {
//         total += num
//     }
//     return total
// }
//
// func main() {
//     fmt.Println(sumAll(1, 2, 3))       // 6
//     fmt.Println(sumAll(10, 20, 30, 40)) // 100
//
//     // BISA JUGA MEMPASING SLICE LANGSUNG DENGAN `...`
//     nums := []int{5, 10, 15}
//     fmt.Println(sumAll(nums...))       // 30
// }

// Anonymous func
func main() {
    // 1. Anonymous Function disimpan dalam variabel
    multiply := func(a, b int) int {
        return a * b
    }
    fmt.Println("Hasil perkalian:", multiply(4, 5))

    // 2. Closure mengakses variabel luar
    counter := 0
    increment := func() int {
        counter++ // Memanipulasi variabel di luar scope fungsi
        return counter
    }

    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
}
