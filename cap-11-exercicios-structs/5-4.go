/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
- Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG
*/

package main

import "fmt"

func main() {
	x := struct {
		nome      string
		sabor     string
		ondeTem   []string
		vaiBemCom map[string]string
	}{
		nome:    "Stroopwaffel",
		sabor:   "Doce",
		ondeTem: []string{"Holanda", "Casa do tio rico"},
		vaiBemCom: map[string]string{
			"Café da manhã": "Chá",
			"Almoço":        "Café",
			"Jantar":        "Chocolate Quente",
		},
	}

	fmt.Println(x)
}
