package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Struktur Data Barang
type Barang struct {
	Kode       string
	Nama       string
	Harga      int
	JumlahStok int
}

// Variabel global untuk menyimpan daftar barang
var daftarBarang [100]Barang
var jumlahBarang int

// Fungsi untuk menampilkan menu utama
func tampilkanMenu() {
	fmt.Println("APLIKASI INVENTORY BARANG:")
	fmt.Println("1. Daftar Data Barang")
	fmt.Println("2. Input barang")
	fmt.Println("3. Hapus barang")
	fmt.Println("4. Edit barang")
	fmt.Println("5. Cari barang")
	fmt.Println("6. Exit")
	fmt.Print("Pilih menu: ")
}

// Fungsi untuk menambah barang
func inputBarang() {
	if jumlahBarang < len(daftarBarang) {
		var barang Barang
		fmt.Print("Masukkan kode barang: ")
		fmt.Scan(&barang.Kode)
		fmt.Print("Masukkan nama barang: ")
		fmt.Scan(&barang.Nama)
		fmt.Print("Masukkan jumlah stok: ")
		fmt.Scan(&barang.JumlahStok)
		fmt.Print("Masukkan harga: ")
		fmt.Scan(&barang.Harga)
		daftarBarang[jumlahBarang] = barang
		jumlahBarang++
		fmt.Println("Barang berhasil ditambahkan!")
	} else {
		fmt.Println("Data barang penuh!")
	}
	kembaliKeMenu()
}

// Fungsi untuk menghapus barang
func hapusBarang() {
	var kode string
	fmt.Print("Masukkan kode barang yang ingin dihapus: ")
	fmt.Scan(&kode)

	index := cariBarangByKode(kode)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
	} else {
		// Menggeser elemen setelah index yang ditemukan ke depan
		for i := index; i < jumlahBarang-1; i++ {
			daftarBarang[i] = daftarBarang[i+1]
		}
		jumlahBarang--
		fmt.Println("Barang berhasil dihapus.")
	}
	kembaliKeMenu()
}

// Fungsi untuk membaca input string
// func bacaInputString(prompt string) string {
// 	fmt.Print(prompt)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	return scanner.Text()
// }

func bacaInputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n') // '\n' adalah delimiter untuk berhenti membaca
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		return ""
	}

	// Menghapus karakter newline dari input, jika ada
	return strings.TrimSpace(input)
}

// Fungsi untuk mengedit barang
func editBarang() {
	var stokBaru, hargaBaru int
	var kode, namaBaru string
	fmt.Print("Masukkan kode barang yang ingin diubah: ")
	fmt.Scan(&kode)

	index := cariBarangByKode(kode)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
	} else {
		// namaBaru := bacaInputString("Masukkan nama baru barang: ")
		fmt.Print("Masukkan Nama baru: ")
		fmt.Scan(&namaBaru)

		fmt.Print("Masukkan jumlah stok baru: ")
		fmt.Scan(&stokBaru)

		// Menangani input harga baru
		fmt.Print("Masukkan harga baru: ")
		fmt.Scan(&hargaBaru)

		daftarBarang[index].Nama = namaBaru
		daftarBarang[index].JumlahStok = stokBaru
		daftarBarang[index].Harga = hargaBaru

		fmt.Println("Barang berhasil diubah.")
	}

	kembaliKeMenu()
}

