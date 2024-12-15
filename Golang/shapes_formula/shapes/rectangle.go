package shapes

import (
	"fmt" // import package fmt untuk operasi input output
)

// Membuat fungsi untuk menghitung luas persegi panjang
func LuasPersegiPanjang(){
	// Deklarasi variabel panjang dan lebar
	var panjang, lebar float64

	// Input panjang dan lebar persegi panjang
	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scanln(&lebar)

	// Menghitung luas persegi panjang
	luas := panjang * lebar

	// Menampilkan hasil luas persegi panjang
	fmt.Println("Luas persegi panjang adalah: ", luas)
}
