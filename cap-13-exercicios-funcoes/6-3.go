/*
- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
- Solução: https://play.golang.org/p/b5Ua2kNN8a
*/

package main

import "fmt"

func main() {
	defer fmt.Println("This will appear later")
	fmt.Println("This will appear first")
}
