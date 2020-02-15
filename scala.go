package main

import (
	"fmt"
)

func reset(x int) int {
	if x == 10 {
		x = 0
	}
	return x
}

func main() {
	var n, l int
	K := 0
	skipper := 0
	Molt := 0
	fmt.Print("inserire numero delle scale e lunghezza ")
	fmt.Scanln(&n, &l)

	for i := 1; i <= 2*n; i++ {
		if i%2 != 0 && i == 1 {
			for j := 1; j < 2*l; j++ {
				if j%2 != 0 {
					fmt.Print(K)
					K++
					K = reset(K)
				} else {
					fmt.Print("-")
				}
			}
			Molt = 1
		} else if i%2 == 0 {
			for t := 1; t < skipper; t++ {
				if t == skipper-1 {
					fmt.Print(K)
					K++
					K = reset(K)
				} else {
					fmt.Print(" ")
				}
			}
		} else if i%2 != 0 && i != 1 {
			for j := 1; j < skipper+(2*l); j++ {
				if j < skipper-1 {
					fmt.Print(" ")
				} else if j%2 != 0 {
					fmt.Print(K)
					K++
					K = reset(K)
				} else {
					fmt.Print("-")
				}
			}
			Molt++
		}
		skipper = Molt * (2 * l)
		fmt.Println()
	}
	return
}
