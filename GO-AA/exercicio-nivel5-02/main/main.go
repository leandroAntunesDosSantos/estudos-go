package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	//var meuMapa map[string]pessoa
	meuMapa := make(map[string]pessoa)
	meuMapa["mendes"] = pessoa{
		"jailson",
		"mendes",
		[]string{"morango", "chocolate"},
	}

	meumapa2 := map[string]pessoa{
		"mendes": {
			"jailson",
			"mendes",
			[]string{"morango", "chocolate"},
		},
		"pereira": {
			"marcos",
			"pereira",
			[]string{"baunilha", "milho", "batata"},
		},
	}

	meuMapa["pereira"] = pessoa{
		"marcos",
		"pereira",
		[]string{"baunilha", "milho", "batata"},
	}
	for _, valor := range meuMapa {
		fmt.Println(valor)
	}
	fmt.Println()
	for _, valor := range meumapa2 {
		fmt.Println("meu nome é", valor.nome, " e meus sorvetes favoritos sao:")
		for _, valor := range valor.sabores {
			fmt.Println("-", valor)
		}
		fmt.Println("\n")
	}
}
