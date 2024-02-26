package main

import "fmt"

type batata int

var x batata
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%v\n %T\n", y, y)

}
