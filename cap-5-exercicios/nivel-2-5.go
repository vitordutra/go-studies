// - Crie uma variável de tipo string utilizando uma raw string literal.
// - Demonstre-a.
// - Solução: https://play.golang.org/p/RkpqPpRWuo

package main

import "fmt"

func main() {
	x := `tinha uma pedra
				na estrada
				e a estrada
					esta para la
						de pasargada`
	fmt.Println(x)
}

