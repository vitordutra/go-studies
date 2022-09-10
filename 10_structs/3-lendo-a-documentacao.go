/*
- É importante se familiarizar com a documentação da linguagem Go.
- Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
- Veremos:
    - ref/spec
        - Já vimos mais da metade dos tipos em Go!
        - Struct types.
            - x, y int
            - anonymous fields
            - promoted fields
- Go Playground: https://play.golang.org/p/z9UQej4IQT
*/

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}

	pessoa3 := pessoa{"Mauricio", 40}
	pessoa4 := profissional{pessoa{"Vanderlei", 50}, "Político", 10000000}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.pessoa.nome)
	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa4)
}
