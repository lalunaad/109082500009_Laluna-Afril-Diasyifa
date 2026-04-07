package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	barisanBilangan(n)
}

func barisanBilangan(bilangan int) {
	if bilangan == 1 {
		fmt.Print(1, " ")
	} else {
		fmt.Print(bilangan, " ")
		barisanBilangan(bilangan - 1)
		fmt.Print(bilangan, " ")
	}
}
