package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var tentativo string
	var sconfitte, vittorie int
	for {
		rand := random()
		r := game(rand)
		fmt.Println("Dove sto guardando?")
		fmt.Scan(&tentativo)
		if tentativo == "0" {
			break
		}
		if tentativo == r {
			fmt.Println("Giusto, hai vinto!")
			vittorie++
		} else {
			fmt.Println("Sbagliato, hai perso!")
			sconfitte++
		}

	}
	pVittorie := vittorie * 100 / (vittorie + sconfitte)
	pSconfitte := 100 - pVittorie
	fmt.Println("Vinto: ", pVittorie, "%", "Perso: ", pSconfitte, "%")

}
func random() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	return random
}
func game(random int) string {
	var r string
	if random%2 == 0 {
		r = "d"
	} else {
		r = "s"
	}
	return r
}