// Fungsi untuk mencari barang
func cariBarang() {
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Println("3. Cari barang termahal")
	fmt.Println("4. Cari barang termurah")
	fmt.Println("5. Cari barang berdasarkan nama")
	fmt.Println("6. Kembali ke menu utama")
	fmt.Print("Pilih metode pencarian: ")
	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var kode string
		fmt.Print("Masukkan kode barang yang dicari: ")
		fmt.Scan(&kode)
		index := cariBarangByKode(kode)
		if index == -1 {
			fmt.Println("Barang tidak ditemukan.")
		} else {
			fmt.Printf("Barang ditemukan: Kode: %s, Nama: %s, Stok: %d, Harga: %d\n",
				daftarBarang[index].Kode, daftarBarang[index].Nama, daftarBarang[index].JumlahStok, daftarBarang[index].Harga)
		}
	case 2:
		var kode string
		fmt.Print("Masukkan kode barang yang dicari: ")
		fmt.Scan(&kode)
		index := binarySearch(kode)
		if index == -1 {
			fmt.Println("Barang tidak ditemukan.")
		} else {
			fmt.Printf("Barang ditemukan: Kode: %s, Nama: %s, Stok: %d, Harga: %d\n",
				daftarBarang[index].Kode, daftarBarang[index].Nama, daftarBarang[index].JumlahStok, daftarBarang[index].Harga)
		}
	case 3:
		cariBarangTermahal()
	case 4:
		cariBarangTermurah()
	case 5:
		var nama string
		fmt.Print("Masukkan nama barang yang dicari: ")
		fmt.Scan(&nama)
		cariBarangByNama(nama)
	case 6:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
		cariBarang()
	}
	kembaliKeMenu()
}

// Fungsi untuk mencari barang yang paling mahal
func cariBarangTermahal() {
	if jumlahBarang == 0 {
		fmt.Println("Daftar barang kosong.")
		return
	}
	maxHarga := math.MinInt
	index := -1
	for i := 0; i < jumlahBarang; i++ {
		if daftarBarang[i].Harga > maxHarga {
			maxHarga = daftarBarang[i].Harga
			index = i
		}
	}
	fmt.Printf("Barang termahal: Kode: %s, Nama: %s, Stok: %d, Harga: %d\n",
		daftarBarang[index].Kode, daftarBarang[index].Nama, daftarBarang[index].JumlahStok, daftarBarang[index].Harga)
}

// Fungsi untuk mencari barang yang paling murah
func cariBarangTermurah() {
	if jumlahBarang == 0 {
		fmt.Println("Daftar barang kosong.")
		return
	}
	minHarga := math.MaxInt
	index := -1
	for i := 0; i < jumlahBarang; i++ {
		if daftarBarang[i].Harga < minHarga {
			minHarga = daftarBarang[i].Harga
			index = i
		}
	}
	fmt.Printf("Barang termurah: Kode: %s, Nama: %s, Stok: %d, Harga: %d\n",
		daftarBarang[index].Kode, daftarBarang[index].Nama, daftarBarang[index].JumlahStok, daftarBarang[index].Harga)
}

// Fungsi untuk mencari barang berdasarkan nama
func cariBarangByNama(nama string) {
	fmt.Println("Hasil pencarian barang berdasarkan nama:")
	found := false
	for i := 0; i < jumlahBarang; i++ {
		if contains(daftarBarang[i].Nama, nama) {
			fmt.Printf("Kode: %s, Nama: %s, Stok: %d, Harga: %d\n",
				daftarBarang[i].Kode, daftarBarang[i].Nama, daftarBarang[i].JumlahStok, daftarBarang[i].Harga)
			found = true
		}
	}
	if !found {
		fmt.Println("Barang tidak ditemukan.")
	}

	fmt.Println("1. Cari lagi")
	fmt.Println("2. Kembali ke menu utama")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		namaBaru := bacaInputString("Masukkan nama barang yang dicari: ")
		cariBarangByNama(namaBaru)
	case 2:
		return
	default:
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama.")
	}
	kembaliKeMenu()
}

// Fungsi untuk memeriksa apakah string `sub` terdapat dalam `str`
func contains(str, sub string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(sub))
}

// Fungsi untuk mencari barang berdasarkan kode menggunakan Sequential Search
func cariBarangByKode(kode string) int {
	for i := 0; i < jumlahBarang; i++ {
		if daftarBarang[i].Kode == kode {
			return i
		}
	}
	return -1
}

