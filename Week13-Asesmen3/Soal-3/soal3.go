package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func main() {
	var p tabPartai
	var n, inputVal, indeksPartai int

	fmt.Println("Masukkan proses input suara :")  
	fmt.Scan(&inputVal)
	for inputVal != -1 {
		indeksPartai = posisi(p, n, inputVal)
		if indeksPartai == -1 {
			p[n].nama = inputVal
			p[n].suara = 1
			n++
		} else {
			p[indeksPartai].suara++
		}
		fmt.Scan(&inputVal)
	}

	var langkah, j int
	var tmp partai
	for langkah = 1; langkah <= n-1; langkah++ {
		j = langkah
		tmp = p[j]
		for j > 0 && tmp.suara > p[j-1].suara {
			p[j] = p[j-1]
			j--
		}
		p[j] = tmp
	}

	fmt.Println("Hasil Perhitungan suara :")  
	for i := 0; i < n; i++ {
		fmt.Printf("%v(%v) ", p[i].nama, p[i].suara)
	}
	fmt.Println()  
}

func posisi(t tabPartai, n int, nama int) int {
	var counter, ditemukan int
	counter = 0
	ditemukan = -1
	for counter < n && ditemukan == -1 {
		if t[counter].nama == nama {
			ditemukan = counter
		}
		counter++
	}
	return ditemukan
}