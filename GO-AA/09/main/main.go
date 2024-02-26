package main

import "fmt"

func main() {
	slice1 := []int{20, 21, 22, 89}
	total := 0

	for _, valor := range slice1 {
		total += valor
	}
	fmt.Println("O valor total é: ", total)
}
