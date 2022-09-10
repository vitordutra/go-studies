/*
- Range: for k, v := range map { }
- Reiterando: maps não tem ordem e um range usará uma ordem aleatória.
- Go Playground: https://play.golang.org/p/6zEMfIP-AE
- delete(map, key)
- Deletar uma key não-existente não retorna erros!
- Go Playground: https://play.golang.org/p/0uuIicU3Zz
*/

package main

import (
	"fmt"
)

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)
}
