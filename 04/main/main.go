package main

import "fmt"

const x = 10 //tipo nao é definido na criaçao

var y = 10 // tipo int definido mesmo que nao tenha

var c int
var d float64

func main() {
	k := x + 10.5
	fmt.Printf("%v %T\n", k, k)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
}
