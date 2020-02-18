package main

import (
	"bufio"
	"os"
	"math/rand"
	"time"
	"strconv"
	"math"
)

func main() {
	var slice []string
	//Bufio
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo := scanner.Text()
		slice = append(slice, testo)
	}
	//Conta ripetizioni in una stringa
	cont := strings.Count(numeri, K)

	//Da slice a stringa
	lettere := strings.Join(lettereSlice, "")

	//Passaggio a removeCharacters
	lettere = removeCharacters(lettere, K)

	//Ripetere caratteri
	fmt.Println(K, strings.Repeat("*", cont))

	//Splittare una stringa
	stringa := strings.Split(stringaDaSplittare, ";")

	//Conversione in float64
	nfloat, _ := strconv.ParseFloat(numero, 64)

	//Conversione in Int
	numeri, _ := strconv.Atoi(n)

	//Da runa a stringa
	ch := string(char)

	//Per uscire dal programma
	os.Exit(0)

	//Elevare al quadrato
	result := math.Pow(float64(x), 2)

	//Radice quadrata
	radice := math.Sqrt(result)

	//Numero randomico
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
}
//Rimozione caratteri da una stringa
func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)

}
//Creazione e gestione delle struct
type Point struct {
	x float64
	y float64
	z float64
}
func main(){
	var p4 Point
	p1 := Point{1.3, 2.5, -2}
	p2 := Point{-3, -2.5, -2}
	p3 := Point{1.1, -3.7, 4}
	proiezione(&p1,&p2)
	s := []Point3{p1, p2, p3}
	//Dichiarazione variabile della struct
	fmt.Println("inserisci x, y, raggio del primo cerchio")
	fmt.Scan(&p4.x, &p4.y, &p4.r)
}
func proiezione(str, str2 *Point) {
	str.z = 0
	str2.z = 0
}
