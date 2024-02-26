package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{"leandro", 32}
	u2 := user{"mariana", 27}
	u3 := user{"ronaldo", 33}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(sb))
}
