/*
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
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

	fmt.Println("---------------------------------------")
	
	delete(nomeEHobbies, "vitor_dutra")

	for key, values := range nomeEHobbies {
		fmt.Println(key)
		for _, hobby := range values {
			fmt.Printf("\t%v\n", hobby)
		}
	}
}
