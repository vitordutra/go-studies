/*
- Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
- Exemplos:
    - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
        - Go Playground: https://play.golang.org/p/k8O3_8UDa
    - Pode-se passar zero ou mais valores
        - Go Playground: https://play.golang.org/p/C238I9n7Vs
    - O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
        - Go Playground: https://play.golang.org/p/8wc2TA9xH
        - Não roda: https://play.golang.org/p/2qTAnLWfgB
*/

package main

import "fmt"

func main() {
	si := []int{10, 10, 1, 2, 3, 5}
	total := soma(si...)
	fmt.Println(total)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
