package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "joao",
		sobrenome: "da silva",
		fumante:   false,
	}
	cliente2 := cliente{"joana", "pereira", true}
	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
