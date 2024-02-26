package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	zezinho := pessoa{"zezinho", "primo", 10}
	fmt.Println(zezinho)
	mudarPessoa(&zezinho)
	fmt.Println(zezinho)
}

func mudarPessoa(p *pessoa) {
	(*p).nome = "Luizinho"
	(*p).sobrenome = "irmao do luizinho"
}
