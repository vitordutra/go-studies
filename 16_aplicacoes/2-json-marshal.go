/*
- Exemplo: transformando structs em Go em código JSON.
- No improviso tambem.
- Go Playground: https://play.golang.org/p/_JvCOlK-H9
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {

	jamesBond := Pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		ContaBancaria: 1000000.50,
	}

	darthVader := Pessoa{
		Nome:          "Anakin",
		Sobrenome:     "Skywalker",
		Idade:         50,
		Profissao:     "Intergalactic Villain",
		ContaBancaria: 1345123452345.60,
	}

	j, err := json.Marshal(jamesBond)

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(darthVader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))
}
