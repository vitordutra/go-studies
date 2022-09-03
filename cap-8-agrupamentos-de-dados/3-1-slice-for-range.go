/*
- Slices:
    - Tamanho: len(x)
    - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}
- Go Playground:
    - https://play.golang.org/p/h5-RFJn-Fh
    - https://play.golang.org/p/2wj02m3-eM
*/

package main

import (
	"fmt"
)

func main() {

	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	//for índice, valor := range slice {fmt.Println("No índice", índice, "temos o valor:", valor)}

	//slice[4] = "melancia"
	slice = append(slice, "melancia")

	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}
}
