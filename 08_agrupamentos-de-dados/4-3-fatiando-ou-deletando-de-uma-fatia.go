/*
- x[:], x[a:], x[:b], x[a:b]
- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
    - Off-by-one error.
- Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
- É fatiando que se deleta um item de uma slice. Na prática:
    - x := append(x[:i], x[:i]...)
    - Go Playground: https://play.golang.org/p/xK2HwCqvwd
- Exercício: tente acessar todos os itens de uma slice sem utilizar range.
- Solução: https://play.golang.org/p/aUC9qVCobH
*/

package main

import (
	"fmt"
)

func main() {

	//			    0.           1.           2.         3.               4.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	fatia := sabores[:]

	fmt.Println(fatia)

	fmt.Println(sabores[0], sabores[1], sabores[2], "\n")

	for i := 0; i < len(sabores); i++ {

		fmt.Println(sabores[i])
	}

}
