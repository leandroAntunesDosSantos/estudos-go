package main

import "fmt"

func main() {
	sliceOfInt := []int{1, 2, 3, 4, 5}
	fmt.Println(funcao1(sliceOfInt...))
	fmt.Println(funcao2(sliceOfInt))
}

func funcao1(x ...int) int {
	total := 0
	for _, valor := range x {
		total += valor
	}
	return total
}

func funcao2(x []int) int {
	total := 0
	for _, valor := range x {
		total += valor
	}
	return total
}
