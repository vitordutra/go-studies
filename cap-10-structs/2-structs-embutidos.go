/*
- Struct é um tipo de dados composto que nos permite armazenar valores de tipos diferentes.
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
- Go Playground: https://play.golang.org/p/5i0DqxuBp1
*/

package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}

	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
