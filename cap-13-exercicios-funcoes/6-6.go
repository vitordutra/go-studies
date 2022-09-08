/*
- Crie e utilize uma função anônima.
- Solução: https://play.golang.org/p/Kgo6hVr5G5
*/

package main

import "fmt"

func main() {
	x := 2
	y := func (n int) int{
		return n * 2
	} (x)
	fmt.Println(y)
}
