package main

import "fmt"

func main() {
    var votes [21]int 
    var input int
    var totalVotes int  
    var validVotes int 
    

    for {
        fmt.Scan(&input)
        if input == 0 {
            break
        }
        totalVotes++
   
        if input >= 1 && input <= 20 {
            votes[input]++
            validVotes++
        }
    }
    
    fmt.Println("Suara masuk:", totalVotes)
    fmt.Println("Suara sah:", validVotes)
    
    maxVotes := -1
    for i := 1; i <= 20; i++ {
        if votes[i] > maxVotes {
            maxVotes = votes[i]
        }
    }
    

    ketua := -1
    wakil := -1
    
    for i := 1; i <= 20; i++ {
        if votes[i] == maxVotes {
            if ketua == -1 {
                ketua = i
            } else if wakil == -1 {
                wakil = i
                break
            }
        }
    }
    
    fmt.Println("Ketua RT:", ketua)
    fmt.Println("Wakil ketua:", wakil)
}