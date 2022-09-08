/*
- Todos os valores ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um pointeiro se refere a esse endereço.
- Notações:
    - &variável mostra o endereço de uma variável
        - %T: variável vs. &variável
    - *variável faz de-reference, mostra o valor que consta nesse endereço
    - ????: *&var funciona!
    - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
- Exemplo: a := 0; b := &a; *b++
- Go Playground: https://play.golang.org/p/gC1qGFUYrV
*/

package main

import "fmt"

func main() {
	x := 0
	y := &x

	fmt.Println(x, y)
	*y++
	fmt.Println(*y)
	fmt.Printf("%T, %T\n", x, y)
	fmt.Println(x, y)
}
