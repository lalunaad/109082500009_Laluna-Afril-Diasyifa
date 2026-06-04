package main

import "fmt"

type player struct {
    nama   string
    gol    int
    assist int
}

const NMAX = 1000
type Arrplayer [NMAX]player

func main() {
    var data Arrplayer
    var idx, total int
    var depan, belakang string

    fmt.Println("Masukkan Data Input :") 
    fmt.Scan(&total)
    for idx = 0; idx < total && idx < NMAX; idx++ {
        fmt.Scan(&depan, &belakang, &data[idx].gol, &data[idx].assist)
        data[idx].nama = depan + " " + belakang
    }

    var langkah, idxMax int
    var tmp player

    for langkah = 1; langkah <= total-1; langkah++ {
        idxMax = langkah - 1
        for idx = langkah; idx < total; idx++ {
            if data[idxMax].gol < data[idx].gol || (data[idxMax].gol == data[idx].gol && data[idxMax].assist <= data[idx].assist) {
                idxMax = idx
            }
        }
        tmp = data[idxMax]
        data[idxMax] = data[langkah-1]
        data[langkah-1] = tmp
    }

    fmt.Println("\nHasil Sorting :")  
    for idx = 0; idx < total; idx++ {
        fmt.Println(data[idx].nama, data[idx].gol, data[idx].assist)
    }
}