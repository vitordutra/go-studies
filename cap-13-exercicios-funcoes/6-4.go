/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
- Solução: https://play.golang.org/p/GBZcnu0Ajp
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) MostrarNomeCompletoEIdade() {
	fmt.Println(p.nome, p.sobrenome, "idade:", p.idade, "anos")
}

func main() {
	pessoinha := pessoa{
		nome: "Joao Vitor",
		sobrenome: "Dutra",
		idade: 28,
	}

	pessoinha.MostrarNomeCompletoEIdade()
}
