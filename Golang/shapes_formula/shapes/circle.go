package shapes // deklarasi nama package

import (
	"math" // import package math untuk operasi matematika
	"fmt"  // import package fmt untuk operasi input output
)

// Membuat fungsi untuk menghitung luas lingkaran
func LuasLingkaran(){
	// Deklarasi variabel jari-jari
	var jariJari float64

	// Input jari-jari lingkaran
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	// Menghitung luas lingkaran
	luas := math.Pi * jariJari * jariJari

	// Menampilkan hasil luas lingkaran
	fmt.Println("Luas lingkaran adalah: ", luas)
}

