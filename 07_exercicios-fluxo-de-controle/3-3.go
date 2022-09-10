// - Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
// - Solução: https://play.golang.org/p/qnFjiDJzLor

package main

import "fmt"

func main() {
	birthYear := 1994
	currentYear := 2022

	for birthYear <= currentYear {
		fmt.Println(birthYear)
		birthYear++
	}
}
