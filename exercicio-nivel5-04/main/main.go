package main

import "fmt"

func main() {
	x := struct {
		nome      string
		sabor     string
		ondeTem   []string
		vaiBemCom map[string]string
	}{
		nome:    "macarronada",
		sabor:   "macarrao",
		ondeTem: []string{"italia", "brasil"},
		vaiBemCom: map[string]string{
			"no cafe da manha": "chá",
			"no almoço":        "almondegas",
			"na janta":         "aperitivo",
		},
	}
	fmt.Println(x)
}
