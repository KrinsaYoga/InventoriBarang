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
	var pilihan, urutPilihan, subPilihan, idx int
	var lanjut bool
	var katCari, searchNm string
	var searchID int

	isiDataDummy()
	lanjut = true

	for lanjut {
		fmt.Println("\n=== APLIKASI INVENTORI BARANG ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Ubah Data Barang (Edit)")
		fmt.Println("4. Menu Pencarian & Tampilan Barang")
		fmt.Println("5. Catat Transaksi (Masuk/Keluar)")
		fmt.Println("6. Lihat History Transaksi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahBarang()
		case 2:
			hapusBarang()
		case 3:
			ubahBarang()
		case 4:
			// SUB-MENU MENGGUNAKAN SWITCH
			fmt.Println("\n--- MENU PENCARIAN & TAMPILAN ---")
			fmt.Println("1. Cari Nama (Sequential)")
			fmt.Println("2. Cari ID (Binary)")
			fmt.Println("3. Cari Kategori (Sequential)")
			fmt.Println("4. Urutkan Stok (Selection)")
			fmt.Println("5. Urutkan Nama (Insertion)")
			fmt.Print("Pilih: ")
			fmt.Scan(&subPilihan)

			switch subPilihan {
			case 1:
				fmt.Print("Nama Barang: ")
				fmt.Scan(&searchNm)
				idx = cariNama_Sequential(searchNm)
				if idx != -1 {
					fmt.Println("Ketemu! Nama:", dataBarang[idx].Nama, "| Stok:", dataBarang[idx].Stok)
				} else {
					fmt.Println("Tidak ditemukan.")
				}
			case 2:
				fmt.Print("ID Barang: ")
				fmt.Scan(&searchID)
				idx = cariID_Binary(searchID)
				if idx != -1 {
					fmt.Println("Ketemu! Nama:", dataBarang[idx].Nama, "| Stok:", dataBarang[idx].Stok)
				} else {
					fmt.Println("Tidak ditemukan.")
				}
			case 3:
				var i int
				var ketemu bool
				fmt.Print("Kategori: ")
				fmt.Scan(&katCari)
				ketemu = false
				fmt.Printf("\n--- Hasil Kategori: %s ---\n", katCari)
				for i = 0; i < nBarang; i++ {
					if dataBarang[i].Kategori == katCari {
						fmt.Printf("ID: %d | Nama: %s | Stok: %d\n", dataBarang[i].ID, dataBarang[i].Nama, dataBarang[i].Stok)
						ketemu = true
					}
				}
				if !ketemu {
					fmt.Println("Kategori tidak ditemukan.")
				}
			case 4:
				fmt.Print("1. Asc (Sedikit-Banyak)  2. Desc (Banyak-Sedikit): ")
				fmt.Scan(&urutPilihan)
				switch urutPilihan {
				case 1:
					sortByStok_Selection(true)
					cetakData()
				case 2:
					sortByStok_Selection(false)
					cetakData()
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			case 5:
				fmt.Print("1. Asc (A-Z)  2. Desc (Z-A): ")
				fmt.Scan(&urutPilihan)
				switch urutPilihan {
				case 1:
					sortByNama_Insertion(true)
					cetakData()
				case 2:
					sortByNama_Insertion(false)
					cetakData()
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			default:
				fmt.Println("Pilihan sub-menu tidak valid.")
			}
		case 5:
			catatTransaksi()
		case 6:
			fmt.Print("Urut Jml Transaksi? 1. Asc  2. Desc  3. Tidak: ")
			fmt.Scan(&subPilihan)
			switch subPilihan {
			case 1:
				sortHistory(true)
			case 2:
				sortHistory(false)
			}
			tampilkanHistory()
		case 0:
			fmt.Println("Terima kasih! Program selesai.")
			lanjut = false
		default:
			fmt.Println("Menu tidak tersedia.")
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
	var idBaru, idx int

	if nBarang < NMAX {
		fmt.Print("Masukkan ID Barang Baru: ")
		fmt.Scan(&idBaru)

		// Cek apakah ID sudah ada menggunakan fungsi yang sudah kita buat
		idx = cariID_Binary(idBaru)

		if idx != -1 {
			// Jika idx bukan -1, berarti ID sudah ditemukan di sistem
			fmt.Println("Error: ID sudah digunakan oleh barang:", dataBarang[idx].Nama)
			fmt.Println("Gagal menambah data. Gunakan ID lain.")
		} else {
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
		}
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

func cariKategori_Sequential(kat string) {
	var i int
	var ditemukan bool
	i = 0
	ditemukan = false
	fmt.Printf("\n--- Hasil Pencarian Kategori: %s ---\n", kat)
	for i < nBarang {
		if dataBarang[i].Kategori == kat {
			fmt.Printf("ID: %d | Nama: %s | Stok: %d\n", dataBarang[i].ID, dataBarang[i].Nama, dataBarang[i].Stok)
			ditemukan = true
		}
		i = i + 1
	}
	if !ditemukan {
		fmt.Println("Tidak ada barang dalam kategori tersebut.")
	}
}

func ubahBarang() {
	var id, idx int
	fmt.Print("Masukkan ID Barang yang akan diubah: ")
	fmt.Scan(&id)

	// Menggunakan Binary Search
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

func catatTransaksi() {
	var id, idx, jml int
	var tipe string
	fmt.Print("Masukkan ID Barang: ")
	fmt.Scan(&id)
	idx = cariID_Binary(id)
	if idx != -1 {
		fmt.Print("Tipe Transaksi (Masuk/Keluar): ")
		fmt.Scan(&tipe)
		fmt.Print("Jumlah Barang: ")
		fmt.Scan(&jml)

		if tipe == "Masuk" {
			dataBarang[idx].Stok = dataBarang[idx].Stok + jml
			// PERBAIKAN: Mengisi IDTransaksi agar tidak 0
			dataTransaksi[nTransaksi].IDTransaksi = nTransaksi + 1
			dataTransaksi[nTransaksi].IDBarang = id
			dataTransaksi[nTransaksi].Tipe = "Masuk"
			dataTransaksi[nTransaksi].Jumlah = jml
			nTransaksi = nTransaksi + 1
			fmt.Println("Transaksi masuk berhasil dicatat.")
		} else if tipe == "Keluar" {
			if dataBarang[idx].Stok >= jml {
				dataBarang[idx].Stok = dataBarang[idx].Stok - jml
				// PERBAIKAN: Mengisi IDTransaksi agar tidak 0
				dataTransaksi[nTransaksi].IDTransaksi = nTransaksi + 1
				dataTransaksi[nTransaksi].IDBarang = id
				dataTransaksi[nTransaksi].Tipe = "Keluar"
				dataTransaksi[nTransaksi].Jumlah = jml
				nTransaksi = nTransaksi + 1
				fmt.Println("Transaksi keluar berhasil dicatat.")
			} else {
				fmt.Println("Stok tidak mencukupi!")
			}
		} else {
			fmt.Println("Tipe tidak valid.")
		}
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

func sortHistory(ascending bool) {
	/* Spesifikasi: Mengurutkan riwayat transaksi berdasarkan JUMLAH
	   menggunakan Insertion Sort */
	var i, j int
	var key Transaksi
	i = 1
	for i < nTransaksi {
		key = dataTransaksi[i]
		j = i - 1
		if ascending {
			for j >= 0 && dataTransaksi[j].Jumlah > key.Jumlah {
				dataTransaksi[j+1] = dataTransaksi[j]
				j = j - 1
			}
		} else {
			for j >= 0 && dataTransaksi[j].Jumlah < key.Jumlah {
				dataTransaksi[j+1] = dataTransaksi[j]
				j = j - 1
			}
		}
		dataTransaksi[j+1] = key
		i = i + 1
	}
}

func tampilkanHistory() {
	var i int
	fmt.Println("\n--- RIWAYAT TRANSAKSI (LOG) ---")
	fmt.Printf("%-5s %-10s %-10s %-10s\n", "No", "ID Brg", "Tipe", "Jumlah")
	i = 0
	for i < nTransaksi {
		fmt.Printf("%-5d %-10d %-10s %-10d\n",
			dataTransaksi[i].IDTransaksi,
			dataTransaksi[i].IDBarang,
			dataTransaksi[i].Tipe,
			dataTransaksi[i].Jumlah)
		i = i + 1
	}
}
