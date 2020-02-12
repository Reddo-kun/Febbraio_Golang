package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	lista := inserimento()
	minore := minDispari(lista)
	result := superiore(minore, lista)
	fmt.Println(result)
}
func inserimento() []int {
	var listaNumeri []int
	stringa := os.Args[1:]
	var cont int
	for _, numero := range stringa {

		n, _ := strconv.Atoi(numero)

		if n%2 != 0 {
			cont++
		}

		listaNumeri = append(listaNumeri, n)
	}
	if cont == 0 {
		fmt.Println("Non sono presenti numeri dispari")
		os.Exit(0)
	}
	return listaNumeri
}
func minDispari(intSlice []int) int {
	var cont, dispari int
	for _, numero := range intSlice {
		if numero%2 != 0 && cont == 0 {
			dispari = numero
			cont++
		}
		if numero < dispari {
			dispari = numero
		}
	}
	return dispari

}
func superiore(minoreDispari int, Slice []int) []int {
	var risultato []int
	for _, pari := range Slice {
		if pari%2 == 0 && pari > minoreDispari {
			risultato = append(risultato, pari)

		}
	}
	return risultato

}
