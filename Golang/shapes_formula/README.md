# PENGENALAN GOLANG

## **Instalasi Golang**
Ikuti langkah-langkah berikut untuk menginstal Golang pada komputer Anda:

1. **Unduh Golang**
   - Kunjungi situs resmi Golang di [https://golang.org/dl/](https://golang.org/dl/).
   - Pilih versi yang sesuai dengan sistem operasi Anda (Windows, macOS, atau Linux).

2. **Instal Golang**
   - Jalankan file instalasi yang telah diunduh.
   - Ikuti langkah-langkah instalasi hingga selesai.

3. **Verifikasi Instalasi**
   - Setelah proses instalasi selesai, buka terminal atau command prompt.
   - Jalankan perintah berikut untuk memastikan Golang telah terinstal dengan benar:
     ```bash
     go version
     ```
   - Jika berhasil, akan muncul informasi versi Golang yang terinstal, misalnya:
     ```
     go version go1.21.1 windows/amd64
     ```

---

## **Quickstart: Menulis Program Golang di VS Code**
Untuk memulai membuat program menggunakan Golang dalam Visual Studio Code (VS Code), ikuti langkah-langkah berikut:

### **Persiapan VS Code**
1. **Install VS Code**
   - Jika Anda belum memiliki VS Code, unduh dan instal dari situs resmi: [https://code.visualstudio.com/](https://code.visualstudio.com/).

2. **Pasang Ekstensi Golang**
   - Jalankan VS Code.
   - Buka **Extension Manager** dengan menekan `Ctrl + Shift + X`.
   - Ketik "go" pada kotak pencarian, lalu tekan `Enter`.
   - Temukan ekstensi **"Go" by the Go Team at Google** dan klik tombol **Install**.

3. **Instalasi Tools Pendukung**
   - Setelah ekstensi terinstal, buka **Command Palette** dengan menekan `Ctrl + Shift + P`.
   - Cari dan pilih perintah **"Go: Install/Update Tools"**.
   - Pilih semua tools yang tersedia, lalu klik **OK**.
   - Tunggu hingga proses instalasi tools selesai.

### **Membuat Program Pertama dengan Golang**
1. **Persiapan Folder Proyek**
   - Buat sebuah folder untuk proyek Anda, misalnya: `hello`.
   - Buka folder tersebut di VS Code.

2. **Inisialisasi Modul Golang**
   - Buka terminal di VS Code (tekan `Ctrl +  ` atau buka menu **View > Terminal**).
   - Jalankan perintah berikut untuk menginisialisasi modul Golang:
     ```bash
     go mod init hello
     ```
   - Setelah berhasil, akan muncul file `go.mod` di folder proyek Anda. File ini digunakan untuk mengelola dependensi proyek.

3. **Buat File Program**
   - Buat file baru bernama `main.go` di dalam folder proyek.
   - Tulis kode berikut di dalam file `main.go`:
     ```go
     package main

     import (
         "fmt"
     )

     func main() {
         fmt.Println("Hello, World!")
     }
     ```
   
   #### Penjelasan Sintaks:
   - `package main`: Mendeklarasikan bahwa file ini merupakan bagian dari package utama yang menjadi titik awal eksekusi program.
   - `import "fmt"`: Mengimpor package `fmt` yang digunakan untuk mencetak teks ke terminal.
   - `func main()`: Fungsi `main` adalah fungsi utama yang dieksekusi saat program dijalankan.
   - `fmt.Println("Hello, World!")`: Fungsi `Println` dari package `fmt` mencetak teks "Hello, World!" ke terminal dan menambahkan baris baru di akhir.

4. **Menjalankan Program**
   - Untuk menjalankan program, buka terminal di folder proyek Anda.
   - Jalankan perintah berikut:
     ```bash
     go run main.go
     ```
   - Jika berhasil, output berikut akan muncul di terminal:
     ```
     Hello, World!
     ```

---

## **Tips dan Catatan**
- **File Utama**: Program Golang selalu dimulai dari fungsi `main` dalam package `main`.
- **Komunitas dan Dokumentasi**:
  - Dokumentasi resmi Golang: [https://go.dev/doc/](https://go.dev/doc/)
  - Golang Playground (untuk mencoba kode secara online): [https://play.golang.org/](https://play.golang.org/)
- **Pengaturan PATH**: Jika Anda mengalami masalah dengan perintah `go`, pastikan direktori Golang telah ditambahkan ke variabel `PATH` pada sistem Anda.
---

## **Penjelasan Singkat**
- **fmt**: `fmt` adalah salah satu package yang sangat sering digunakan di Golang. Package ini digunakan untuk mencetak teks atau data di layar. Misalnya, saat kamu ingin menampilkan pesan "Hello, World!" di layar, kamu akan menggunakan `fmt.Println()`.
- **Package**: Package adalah kumpulan kode yang punya tujuan atau fungsi tertentu. Bayangkan jika kamu menulis banyak kode, kamu bisa mengelompokkan kode yang punya fungsi sama dalam satu package. Misalnya, package `fmt` berisi kode untuk mencetak tulisan di layar.
- **Dependensi**: Dependensi adalah "pustaka" atau "perpustakaan" tambahan yang digunakan oleh proyek kamu. Misalnya, jika kamu butuh fitur untuk mencetak teks di layar, kamu akan menggunakan pustaka `fmt`. Dependensi ini membantu kamu untuk tidak menulis ulang kode yang sudah ada, cukup gunakan pustaka yang sudah ada.
- **Modul**: Modul itu seperti "kotak" yang berisi semua kode yang diperlukan untuk menjalankan proyek di Golang. Kotak ini juga mencatat paket apa saja yang dibutuhkan proyek kamu. Jadi, modul membantu mengatur kode dan dependensi (pustaka lain) yang digunakan agar semuanya tetap terorganisir.
  - `go.mod` adalah file yang mencatat informasi tentang modul, seperti nama modul dan versi dependensi yang digunakan. Ini juga memungkinkan Go untuk mengelola dependensi proyek secara otomatis dan konsisten, baik untuk pengembangan lokal maupun produksi.
- **Ekstensi**: Ekstensi adalah alat tambahan yang bisa dipasang di program seperti Visual Studio Code (VS Code). Ekstensi membantu kamu untuk menulis kode dengan lebih mudah, misalnya memberikan saran otomatis saat kamu mengetik, atau membantu memformat kode supaya rapi.

Selamat belajar dan eksplorasi dengan Golang!

