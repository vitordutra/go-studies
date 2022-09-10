/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
- Solução: https://play.golang.org/p/GLK11Q1_x8y
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	// var meuMapa map[string]pessoa
	meuMapa := make(map[string]pessoa)

	meuMapa["Pimentão"] = pessoa{
		nome:      "Renata",
		sobrenome: "pimentão",
		sabores:   []string{"pistache", "morango", "baunilha"},
	}

	meuMapa["da Prússia"] = pessoa{
		"Frederico",
		"da Prússia",
		[]string{"sabão em pó", "pé de coelho", "feijão"},
	}

	for _, pessoa := range meuMapa {
		fmt.Println("Meu nome é", pessoa.nome, pessoa.sobrenome)
		for _, sabor := range pessoa.sabores {
			fmt.Println("\t", sabor)
		}
	}
}
