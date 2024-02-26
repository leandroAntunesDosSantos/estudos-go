package main

import "fmt"

func main() {
	x := retornaUmaFuncao()
	x()
}

func retornaUmaFuncao() func() {
	return func() {
		fmt.Println("Olha eu aqui")
	}
}
