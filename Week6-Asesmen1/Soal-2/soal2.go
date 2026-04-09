package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		return 3 
	} else if tujuan == "mancanegara" {
		return 8
	}
	return 0
}

func biayaPerHari(jumlahMhs int) int {
	makan := 35000
	penginapan := 250000
	return makan + penginapan
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	uangSaku := 300000

	hari := tanggunganHari(lamaPerjalanan, tujuan)
	biayaHarian := biayaPerHari(jumlahMhs)

	total := jumlahMhs * hari * biayaHarian
	total = total + (jumlahMhs * uangSaku)

	*totalBiaya = float64(total)
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	fmt.Println()

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U: Rp %.2f\n", biaya)
}