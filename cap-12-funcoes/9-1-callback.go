/*
- Primeiro veja se você entende isso: https://play.golang.org/p/QkAtwMZU-z
- Callback é passar uma função como argumento.
- Exemplo:
    - Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.
    - Go Playground:
- Desafio: Crie uma função no programa acima que utilize somente os números ímpares.
*/

package main

import "fmt"

func main() {
	fib := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	sumOfEvens := onlyEvenNumbers(sum, fib...)
	fmt.Println(sumOfEvens)

	sumOfOdds := onlyOddNumbers(sum, fib...)
	fmt.Println(sumOfOdds)

	fmt.Println(sum(fib...))
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func onlyEvenNumbers(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}

	total := f(slice...)

	return total
}

func onlyOddNumbers(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
