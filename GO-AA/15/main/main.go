package main

import "fmt"

func main() {
	qualquerCoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra cadeia",
	}

	for key, value := range qualquerCoisa {
		fmt.Println(key, value)

	}
	println()
	total := 0
	contador := 0
	for key, _ := range qualquerCoisa {
		total += key
		contador++
	}
	fmt.Println(total)
	fmt.Println(contador)

	delete(qualquerCoisa, 98)
	fmt.Println(qualquerCoisa)
}
