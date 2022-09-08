/*
- Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
    - Não queremos passar grandes volumes de dados pra lá e pra cá
    - Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)
- Exemplos:
    - x := 0; funçãoquemudaovalordoargumentopra1(x); Print(x)
    - x := 0; funçãoquemudaovalordo*argumentopra1(&x); Print(x)
- Go Playground: https://play.golang.org/p/VZmfWfw76s
*/

package main

import "fmt"

func main() {
	x := 10

	// 10 outside the Function
	// 11 inside the function
	// passByValue(x)

	passByReference(&x)

	fmt.Println("Outside the function", x)
}

func passByValue(x int) {
	x++
	fmt.Println("Inside the function:", x)
}

func passByReference(x *int) {
	*x++
	fmt.Println("Inside the function:", *x)
}
