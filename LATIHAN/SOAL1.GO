package main

import "fmt"

const NMAX = 1000000
type arrInt [NMAX]int

func sorting(T *arrInt, n int) {
    var langkah, minIdx, j, tmp int
    for langkah = 1; langkah <= n-1; langkah++ {
        minIdx = langkah - 1
        for j = langkah; j <= n-1; j++ {
            if T[minIdx] > T[j] {
                minIdx = j
            }
        }
        tmp = T[minIdx]
        T[minIdx] = T[langkah-1]
        T[langkah-1] = tmp
    }
}

func median(T arrInt, n int) float64 {
    var tengah int = n / 2
    if n%2 == 0 {
        return float64(T[tengah-1] + T[tengah]) / 2.0
    } else {
        return float64(T[tengah])
    }
}

func main() {
    var data arrInt
    var val, count int
    
    fmt.Scan(&val)
    count = 0
    
    for val != -5313541 && count < NMAX {
        if val == 0 {
            sorting(&data, count)
            fmt.Println(median(data, count))
        } else {
            data[count] = val
            count++
        }
        fmt.Scan(&val)
    }
}