package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := os.Args[1:]
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			numero, _ := strconv.Atoi(string(str[i+1]))
			fmt.Print(strings.Repeat(string(str[i]), numero))
		}
	}

}
