/*
- Qual a utilidade de funções?
    - Abstrair funcionalidade
    - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
    - Funções são definidas com parâmetros
    - Funções são chamadas com argumentos
- Tudo em Go é pass by value.
    - Pass by reference, pass by copy, ... não.
- Parâmetro pode ser ...variádico.
- Exemplos:
    - Função básica.
        - Go Playground: https://play.golang.org/p/FebJblBenP
    - Função que aceita um argumento.
        - Go Playground:
        https://play.golang.org/p/CE6Ij3U4QB
    - Função com retorno.
        - Go Playground: https://play.golang.org/p/gKxwYe6btP
    - Função com múltiplos retornos e parâmetro variádico.
        - Go Playground: https://play.golang.org/p/OcQ1wXwM2c
    - Mais um: https://play.golang.org/p/8wc2TA9xH_
*/

package main

import "fmt"

func main() {
	soma, tamanho, oi := soma("tarde", 12, 13, 14, 15)
	fmt.Println(soma, tamanho, oi)
}

func soma(s string, x ...int) (int, int, string) {
	oi := ""
	if s == "manhã" {
		oi = "Oi, bom dia!"
	} else if s == "tarde" {
		oi = "Oi, boa tarde"
	} else {
		oi = "Oi, boa noite"
	}

	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma, len(x), oi
}
