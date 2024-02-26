package main

import "fmt"

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))
	fmt.Println()
	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array1)
	slice1 := []int{1, 2, 3, 4, 5, 6} //slice nao tem limitador
	fmt.Println(slice1)
	fmt.Println()
	slice2 := append(slice1, 7, 8, 9)
	slice2[2] = 34
	fmt.Println(slice2)

}
