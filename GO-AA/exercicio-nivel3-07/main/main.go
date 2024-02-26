package main

import "fmt"

func main() {
	idade := 66

	if idade < 18 {
		fmt.Println("Menor de idade")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Adulto")
	} else {
		fmt.Println("Idoso")
	}
}
