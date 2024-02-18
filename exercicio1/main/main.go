package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	println(x, y, z)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", z)
	fmt.Printf("x: %v, y: %v, z: %v", x, y, z)
}
