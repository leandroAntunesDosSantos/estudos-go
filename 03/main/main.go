package main

import "fmt"

func main() {
	s := "Hello, playground"

	sb := []byte(s)

	fmt.Printf("%v\n%T\n", sb, sb)

	for _, v := range sb {
		fmt.Printf("%v- %T - %#U - %#x\n", v, v, v, v)
	}
	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v- %T - %#U - %#x\n", s[i], s[i], s[i], s[i])

	}
}
