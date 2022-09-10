/*
- Callback: passe uma função como argumento a outra função.
- Solução: https://play.golang.org/p/2epLD_Yyap
*/

package main

import "fmt"

func main() {
	x := []int{2, 4, 8, 16, 32, 64}
	y := myMap(double, x)
	fmt.Println(y)
}

func myMap(f func(x int) int, s []int) []int {
	var newSlice []int
	for _, v := range s {
		newSlice = append(newSlice, f(v))
	}
	return newSlice
}

func double(n int) int {
	return n * 2
}
