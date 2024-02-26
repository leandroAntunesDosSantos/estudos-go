package main

import "fmt"

func main() {
	recebeArqumento(argumento)
}

func argumento() {
	fmt.Println("olha eu aqui")
}

func recebeArqumento(x func()) {
	fmt.Println("presta atençao")
	x()
}
