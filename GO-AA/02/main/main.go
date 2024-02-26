package main

import (
	"fmt"
)

var x int //declaraçao

type hotdog int

var y hotdog = 10

func main() {
	x = 10 //inicializacao atribui valor
	fmt.Printf("x: %v, %T\n", x, x)

	x = 20 //atribuicao
	fmt.Printf("x: %v, %T\n", x, x)

	fmt.Printf("y: %v, %T\n", y, y)

	//z := x + y //tipos diferentes
	//fmt.Printf("z: %v, %T\n", z, z)

}
