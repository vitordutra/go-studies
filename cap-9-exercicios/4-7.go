/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
- Solução: https://play.golang.org/p/Gh81-d5tMi
*/

package main

import "fmt"

func main() {
	ss := [][]string{
		{"Jessica", "Veloso", "Medicina"},
		{"Vitor", "Dutra", "Programação"},
		{"Aline", "Gois", "Psicanalise"},
	}

	for _, v := range ss {
		for _, item := range v {
			fmt.Println(item)
		}
	}
}
