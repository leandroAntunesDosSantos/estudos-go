package main

import "fmt"

func main() {
	informacoes := [][]string{
		{"miu", "milton", "encher o saco"},
		{"mimi", "martha", "pedir comida"},
		{"meus alunos", "que estudam", "fazer tarefas"},
	}

	for _, v := range informacoes {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}
}
