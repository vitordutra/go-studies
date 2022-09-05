/*
- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
- Solução: https://play.golang.org/p/ST3TKusuOd
*/

package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 90, 100}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)
}
