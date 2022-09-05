/*
- São structs sem identificadores.
- x := struct { name type }{ name: value }
- Go Playground: https://play.golang.org/p/xyhNnSCu1f
*/

package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50}

	fmt.Println(x)
}
