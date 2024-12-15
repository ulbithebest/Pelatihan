package main

import (
	"fmt"
	"shapes_formula/shapes" // Import package shapes
)

func main() {
	for {
		
		// Tampilkan menu pilihan
		fmt.Println("\n=== Pilih Operasi ===")
		fmt.Println("1. Hitung Luas Lingkaran")
		fmt.Println("2. Hitung Luas Persegi Panjang")
		fmt.Println("3. Hitung Luas Persegi")
		fmt.Println("4. Hitung Luas Segitiga")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan pilihan (1-5): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		// Pilih operasi berdasarkan input
		switch pilihan {
		case 1:
			shapes.LuasLingkaran()
		case 2:
			shapes.LuasPersegiPanjang()
		case 3:
			shapes.LuasPersegi()
		case 4:
			shapes.LuasSegitiga()
		case 5:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
