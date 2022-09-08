/*
- Pode-se usar uma função como retorno de uma função
- Declaração: func f() return
- Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
    - ????: fmt.Println(f()())
- Go Playground: https://play.golang.org/p/zPjoWNrCJF
*/

package main

import "fmt"

func main() {
	x := returnsAFunction()
	y := x(3)
	fmt.Println(y)
	
	fmt.Println(returnsAFunction()(3))
}

func returnsAFunction() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
