package main

import "fmt"

func retornaInt() int {
	return 10
}
func retornaIntEString() (int, string) {
	return 20, "vinte"
}
func main() {
	fmt.Println(retornaInt())
	fmt.Println(retornaIntEString())
}
