/*
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
- Solução: https://play.golang.org/p/1aPXVeR1mf
*/

package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", slice)

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:8])
	fmt.Println(slice[2:len(slice) - 1])

}
