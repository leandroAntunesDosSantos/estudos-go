package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"dasilva_guria": {"desaparecer", "ser assustadora", "raspar o cabelo"},
		"senna_ayrton":  {"andar de jetski", "ser milionario", "sotaque paulista"},
		"pike_rob":      {"criar linguagens", "usar ternos muito loucos"},
	}
	for i, v := range pessoas {
		fmt.Println(i)
		for i2, v2 := range v {
			fmt.Println("\t", i2, " - ", v2)
		}
	}
	//fmt.Println(pessoas)
}
