package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oiBomDia() {
	fmt.Println(p.nome, "diz bom dia")
}

func main() {
	mauricio := pessoa{"mauricio", 30}
	mauricio.oiBomDia()
}
