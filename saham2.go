package main

import (
	"fmt"
)

const MAKS_SAHAM = 4

type Saham struct {
	Kode, Nama    string
	Harga, Volume int
}

type Transaksi struct {
	ID        int
	Kode      string
	Lot       int
	Tipe      string
	HargaSesi int
}

var daftarPasar = [MAKS_SAHAM]Saham{
	{"BBCA", "Bank Central Asia", 9500, 500},
	{"GOTO", "GoTo", 50, 10000},
	{"TLKM", "Telkom", 3800, 1200},
	{"ASII", "Astra", 5200, 2000},
}

var porto [100]Transaksi
var jumlahPorto = 0
var duitSaldo = 10000000
var idTransaksi = 1
var sirkulasi = 0

func main() {
	var selesai bool
	selesai = false

	for selesai == false {
		fmt.Println("\n--- MENU SIMULASI SAHAM ---")
		fmt.Println("1. Transaksi (Beli/Jual)")
		fmt.Println("2. Cek Portofolio & Statistik")
		fmt.Println("3. Cari Saham")
		fmt.Println("4. Urutkan Saham")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			transaksiBaru()
		case 2:
			cekStatistik()
		case 3:
			cariSaham()
		case 0:
			selesai = true
		}
	}
}

func transaksiBaru() {
	fmt.Print("\n1. Tambah  2. Ubah Lot  3. Hapus  0. Kembali | Pilih: ")
	var p int
	fmt.Scan(&p)

	if p == 0 {
		return
	}

	if p == 1 {
		if jumlahPorto >= 100 {
			fmt.Println("Kapasitas riwayat transaksi penuh.")
			return
		}

		var kode, tipe string
		var jmlLot int
		fmt.Print("Masukkan Kode, Tipe (Beli/Jual), dan Lot: ")
		fmt.Scan(&kode, &tipe, &jmlLot)

		h := 0
		for i := 0; i < MAKS_SAHAM; i++ {
			if daftarPasar[i].Kode == kode {
				h = daftarPasar[i].Harga
			}
		}

		porto[jumlahPorto] = Transaksi{idTransaksi, kode, jmlLot, tipe, h}
		jumlahPorto++

		if tipe == "Beli" {
			duitSaldo -= h * jmlLot * 100
		} else {
			duitSaldo += h * jmlLot * 100
		}
		fmt.Println("Transaksi berhasil. ID:", idTransaksi)
		idTransaksi++

	} else if p == 2 {
		var id, lotBaru int
		fmt.Print("Masukkan ID Transaksi dan Lot baru: ")
		fmt.Scan(&id, &lotBaru)

		ketemu := false
		for i := 0; i < jumlahPorto && ketemu == false; i++ {
			if porto[i].ID == id {
				selisihLot := lotBaru - porto[i].Lot
				nilai := porto[i].HargaSesi * selisihLot * 100

				if porto[i].Tipe == "Beli" {
					duitSaldo -= nilai
				} else {
					duitSaldo += nilai
				}
				porto[i].Lot = lotBaru
				fmt.Println("Data lot telah diperbarui.")
				ketemu = true
			}
		}

	} else if p == 3 {
		var id int
		fmt.Print("Masukkan ID yang akan dihapus: ")
		fmt.Scan(&id)

		ketemu := false
		for i := 0; i < jumlahPorto && ketemu == false; i++ {
			if porto[i].ID == id {
				if porto[i].Tipe == "Beli" {
					duitSaldo += porto[i].HargaSesi * porto[i].Lot * 100
				} else {
					duitSaldo -= porto[i].HargaSesi * porto[i].Lot * 100
				}

				for j := i; j < jumlahPorto-1; j++ {
					porto[j] = porto[j+1]
				}
				jumlahPorto--
				fmt.Println("Data berhasil dihapus.")
				ketemu = true
			}
		}
	}
}

func cekStatistik() {
	sirkulasi++
	for i := 0; i < MAKS_SAHAM; i++ {
		if (i+sirkulasi)%2 == 0 {
			daftarPasar[i].Harga += 100
		} else {
			daftarPasar[i].Harga -= 100
		}
	}

	fmt.Println("\nSaldo Saat Ini:", duitSaldo)
	totLR := 0

	for i := 0; i < jumlahPorto; i++ {
		t := porto[i]
		hargaKini := 0
		for j := 0; j < MAKS_SAHAM; j++ {
			if daftarPasar[j].Kode == t.Kode {
				hargaKini = daftarPasar[j].Harga
			}
		}

		lr := 0
		if t.Tipe == "Beli" {
			lr = (hargaKini - t.HargaSesi) * t.Lot * 100
		}

		totLR += lr
		fmt.Printf("[%s] Tipe: %s | Lot: %d | Harga Awal: %d | Harga Kini: %d | L/R: %d\n",
			t.Kode, t.Tipe, t.Lot, t.HargaSesi, hargaKini, lr)
	}
	fmt.Println("Total Keuntungan/Kerugian:", totLR)
}

func cariSaham() {
	fmt.Println("\n--- DAFTAR SAHAM TERSEDIA ---")
	for i := 0; i < MAKS_SAHAM; i++ {
		fmt.Printf("- %s | %s\n", daftarPasar[i].Kode, daftarPasar[i].Nama)
	}

	fmt.Print("\n1. Cari Nama (Sequential)  2. Cari Kode (Binary) | Pilih: ")
	var p int
	fmt.Scan(&p)

	ketemu := false

	if p == 1 {
		var cari string
		fmt.Print("Masukkan Nama Perusahaan: ")
		fmt.Scan(&cari)

		i := 0
		for i < MAKS_SAHAM && ketemu == false {
			if daftarPasar[i].Nama == cari {
				fmt.Printf("\n[HASIL DITEMUKAN]\nKode: %s\nNama: %s\nHarga: %d\nVolume: %d\n", 
					daftarPasar[i].Kode, daftarPasar[i].Nama, daftarPasar[i].Harga, daftarPasar[i].Volume)
				ketemu = true
			}
			i++
		}
	} else if p == 2 {
		var kodeCari string
		fmt.Print("Masukkan Kode Saham: ")
		fmt.Scan(&kodeCari)

		for i := 0; i < MAKS_SAHAM-1; i++ {
			for j := i + 1; j < MAKS_SAHAM; j++ {
				if daftarPasar[i].Kode > daftarPasar[j].Kode {
					temp := daftarPasar[i]
					daftarPasar[i] = daftarPasar[j]
					daftarPasar[j] = temp
				}
			}
		}

		kiri, kanan := 0, MAKS_SAHAM-1
		for kiri <= kanan && ketemu == false {
			tengah := (kiri + kanan) / 2
			if daftarPasar[tengah].Kode == kodeCari {
				fmt.Printf("\n[HASIL DITEMUKAN]\nKode: %s\nNama: %s\nHarga: %d\nVolume: %d\n", 
					daftarPasar[tengah].Kode, daftarPasar[tengah].Nama, daftarPasar[tengah].Harga, daftarPasar[tengah].Volume)
				ketemu = true
			} else if daftarPasar[tengah].Kode < kodeCari {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
	}

	if ketemu == false {
		fmt.Println("Maaf, data tidak ditemukan.")
	}
}