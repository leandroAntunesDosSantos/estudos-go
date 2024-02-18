package main

import "fmt"

func main() {
	numeroDeBytes, erros := fmt.Println("Hello, World", "tudo certo com o codigo", 100)
	fmt.Println(numeroDeBytes, erros)

	x := 16
	y := "strings"
	z := true

	fmt.Println(x, y, z)
}
