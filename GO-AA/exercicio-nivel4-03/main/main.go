package main

import "fmt"

func main() {
	valores := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 10}
	fmt.Println(valores[0:3])
	fmt.Println(valores[4:])
	fmt.Println(valores[1:7])
	fmt.Println(valores[2 : len(valores)-1])
}
