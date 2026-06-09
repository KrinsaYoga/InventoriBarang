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
	cetakData()
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

func hitungTotalStok(n int) int {
	/* Spesifikasi: Menghitung total seluruh stok secara rekursif */
	if n == 0 {
		return 0
	}
	return dataBarang[n-1].Stok + hitungTotalStok(n-1)
}

func cetakData() {
	var i int
	fmt.Println("\n---------------------------------------------------------")
	fmt.Printf("%-5s %-15s %-15s %-8s %-10s\n", "ID", "Nama", "Kategori", "Stok", "Harga")
	fmt.Println("---------------------------------------------------------")
	i = 0
	for i < nBarang {
		fmt.Printf("%-5d %-15s %-15s %-8d %-10d\n", 
			dataBarang[i].ID, dataBarang[i].Nama, dataBarang[i].Kategori, 
			dataBarang[i].Stok, dataBarang[i].Harga)
		i = i + 1
	}
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Total Stok (Rekursif):", hitungTotalStok(nBarang))
}

func tambahBarang() {
	if nBarang < NMAX {
		fmt.Print("ID Barang: ")
		fmt.Scan(&dataBarang[nBarang].ID)
		fmt.Print("Nama Barang: ")
		fmt.Scan(&dataBarang[nBarang].Nama)
		fmt.Print("Kategori: ")
		fmt.Scan(&dataBarang[nBarang].Kategori)
		fmt.Print("Stok: ")
		fmt.Scan(&dataBarang[nBarang].Stok)
		fmt.Print("Harga: ")
		fmt.Scan(&dataBarang[nBarang].Harga)
		nBarang = nBarang + 1
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		fmt.Println("Gudang Penuh!")
	}
}