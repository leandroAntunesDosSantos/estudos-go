package main

import "fmt"

func main() {
	valor1 := make([]int, 2, 4) //make(tipo,tamanho,capacidade)
	valor1[0], valor1[1] = 1, 2

	valor1 = append(valor1, 3)
	valor1 = append(valor1, 4)
	valor1 = append(valor1, 5)

	fmt.Println(valor1, len(valor1), cap(valor1))
}
