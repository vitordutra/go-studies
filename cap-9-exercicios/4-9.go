/*
- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
- Solução: https://play.golang.org/p/3fcvHlt8Lm
*/

package main

import "fmt"

func main() {
	nomeEHobbies := map[string][]string{
		"jessica_veloso": {"medicina", "grey's anatomy"},
		"vitor_dutra":    {"programação", "filmes", "series"},
		"aline_gois":     {"psicanalise", "espiritismo"},
	}

	nomeEHobbies["zadock_jr"] = []string{"carro", "futebol", "rambo"}

	for key, values := range nomeEHobbies {
		fmt.Println(key)
		for _, hobby := range values {
			fmt.Printf("\t%v\n", hobby)
		}
	}
}
