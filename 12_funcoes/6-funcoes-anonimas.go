/*
- Anonymous self-executing functions → Funções anônimas auto-executáveis.
- func(p params) { ... }()
- Vamos ver bastante quando falarmos de goroutines.
- Go Playground: https://play.golang.org/p/Rnqmo6X6jh
*/

package main

import "fmt"

func main() {
	x := 387

	func(x int) {
		fmt.Println(x, "times 873648 is:")
		fmt.Println(x * 873648)
	}(x)
}
