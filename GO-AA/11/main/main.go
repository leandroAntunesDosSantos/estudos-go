package main

import "fmt"

func main() {
	sabores := []string{"pepperoni", "mozzarella", "abacaxi", "quatroqueijos", "marquerita"}

	sabores = append(sabores[0:2], sabores[3:]...)

	fmt.Println(sabores)
	fmt.Println()

	valores1 := []int{1, 2, 3, 4}
	valores2 := []int{9, 10, 11, 12}
	valores3 := append(valores1, valores2...)
	fmt.Println(valores3)

}
