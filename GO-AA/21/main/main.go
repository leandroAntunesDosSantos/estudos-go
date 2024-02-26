package main

import "fmt"

func main() {
	x := 10

	y := &x         //endereco memoria
	fmt.Println(*y) //reference
	fmt.Println(*&x)
	fmt.Printf("%T,%T", x, y) //int e ponteiro de int
}
