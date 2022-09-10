// - Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
// - Solução: https://play.golang.org/p/dIbfdiuw0ms

package main

import "fmt"

func main() {
	birthYear := 1994
	currentYear := 2022

	for {
		if birthYear <= currentYear {
			fmt.Println(birthYear)
			birthYear++
		} else {
			break
		}
	}
}
