package main

import (
	"fmt"
	"math"
)

func main() {
	p1, p2 := inserimento()
	distanza := distanza(&p1, &p2)
	sommaraggi := p2.r + p1.r
	if sommaraggi == int(distanza) {
		fmt.Println("Sono tangenti")
	} else {
		fmt.Println("Non sono tangenti")
	}
}
func inserimento() (point, point) {
	var p1, p2 point
	fmt.Println("inserisci x, y, raggio del primo cerchio")
	fmt.Scan(&p1.x, &p1.y, &p1.r)
	fmt.Println("inserisci x, y, raggio del primo cerchio")
	fmt.Scan(&p2.x, &p2.y, &p2.r)
	return p1, p2
}

type point struct {
	x int
	y int
	r int
}

func distanza(p1, p2 *point) float64 {
	x := (p2.x - p1.x)
	resultx := math.Pow(float64(x), 2)
	y := (p2.y - p1.y)
	resulty := math.Pow(float64(y), 2)
	result := resultx + resulty
	distanza := math.Sqrt(result)
	return distanza

}
