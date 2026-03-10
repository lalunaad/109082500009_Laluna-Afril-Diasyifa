package main

import "fmt"

func main() {
	var beratparsel, kg, sisaberat, biaya int
	var biayatambahan, totalbiaya int

	fmt.Print("Masukkan total berat (gram) : ")
	fmt.Scan(&beratparsel)
	fmt.Println()

	kg = beratparsel / 1000
	sisaberat = beratparsel % 1000

	biaya = kg * 10000

	if kg > 10 {
		biayatambahan = 0
	} else {
		if sisaberat >= 500 {
			biayatambahan = sisaberat * 5
		} else if sisaberat > 0 && sisaberat < 500 {
			biayatambahan = sisaberat * 15
		} else {
			biayatambahan = 0
		}
	}

	totalbiaya = biaya + biayatambahan

	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %d kg + %d gram\n", kg, sisaberat)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biaya, biayatambahan)
	fmt.Printf("Total biaya : Rp %d\n", totalbiaya)
}