/*
- Slices multi-dimensionais são slices que contem slices.
- São como planilhas.
- [][]type
- Go Playground: https://play.golang.org/p/vKyHiG1GtM
- Só pra sacanear: https://play.golang.org/p/ZSU_8eJ9Yp
*/

package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		// Índice: 0  1  2                   // Índice:
		{1, 2, 3, 4, 5, 6},       // 0
		{7, 8, 9, 10, 11, 12},    // 1
		{13, 14, 15, 16, 17, 18}, // 2
	}
	fmt.Println(ss[2][4])
}
