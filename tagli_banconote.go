package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	tagli := []int{100, 50, 20, 10, 5, 2, 1}
	numeri, _ := strconv.Atoi(n)
	somma := esecuzione(numeri)
	for i := 0; i < len(somma); i++ {
		fmt.Printf("Taglio : %d  Quantita' : %d \n", tagli[i], somma[i])
	}

}
func esecuzione(numeri int) []int {
	tagli := []int{100, 50, 20, 10, 5, 2, 1}
	var control []int
	for i := 0; i < len(tagli); i++ {
		controllo := numeri / tagli[i]
		numeri = numeri % tagli[i]
		control = append(control, controllo)
	}
	return control
}
