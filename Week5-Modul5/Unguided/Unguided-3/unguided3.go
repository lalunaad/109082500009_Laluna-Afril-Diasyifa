package main

import "fmt"

func main() {
	var basis, eksponen int
	fmt.Scan(&basis, &eksponen)
	fmt.Println(hitungPangkat(basis, eksponen))
}

func hitungPangkat(basis, eksponen int) int {
	if eksponen == 0 {
		return 1
	} else {
		return basis * hitungPangkat(basis, eksponen-1)
	}
}