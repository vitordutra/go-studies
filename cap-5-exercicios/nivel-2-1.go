// - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
// - Solução: https://play.golang.org/p/X7qm3aWSa6

package main

import "fmt"

func main() {
	x := 100
	fmt.Printf("%d, %#x, %b, %#U", x, x, x, x)
}
