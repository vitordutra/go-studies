/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
- Solução: https://play.golang.org/p/Tgv3wwxKV-
*/

package main

import "fmt"

func main() {
	slice := []int{1967, 1968, 1994}
	variadic := variadicSum(slice...)
	fmt.Println(variadic)

    // invalid operation: cannot call non-function variadic (variable of type int)
    // fmt.Println(variadic(slice))

    fmt.Println(sliceSum(slice))
}

func variadicSum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sliceSum(x []int) int {
    sum := 0
    for _, v := range x {
        sum += v
    }
    return sum
}
