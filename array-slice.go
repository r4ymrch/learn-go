package main

import "fmt"

func main() {
	// Bentuk dasar array
    // // 1. Deklarasi array berukuran 3
    // var numbers [3]int
    // numbers[0] = 10
    // numbers[1] = 20
    // numbers[2] = 30
    //
    // // 2. Deklarasi + Inisialisasi langsung
    // cities := [3]string{"Jakarta", "Bandung", "Surabaya"}
    //
    // // 3. Menggunakan [...] agar ukuran dihitung otomatis berdasarkan jumlah elemen
    // colors := [...]string{"Merah", "Hijau", "Biru", "Kuning"}
    //
    // fmt.Println("Numbers:", numbers)
    // fmt.Println("Cities:", cities)
    // fmt.Println("Jumlah elemen colors:", len(colors)) // 4

    // Slice
	// 1. Deklarasi Slice (tanpa angka di dalam [])
    // var fruits = []string{"Apel", "Pisang", "Jeruk"}
    //
    // // 2. Menambah elemen menggunakan append()
    // fruits = append(fruits, "Mangga", "Anggur")
    //
    // fmt.Println("Slice Fruits:", fruits)
    // fmt.Println("Panjang (len):", len(fruits)) // 5
    // fmt.Println("Kapasitas (cap):", cap(fruits))
    //
    // // 3. Membuat Slice dengan fungsi make()
    // // make([]TipeData, length, capacity)
    // scores := make([]int, 3, 5) // len = 3, cap = 5
    // scores[0] = 80
    // scores[1] = 90
    // scores[2] = 100
    //
    // fmt.Println("Scores:", scores)

    // Slicing
    // Array asal
    numbers := [5]int{10, 20, 30, 40, 50}

    // Membuat Slice dari index 1 hingga sebelum index 4
    subSlice := numbers[1:4] // Mengambil index 1, 2, 3

    fmt.Println("Sub Slice:", subSlice) // [20, 30, 40]

    // Karena Slice adalah referensi, mengubah isi Slice akan mengubah Array aslinya
    subSlice[0] = 99
    fmt.Println("Array Asal setelah diubah:", numbers) // [10, 99, 30, 40, 50]
}
