package main

import "fmt"

var numero1 int = 15

func main() {
	fmt.Printf("%d %b %#x", numero1, numero1, numero1)
	fmt.Println()
	numero2 := numero1 << 1
	fmt.Printf("%d %b %#x", numero2, numero2, numero2)

}
