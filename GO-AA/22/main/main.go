package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {
	jamesBond := Pessoa{
		"james",
		"bond",
		40,
		"Agente secreto",
		10000.50,
	}
	darthVader := Pessoa{
		"anakin",
		"skylwalker",
		58,
		"vilao intergalactico",
		5800000,
	}
	j, err := json.Marshal(jamesBond)
	if err != nil {
		fmt.Println(err)
	}
	d, err := json.Marshal(darthVader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))
	fmt.Println(string(d))
}

//para importar é necessario estar em letras maiusculas!!