// Fungsi untuk mencari barang menggunakan Binary Search
func binarySearch(kode string) int {
	// Pastikan data sudah terurut sebelum melakukan Binary Search
	urutkanAscending()
	left, right := 0, jumlahBarang-1

	for left <= right {
		mid := (left + right) / 2
		if daftarBarang[mid].Kode == kode {
			return mid
		} else if daftarBarang[mid].Kode < kode {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Fungsi untuk menampilkan daftar barang
func tampilkanDaftarBarang() {
	fmt.Println("Daftar Barang:")
	fmt.Println("No.   Kode       Nama                 Stok       Harga")
	for i := 0; i < jumlahBarang; i++ {
		fmt.Printf("%-5d %-10s %-20s %-10d %d\n",
			i+1, daftarBarang[i].Kode, daftarBarang[i].Nama, daftarBarang[i].JumlahStok, daftarBarang[i].Harga)
	}
	pilihSorting()
}

// Fungsi untuk memilih metode sorting
func pilihSorting() {
	fmt.Println("1. Urutkan secara ascending")
	fmt.Println("2. Urutkan secara descending")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih metode sorting: ")
	var pilihan int
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		urutkanAscending()
	case 2:
		urutkanDescending()
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
		pilihSorting()
	}
}

// Fungsi untuk mengurutkan barang secara ascending
func urutkanAscending() {
	for i := 0; i < jumlahBarang-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahBarang; j++ {
			if daftarBarang[j].Harga < daftarBarang[minIdx].Harga { // Bandingkan berdasarkan harga
				minIdx = j
			}
		}
		if minIdx != i {
			daftarBarang[i], daftarBarang[minIdx] = daftarBarang[minIdx], daftarBarang[i]
		}
	}
	fmt.Println("Data barang telah diurutkan secara ascending berdasarkan harga.")
	tampilkanDaftarBarang()
}

// Fungsi untuk mengurutkan barang secara descending
func urutkanDescending() {
	for i := 0; i < jumlahBarang-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahBarang; j++ {
			if daftarBarang[j].Harga > daftarBarang[maxIdx].Harga { // Bandingkan berdasarkan harga
				maxIdx = j
			}
		}
		if maxIdx != i {
			daftarBarang[i], daftarBarang[maxIdx] = daftarBarang[maxIdx], daftarBarang[i]
		}
	}
	fmt.Println("Data barang telah diurutkan secara descending berdasarkan harga.")
	tampilkanDaftarBarang()
}

// Fungsi untuk kembali ke menu utama
func kembaliKeMenu() {
	fmt.Println("1. Kembali ke menu utama")
	fmt.Print("Pilih: ")
	var pilihan int
	for {
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			fmt.Print("Pilih: ")
		}
	}
}

func main() {
	daftarBarang = [100]Barang{
		{"B01", "Citato", 10000, 100},
		{"B02", "Nabati", 12000, 80},
		{"B03", "Tenggo", 15000, 60},
		{"B04", "Chitato", 11000, 90},
		{"B05", "Pringles", 20000, 50},
		{"B06", "Oops!", 13000, 70},
		{"B07", "Kaleng", 14000, 65},
		{"B08", "Kacang Garuda", 18000, 55},
		{"B09", "Richeese", 87000, 75},
		{"B10", "Lays", 22000, 40},
		{"B11", "Gery", 5000, 85},
		{"B12", "Cheetos", 20000, 45},
		{"B13", "Stick Snack", 14000, 70},
		{"B14", "Kue Cubir", 12000, 95},
		{"B15", "Kentang Goreng", 17000, 60},
		{"B16", "Pop Mie Snack", 13000, 85},
		{"B17", "Supermi", 11000, 80},
		{"B18", "Gula Melaka", 50000, 50},
		{"B19", "Barbeque Snack", 14000, 70},
		{"B20", "Snack Coklat", 19000, 40},
	}
	jumlahBarang = 20

	for {
		tampilkanMenu()
		var pilihan int
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tampilkanDaftarBarang()
		case 2:
			inputBarang()
		case 3:
			hapusBarang()
		case 4:
			editBarang()
		case 5:
			cariBarang()
		case 6:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
