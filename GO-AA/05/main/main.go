package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Horas: ", horas)
		for min := 0; min < 60; min++ {
			fmt.Print(" ", min)
		}
		fmt.Println("\n")
	}
}
