/*
- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
- Se precisar de dicas, veja: https://gobyexample.com/interfaces
- Solução: https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/20_exercicios-ninja-9/02/main.go
*/

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type humano interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz oi!")
}

func dizerAlgumaCoisa(h humano) {
	h.falar()
}

func main() {
	p1 := pessoa{"Genghis Khan", 1000}
	p1.falar()
	dizerAlgumaCoisa(&p1)
	// dizerAlgumaCoisa(&p1)
}
