// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
// - Solução: https://play.golang.org/p/zcEsXqnBr8

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
