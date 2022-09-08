/*
- f := func(p params){ ... }
- f()
- Go Playground: https://play.golang.org/p/cPxhPUbfLy
*/

package main

import "fmt"

func main() {
	x := 12
	y := func(x int) int {
		return x * 873648
	}

	fmt.Println(x, "times 873648 is", y(x))
}
