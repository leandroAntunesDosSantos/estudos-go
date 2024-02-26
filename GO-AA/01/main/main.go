package main

import "fmt"

//marmota tipagem automatica
//tipos promitivos, int, string,bol
//tipos compostos, slice, array, struct, map

var finalizar string = "ate mais"
var pode bool = true

func main() {
	x := 10.2

	y := "Ola, bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20
	y = "ok vc venceu"

	fmt.Printf("x: %v, %T,\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	valor := 1532
	qualquerCoisa(valor)

	println(finalizar)
	println(pode)

}

func qualquerCoisa(valor int) {
	fmt.Println(valor)
}
