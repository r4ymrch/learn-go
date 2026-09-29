package main

import "fmt"

// func main() {
// 	// map dengan key berupa string dan value berupa int
// 	ages := map[string]int{
//     	"Adnan":  20,
//     	"Naufal": 22,
//     	"Budi":   25,
// 	}
//
// 	// Membuat map kosong
// 	scores := make(map[string]int)
//
// 	// Mengisi data
// 	scores["Matematika"] = 90
// 	scores["Fisika"] = 85
// }

// Operasi dasar Map
// func main() {
//     // 1. Inisialisasi Map
//     student := map[string]string{
//         "npm":    "2026001",
//         "name":   "Adnan",
//         "jurusan": "Informatika",
//     }
//
//     // 2. Menambah & Mengubah Data
//     student["angkatan"] = "2026" // Menambah key baru
//     student["jurusan"] = "Teknik Informatika" // Mengubah value yang sudah ada
//
//     // 3. Membaca Data
//     fmt.Println("Nama Student:", student["name"])
//
//     // 4. Menghapus Data (menggunakan delete)
//     delete(student, "npm") // Menghapus key "npm"
//
//     // 5. Checking Exists (Comma-ok Idiom)
//     // Saat mengakses key, Go mengembalikan 2 nilai: (value, isExist)
//     value, isExist := student["npm"]
//     if isExist {
//         fmt.Println("NPM ditemukan:", value)
//     } else {
//         fmt.Println("NPM tidak ditemukan / sudah dihapus!")
//     }
//
//     // 6. Mengecek Jumlah Pasangan Data
//     fmt.Println("Jumlah data di map:", len(student))
// }

// Loop di Map
func main() {
    prices := map[string]int{
        "Kopi":  15000,
        "Teh":   10000,
        "Roti":  12000,
    }

    // Iterasi key dan value
    for item, price := range prices {
        fmt.Printf("Harga %s adalah Rp %d\n", item, price)
    }

    // Jika hanya butuh key-nya saja
    for item := range prices {
        fmt.Println("Item:", item)
    }
}
