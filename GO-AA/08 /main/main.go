package main

import "fmt"

func main() {
	primeiroSlice := []string{"banana", "maça", "jaca"}
	primeiroSlice = append(primeiroSlice, "pessego")

	for indice, valor := range primeiroSlice {
		fmt.Println("No indice", indice, "temos o valor: ", valor)
	}
	fmt.Println()
	for indice, _ := range primeiroSlice {
		fmt.Println("esse slice tem", indice, "elementos")
	}

}
