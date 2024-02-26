package main

import "fmt"

func main() {
	sabores := []string{"pepperoni", "mozzarella", "abacaxi", "quatro queijos", "marquerita"}

	fatia := sabores[:]
	fatia2 := sabores[2:]

	fmt.Println(fatia)
	fmt.Println(fatia2)

	for i, valor := range sabores {
		fmt.Println(i, ":", valor)
	}
	fmt.Println()
	for i := 0; i < len(fatia); i++ {
		fmt.Println(fatia[i])
	}

}
