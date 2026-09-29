package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n=====[ Form Isi Biodata Mahasiswa]=====\n\n")

	fmt.Printf("Masukan Nama: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Masukan NIM: ")
	nim, _ := reader.ReadString('\n')
	nim = strings.TrimSpace(nim)

	fmt.Printf("Masukan Fakultas: ")
	fakultas, _ := reader.ReadString('\n')
	fakultas = strings.TrimSpace(fakultas)

	fmt.Printf("Masukan Prodi: ")
	prodi, _ := reader.ReadString('\n')
	prodi = strings.TrimSpace(prodi)

	ClearScreen()

	fmt.Printf("=====[ Biodata Mahasiswa ]=====\n\nNama: %s\nNIM: %s\nFakultas: %s\nProdi: %s\n\n", name, nim, fakultas, prodi)
}
