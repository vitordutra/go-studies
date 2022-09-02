// - Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
// - Demonstre estes valores.
// - Solução: https://play.golang.org/p/BMYEch6_s8

package main

import "fmt"

func main() {
	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 10)
	d := (10 < 10)
	e := (10 >= 10)
	f := (10 > 10)

	fmt.Println(a, b, c, d, e, f)
}
