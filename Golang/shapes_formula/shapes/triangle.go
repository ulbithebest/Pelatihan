package shapes

import (
	"fmt" // import package fmt untuk operasi input output
)

// Membuat fungsi untuk menghitung luas segitiga
func LuasSegitiga(){
	// Deklarasi variabel alas dan tinggi
	var alas, tinggi float64

	// Input alas dan tinggi segitiga
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scanln(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)

	// Menghitung luas segitiga
	luas := 0.5 * alas * tinggi

	// Menampilkan hasil luas segitiga
	fmt.Println("Luas segitiga adalah: ", luas)
}