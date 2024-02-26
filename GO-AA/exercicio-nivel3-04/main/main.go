package main

import "fmt"

func main() {
	anoNasceu := 1990
	anoAtual := 2024
	for {
		if anoNasceu > anoAtual {
			break
		}
		fmt.Println(anoNasceu)
		anoNasceu++
	}
}
