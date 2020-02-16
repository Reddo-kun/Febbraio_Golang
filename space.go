package main

import "fmt"

type Point3 struct {
	x float64
	y float64
	z float64
}

func main() {

	p1 := Point3{1.3, 2.5, -2}
	p2 := Point3{-3, -2.5, -2}
	p3 := Point3{1.1, -3.7, 4}
	proiezione(&p1, &p2)
	fmt.Print("p1: ", p1, "\n", "p2: ", p2, "\n", "p3: ", p3, "\n")
	punti2D(p1, p2, p3)
}
func proiezione(str, str2 *Point3) {
	str.z = 0
	str2.z = 0
}

func punti2D(p1, p2, p3 Point3) {
	s := []Point3{p1, p2, p3}
	for _, f := range s {
		if f.z == 0 {
			fmt.Println("punti con z = 0", f)
		}

	}
}
