/*
- Isso tudo aqui a gente já viu:
- Toda slice tem um array subjacente.
- Um slice é: um ponteiro/endereço para um array, mais len e cap (que é o len to array).
- Exemplo:
    - x := []int{...números}
    - y := append(x[:i], x[:i]...)
    - pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
    - Ou seja, y utiliza o mesmo array subjacente que x.
    - O que nos dá um resultado inesperado.
- Ou seja, bom saber de antemão pra não ter que aprender na marra.
- Go Playground: https://play.golang.org/p/BBJLuIjU_i
*/

/*
     [1 2 3 4 5]
     [1 2 5]
     [1 2 5 4 5] --> [[1 2 5] 4 5]

ATENÇÃO!!!
Perceba que houve um 'refatiamento', causando alteração ao primeiro 'slice', sendo fusionado com a modificação causada pelo segundo 'slice'.

Trata-se de uma "ineficácia da linguagem", aprende-se por prevenção, para não fazê-lo!

Como evitar:
• Operar tudo com a mesma variável ou copiar "item por item" por iteração;
• Evitar utilizar 'slices'  que se utilizam de um mesmo vetor subjacente ; alterando-o;

i.e. Manter apenas um único 'slice' declarado por variável, copiando ou recortando elementos.

*/

package main

import (
	"fmt"
)

func main() {

	primeiroslice := []int{1, 2, 3, 4, 5}

	fmt.Println(primeiroslice)

	fmt.Println(primeiroslice[:2])

	fmt.Println(primeiroslice[4:])

	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)
}
