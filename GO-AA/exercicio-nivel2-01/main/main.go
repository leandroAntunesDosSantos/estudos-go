package main

import "fmt"

//decimal binário hexadecimal

var x = 100

func main() {
	fmt.Printf("%d %b %#x", x, x, x) //o # adiciona 0x antes do valor hexadecimal
}
