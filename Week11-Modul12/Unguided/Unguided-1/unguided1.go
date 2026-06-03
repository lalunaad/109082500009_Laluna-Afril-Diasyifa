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
    
    
    for i := 1; i <= 20; i++ {
        if votes[i] > 0 {
            fmt.Printf("%d: %d\n", i, votes[i])
        }
    }
}