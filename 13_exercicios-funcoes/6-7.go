/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
- Solução: https://play.golang.org/p/A74rufv6Rs
*/

package main

import "fmt"

func main() {
	returnsAFunction()()
}

func returnsAFunction() func() {
	return func() {
		fmt.Println("This function was returned by returnsAFunction")
	}
}
