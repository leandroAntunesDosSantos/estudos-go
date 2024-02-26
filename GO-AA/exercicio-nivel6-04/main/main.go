package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) metodo() {
	fmt.Println("ola", p.nome, p.sobrenome, "voce tem", p.idade, "anos")
}

func main() {
	pessoa1 := pessoa{"alessandra", "rodrigues", 32}
	pessoa1.metodo()
}
