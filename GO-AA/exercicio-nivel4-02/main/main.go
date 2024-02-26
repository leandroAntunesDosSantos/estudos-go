package main

import "fmt"

func main() {

	valores := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	for _, v := range valores {
		fmt.Println(v)
	}
	fmt.Printf("%T", valores)
}
