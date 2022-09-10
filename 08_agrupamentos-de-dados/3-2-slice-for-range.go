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

	slice := []int{20, 21, 22, 23, 1, 13}

	total := 0

	for _, valor := range slice {

		// mesma coisa que total = total + valor
		total += valor

	}

	fmt.Println("O valor total é:", total)
}
