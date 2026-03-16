package main

import "fmt"

func hitungBiaya(jenis string, masuk int, keluar int) int {
	var totalBiaya int
	var jamAkhir int
	var jam int
	var jamSaatIni int
	var tarif int

	jamAkhir = keluar
	if jamAkhir < masuk {
		jamAkhir = jamAkhir + 24
	}

	for jam = masuk; jam < jamAkhir; jam = jam + 1 {
		jamSaatIni = jam % 24
		tarif = 0

		if jamSaatIni >= 0 && jamSaatIni < 17 {
			if jenis == "motor" {
				tarif = 4000
			} else {
				tarif = 6000
			}
		} else {
			if jenis == "motor" {
				tarif = 5000
			} else {
				tarif = 7000
			}
		}

		totalBiaya = totalBiaya + tarif
	}

	return totalBiaya
}

func main() {
	var totalPendapatan int
	var kendaraanKe int
	var jenis string
	var jamMasuk int
	var jamKeluar int
	var biaya int

	kendaraanKe = 1

	fmt.Println("=== Rekap Tarif Parkir Cafe Per Hari ===")

	for true {
		fmt.Printf("*Kendaraan %d\n", kendaraanKe)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai) : ")
		fmt.Scanln(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24) : ")
		fmt.Scanln(&jamMasuk)

		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24) : ")
		fmt.Scanln(&jamKeluar)

		biaya = hitungBiaya(jenis, jamMasuk, jamKeluar)

		fmt.Printf("Biaya parkir %s %d : %d\n", jenis, kendaraanKe, biaya)
		fmt.Println("==========================================")

		totalPendapatan = totalPendapatan + biaya
		kendaraanKe = kendaraanKe + 1
	}

	fmt.Printf("*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", totalPendapatan)
}