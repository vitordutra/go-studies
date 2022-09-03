/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
- Solução: https://play.golang.org/p/4-iTPZwfEz
*/

package main

import "fmt"

func main() {
	esporteFavorito := "Futebol"

	switch esporteFavorito {
	case "Xadrez":
		fmt.Println("Esporte Favorito eh Xadrez")
	case "Futebol":
		fmt.Println("Esporte favorito eh Futebol")
	case "Natacao":
		fmt.Println("Esporte Favorito eh Natacao")
	}
}
