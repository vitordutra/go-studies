/*
- O que são tipos de dados compostos?
    - Wikipedia: Composite_data_type
    - Effective Go: Composite literals
    - ref/spec: Composite literals
- Uma slice agrupa valores de um único tipo.
- Criando uma slice: literal composta → x := []type{values}
- Go Playground: https://play.golang.org/p/W7Cxm8NPZC
*/

package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
	slice[20] = 1
	fmt.Println(slice[20])
}
