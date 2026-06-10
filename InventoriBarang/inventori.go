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
	var pilihan, urutPilihan int
	var lanjut bool

	isiDataDummy()
	lanjut = true

	for lanjut {
		fmt.Println("\n=== APLIKASI INVENTORI BARANG ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Tampilkan Barang Berdasarkan Stok")
		fmt.Println("4. Tampilkan Barang Berdasarkan Urutan Nama")
		fmt.Println("5. Cari Barang Berdasarkan Nama")
		fmt.Println("6. Cari Barang Berdasarkan ID")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahBarang()
		} else if pilihan == 2 {
			hapusBarang()
		} else if pilihan == 3 {
			fmt.Print("1. Sedikit ke Banyak (Asc)\n2. Banyak ke Sedikit (Desc)\nPilihan: ")
			fmt.Scan(&urutPilihan)

			// Perbaikan: Validasi input 1 atau 2
			if urutPilihan == 1 {
				sortByStok_Selection(true)
				cetakData()
			} else if urutPilihan == 2 {
				sortByStok_Selection(false)
				cetakData()
			} else {
				fmt.Println("Pilihan tidak valid! Kembali ke menu utama.")
			}

		} else if pilihan == 4 {
			fmt.Print("1. A ke Z (Asc)\n2. Z ke A (Desc)\nPilihan: ")
			fmt.Scan(&urutPilihan)
			if urutPilihan == 1 {
				sortByNama_Insertion(true)
				cetakData()
			} else if urutPilihan == 2 {
				sortByNama_Insertion(false)
				cetakData()
			} else {
				fmt.Println("Pilihan tidak valid! Kembali ke menu utama.")
			}
		} else if pilihan == 5 {
			var searchNm string
			var idx int
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&searchNm)
			idx = cariNama_Sequential(searchNm)
			if idx != -1 {
				fmt.Println("Ditemukan! Stok:", dataBarang[idx].Stok)
			} else {
				fmt.Println("Tidak ditemukan.")
			}
		} else if pilihan == 6 {
			var searchID, idx int
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&searchID)
			idx = cariID_Binary(searchID)
			if idx != -1 {
				fmt.Println("Ditemukan! Nama:", dataBarang[idx].Nama)
			} else {
				fmt.Println("Tidak ditemukan.")
			}
		} else if pilihan == 0 {
			fmt.Println("Program Selesai")
			lanjut = false
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}

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

func urutID_Asc() {
	// spesifikasi: mengurutkan data berdasarkan ID secara Ascending menggunakan Insertion Sort (untuk mendukung Binary Search)
	var i, j int
	var key Barang
	i = 1
	for i < nBarang {
		key = dataBarang[i]
		j = i - 1
		for j >= 0 && dataBarang[j].ID > key.ID {
			dataBarang[j+1] = dataBarang[j]
			j = j - 1
		}
		dataBarang[j+1] = key
		i = i + 1
	}
}

func cariID_Binary(id int) int {
	var low, mid, high, foundIdx int
	urutID_Asc() //binary search
	low = 0
	high = nBarang - 1
	foundIdx = -1
	for low <= high && foundIdx == -1 {
		mid = (low + high) / 2
		if dataBarang[mid].ID == id {
			foundIdx = mid
		} else if dataBarang[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return foundIdx
}

func hapusBarang() {
	var id, idx, i int
	fmt.Print("Masukkan ID Barang yang akan dihapus: ")
	fmt.Scan(&id)
	idx = cariID_Binary(id)
	if idx != -1 {
		i = idx
		for i < nBarang-1 {
			dataBarang[i] = dataBarang[i+1]
			i = i + 1
		}
		nBarang = nBarang - 1
		fmt.Println("Barang berhasil dihapus.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func cariNama_Sequential(nama string) int {
	var i, foundIdx int
	i = 0
	foundIdx = -1
	for i < nBarang && foundIdx == -1 {
		if dataBarang[i].Nama == nama {
			foundIdx = i
		}
		i = i + 1
	}
	return foundIdx
}

func ubahBarang() {
	var id, idx int
	fmt.Print("Masukkan ID Barang yang akan diubah: ")
	fmt.Scan(&id)

	// Menggunakan Binary Search sesuai permintaan poin (d)
	idx = cariID_Binary(id)

	if idx != -1 {
		fmt.Println("Data ditemukan. Masukkan data baru:")
		fmt.Print("Nama Baru: ")
		fmt.Scan(&dataBarang[idx].Nama)
		fmt.Print("Kategori Baru: ")
		fmt.Scan(&dataBarang[idx].Kategori)
		fmt.Print("Harga Baru: ")
		fmt.Scan(&dataBarang[idx].Harga)
		fmt.Println("Data berhasil diperbarui.")
	} else {
		fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
	}
}

func sortByNama_Insertion(ascending bool) {
	// Spesifikasi: Mengurutkan data berdasarkan nama menggunakan Insertion Sort
	var i, j int
	var key Barang
	i = 1
	for i < nBarang {
		key = dataBarang[i]
		j = i - 1
		if ascending {
			for j >= 0 && dataBarang[j].Nama > key.Nama {
				dataBarang[j+1] = dataBarang[j]
				j = j - 1
			}
		} else {
			for j >= 0 && dataBarang[j].Nama < key.Nama {
				dataBarang[j+1] = dataBarang[j]
				j = j - 1
			}
		}
		dataBarang[j+1] = key
		i = i + 1
	}
}

func sortByStok_Selection(ascending bool) {
	// Spesifikasi: Mengurutkan data berdasarkan stok menggunakan Selection Sort
	var i, j, idx_extreme int
	var temp Barang
	i = 0
	for i < nBarang-1 {
		idx_extreme = i
		j = i + 1
		for j < nBarang {
			if ascending {
				if dataBarang[j].Stok < dataBarang[idx_extreme].Stok {
					idx_extreme = j
				}
			} else {
				if dataBarang[j].Stok > dataBarang[idx_extreme].Stok {
					idx_extreme = j
				}
			}
			j = j + 1
		}
		temp = dataBarang[idx_extreme]
		dataBarang[idx_extreme] = dataBarang[i]
		dataBarang[i] = temp
		i = i + 1
	}
}
