package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	fmt.Println("Inserisci la stringa - per uscire scrivi 'end' (puoi andare a capo)")
	slice := inserimentoStringa()
	//Trasformo la slice in una stringa con strings.Join
	line := strings.Join(slice, "")
	lettere, numeri, contn, contl := smistamento(line)

	fmt.Println("Totale Lettere :", contl)
	letterecont(lettere)
	fmt.Println("Totale Cifre :", contn)
	numericont(numeri)

}
func inserimentoStringa() []string {
	var slice []string
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "end" {
			break
		} //Break out of input loop when the user types the word "end"
		text = scanner.Text()
		slice = append(slice, text)
	}
	return slice
}
func smistamento(line string) (string, string, int, int) {
	var letteresl []string
	var numerisl []string
	var contn, contl int
	var numeri, lettere string
	for _, char := range line {

		if unicode.IsLetter(char) {
			ch := string(char)
			letteresl = append(letteresl, ch)
			contl++
		} else if unicode.IsNumber(char) {
			n := string(char)
			numerisl = append(numerisl, n)
			contn++
		}
	}
	numeri = strings.Join(numerisl, "")
	lettere = strings.Join(letteresl, "")

	return lettere, numeri, contn, contl
}

func letterecont(lettere string) {
	for _, s := range lettere {
		K := string(s)

		cont := 0
		cont = strings.Count(lettere, K)
		if cont > 0 {
			//fmt.Println(K, strings.Repeat("*", cont))
			fmt.Println(K, "--->", cont)

		}
		lettere = removeCharacters(lettere, K)
	}
	return
}
func numericont(numeri string) {
	for _, s := range numeri {
		K := string(s)
		cont := 0

		cont = strings.Count(numeri, K)
		if cont > 0 {
			//fmt.Println(K, strings.Repeat("*", cont))
			fmt.Println(K, "--->", cont)

		}
		numeri = removeCharacters(numeri, K)
	}
	return
}
func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)

}
