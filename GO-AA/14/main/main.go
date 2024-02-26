package main

import "fmt"

func main() {
	amigos := map[string]int{
		"alfredo": 55512345,
		"joana":   999999,
	}
	fmt.Println(amigos)
	fmt.Println()
	fmt.Println(amigos["joana"])
	amigos["gopher"] = 44444

	fmt.Println(amigos)
	fmt.Println()

	if sera, ok := amigos["joana"]; ok { //verifica se possui o nome ao invés de retornar 0;
		fmt.Println(sera)
	} else {
		fmt.Println("nao tem")
	}
}
