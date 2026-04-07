package main

import "fmt"

func main() {
	var jumlahBaris int
	fmt.Scan(&jumlahBaris)
	cetakPola(1, jumlahBaris)
}

func cetakPola(barisSekarang, batasBaris int) {
	if barisSekarang > batasBaris {
		return
	} else {
		cetakBintang(barisSekarang)
		fmt.Println()
		cetakPola(barisSekarang + 1, batasBaris)
	}
}

func cetakBintang(banyakBintang int) {
	if banyakBintang == 0 {
		return
	} else {
		fmt.Print("*")
		cetakBintang(banyakBintang - 1)
	}
}