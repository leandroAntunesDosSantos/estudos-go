package main

import (
	"encoding/json"
	"fmt"
)

type jsonToGo []struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"leandro","Age":32},{"First":"mariana","Age":27},{"First":"ronaldo","Age":33}] `
	fmt.Println(s)

	var resultado jsonToGo

	err := json.Unmarshal([]byte(s), &resultado)
	if err != nil {
		fmt.Println("Deu ruim!", err)
	}
	fmt.Println(resultado)
}
