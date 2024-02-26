package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "Volei":
		fmt.Println("Vamo de volei")
	case "lol":
		fmt.Println("Vamo joga lol")
	case "dota":
		fmt.Println("Vamo joga dota")
	case "futebol":
		fmt.Println("Vamo joga futebol")
	}
}
