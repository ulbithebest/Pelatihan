package shapes

import (
	"fmt" // import package fmt untuk operasi input output
)

// Membuat fungsi untuk menghitung luas persegi
func LuasPersegi(){
	// Deklarasi variabel sisi
	var sisi float64

	// Input sisi persegi
	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)

	// Menghitung luas persegi
	luas := sisi * sisi

	// Menampilkan hasil luas persegi
	fmt.Println("Luas persegi adalah: ", luas)
}