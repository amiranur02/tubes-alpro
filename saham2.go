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

var pasar = [MAKS_SAHAM]Saham{
	{"BBCA", "Bank Central Asia", 9500, 500},
	{"GOTO", "GoTo", 50, 10000},
	{"TLKM", "Telkom", 3800, 1200},
	{"ASII", "Astra", 5200, 2000},
}

var portofolio [100]Transaksi
var jumlahPorto = 0 
var saldo = 10000000
var idTx = 1
var siklusHarga = 0 

func main() {
	jalan := true
	for jalan {
		fmt.Println("\n--- MENU SIMULASI SAHAM ---")
		fmt.Println("1. Transaksi Beli/Jual")
		fmt.Println("2. Cek Portofolio & Statistik")
		fmt.Println("3. Cari Saham")
		fmt.Println("4. Urutkan Saham")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		
		var menu int
		fmt.Scan(&menu)

		switch menu {
		case 1:
			menuTransaksi()
		case 2:
			menuStatistik()
		case 3:
			menuCari()
		case 0:
			jalan = false
		}
	}
}

func menuTransaksi() {
	fmt.Print("\n1. Tambah  2. Ubah Lot  3. Hapus  0. Kembali | Pilih: ")
	var p int
	fmt.Scan(&p)

	if p == 0 {
		return
	}

	if p == 1 {
		if jumlahPorto >= 100 {
			fmt.Println("Gagal: Kapasitas riwayat transaksi penuh (Maks 100)!")
			return
		}

		var kode, tipe string
		var lot int
		fmt.Print("Masukkan Kode, Tipe(Beli/Jual), dan Lot (Pisahkan dg spasi): ")
		fmt.Scan(&kode, &tipe, &lot)

		harga := 0
		for i := 0; i < MAKS_SAHAM; i++ {
			if pasar[i].Kode == kode {
				harga = pasar[i].Harga
			}
		}
		
		portofolio[jumlahPorto] = Transaksi{idTx, kode, lot, tipe, harga}
		jumlahPorto++ 
		
		if tipe == "Beli" { 
			saldo -= harga * lot * 100 
		} else { 
			saldo += harga * lot * 100 
		}
		fmt.Println("Transaksi ditambah. ID:", idTx)
		idTx++

	} else if p == 2 {
		var id, lotBaru int
		fmt.Print("Masukkan ID Transaksi dan Lot Baru (Spasi): ")
		fmt.Scan(&id, &lotBaru)

		ketemu := false 
		for i := 0; i < jumlahPorto && !ketemu; i++ {
			if portofolio[i].ID == id {
				if portofolio[i].Tipe == "Beli" {
					saldo += portofolio[i].HargaSesi * portofolio[i].Lot * 100 
					saldo -= portofolio[i].HargaSesi * lotBaru * 100           
				} else if portofolio[i].Tipe == "Jual" {
					saldo -= portofolio[i].HargaSesi * portofolio[i].Lot * 100 
					saldo += portofolio[i].HargaSesi * lotBaru * 100           
				}
				portofolio[i].Lot = lotBaru
				fmt.Println("Transaksi diubah.")
				ketemu = true 
			}
		}

	} else if p == 3 {
		var id int
		fmt.Print("Masukkan ID Transaksi: ")
		fmt.Scan(&id)

		ketemu := false 
		for i := 0; i < jumlahPorto && !ketemu; i++ {
			if portofolio[i].ID == id {
				if portofolio[i].Tipe == "Beli" {
					saldo += portofolio[i].HargaSesi * portofolio[i].Lot * 100 
				} else if portofolio[i].Tipe == "Jual" {
					saldo -= portofolio[i].HargaSesi * portofolio[i].Lot * 100 
				}
				
				for j := i; j < jumlahPorto-1; j++ {
					portofolio[j] = portofolio[j+1]
				}
				jumlahPorto-- 
				fmt.Println("Transaksi dihapus.")
				ketemu = true 
			}
		}
	}
}

func menuStatistik() {
	siklusHarga++
	for i := 0; i < MAKS_SAHAM; i++ {
		if (i+siklusHarga)%2 == 0 {
			pasar[i].Harga += 100 
		} else {
			pasar[i].Harga -= 100 
		}
	}

	fmt.Println("\nSaldo Virtual:", saldo)
	totalLR := 0

	for i := 0; i < jumlahPorto; i++ {
		t := portofolio[i]
		hargaKini := 0
		for j := 0; j < MAKS_SAHAM; j++ {
			if pasar[j].Kode == t.Kode { 
				hargaKini = pasar[j].Harga 
			}
		}

		UntungRugi := (hargaKini - t.HargaSesi) * t.Lot * 100
		if t.Tipe == "Jual" { UntungRugi = 0 } 

		totalLR += UntungRugi
		fmt.Printf("[%s] Tipe: %s | Lot: %d | Beli: %d | Kini: %d | Untung/Rugi: %d\n", t.Kode, t.Tipe, t.Lot, t.HargaSesi, hargaKini, UntungRugi)
	}
	fmt.Println("Total Keuntungan/Kerugian Portofolio:", totalLR)
}

func menuCari() {
	fmt.Println("\n--- DAFTAR SAHAM YANG TERSEDIA ---")
	for i := 0; i < MAKS_SAHAM; i++ {
		fmt.Printf("- Kode: %s | Nama: %s\n", pasar[i].Kode, pasar[i].Nama)
	}

	fmt.Print("\n1. Cari Nama(Sequential)  2. Cari Kode(Binary)  0. Kembali | Pilih: ")
	var p int
	fmt.Scan(&p)

	if p == 0 {
		return
	}

	ketemu := false

	if p == 1 {
		var nama string
		fmt.Print("Ketik Nama Saham (Harus Sama Persis): ")
		fmt.Scan(&nama)

		for i := 0; i < MAKS_SAHAM && !ketemu; i++ {
			if pasar[i].Nama == nama {
				fmt.Printf("\n[HASIL] %s | %s | Harga: %d | Vol: %d\n", pasar[i].Kode, pasar[i].Nama, pasar[i].Harga, pasar[i].Volume)
				ketemu = true 
			}
		}
	} else if p == 2 {
		var kode string
		fmt.Print("Ketik Kode Saham: ")
		fmt.Scan(&kode)

		for i := 0; i < MAKS_SAHAM-1; i++ {
			for j := i + 1; j < MAKS_SAHAM; j++ {
				if pasar[i].Kode > pasar[j].Kode {
					pasar[i], pasar[j] = pasar[j], pasar[i]
				}
			}
		}

		kiri, kanan := 0, MAKS_SAHAM-1
		
		for kiri <= kanan && !ketemu {
			tengah := (kiri + kanan) / 2
			if pasar[tengah].Kode == kode {
				fmt.Printf("\n[HASIL] %s | %s | Harga: %d | Vol: %d\n", pasar[tengah].Kode, pasar[tengah].Nama, pasar[tengah].Harga, pasar[tengah].Volume)
				ketemu = true 
			} else if pasar[tengah].Kode < kode {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
	}

	if !ketemu {
		fmt.Println("Saham tidak ditemukan.")
	}
}
