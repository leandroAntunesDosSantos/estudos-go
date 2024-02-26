package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{
		"jailson",
		"mendes",
		[]string{"morango", "chocolate"},
	}
	pessoa2 := pessoa{
		"marcos",
		"pereira",
		[]string{"baunilha", "milho", "batata"},
	}
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for i, v := range pessoa1.sabores {
		fmt.Println("\t", i, v)
	}
	fmt.Println()
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for i, v := range pessoa2.sabores {
		fmt.Println("\t", i, v)
	}
}
