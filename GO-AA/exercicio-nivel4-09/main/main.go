package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"dasilva_guria": {"desaparecer", "ser assustadora", "raspar o cabelo"},
		"senna_ayrton":  {"andar de jetski", "ser milionario", "sotaque paulista"},
		"pike_rob":      {"criar linguagens", "usar ternos muito loucos"},
	}
	//adicionar uma entrada ao map pessoas
	pessoas["loureiro_kiko"] = []string{"usar os trequinho no punho", "tacar fogo na guitarra"}

	delete(pessoas, "senna_ayrton")

	for i, v := range pessoas {
		fmt.Println(i)
		for i2, v2 := range v {
			fmt.Println("\t", i2, " - ", v2)
		}
	}
	//fmt.Println(pessoas)
}
