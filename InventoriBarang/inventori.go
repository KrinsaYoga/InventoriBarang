package main

import "fmt"

const NMAX int = 1000

// Tipe Bentukan Struktur
type Barang struct {
	ID       int
	Nama     string
	Kategori string
	Stok     int
	Harga    int
}

type Transaksi struct {
	IDTransaksi int
	IDBarang    int
	Tipe        string
	Jumlah      int
}

// Variabel Global (Hanya array utama)
var dataBarang [NMAX]Barang
var nBarang int = 0

var dataTransaksi [NMAX]Transaksi
var nTransaksi int = 0


func main() {
	isiDataDummy()
	fmt.Println("Data dummy berhasil dimuat")
}

// ==========================================
// SUBPROGRAM DATA DUMMY
// ==========================================
func isiDataDummy() {
	// Spesifikasi: Mengisi data awal ke dalam array untuk mempermudah pengujian 
	if nBarang < NMAX {
		dataBarang[nBarang].ID = 105
		dataBarang[nBarang].Nama = "Beras"
		dataBarang[nBarang].Kategori = "Pokok"
		dataBarang[nBarang].Stok = 50
		dataBarang[nBarang].Harga = 15000
		nBarang = nBarang + 1
	}
	if nBarang < NMAX {
		dataBarang[nBarang].ID = 101
		dataBarang[nBarang].Nama = "Minyak"
		dataBarang[nBarang].Kategori = "Pokok"
		dataBarang[nBarang].Stok = 20
		dataBarang[nBarang].Harga = 18000
		nBarang = nBarang + 1
	}
	if nBarang < NMAX {
		dataBarang[nBarang].ID = 103
		dataBarang[nBarang].Nama = "Gula"
		dataBarang[nBarang].Kategori = "Pokok"
		dataBarang[nBarang].Stok = 35
		dataBarang[nBarang].Harga = 14000
		nBarang = nBarang + 1
	}
	if nBarang < NMAX {
		dataBarang[nBarang].ID = 102
		dataBarang[nBarang].Nama = "Garam"
		dataBarang[nBarang].Kategori = "Bumbu"
		dataBarang[nBarang].Stok = 100
		dataBarang[nBarang].Harga = 2000
		nBarang = nBarang + 1
	}
}