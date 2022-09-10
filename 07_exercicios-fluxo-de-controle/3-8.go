// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
// - Solução: https://play.golang.org/p/TyIGp4Hi8B

package main

import "fmt"

func main() {
	cursoParaEstudar := 3

	switch {
	case cursoParaEstudar == 0:
		fmt.Println("Alura Java")
	case cursoParaEstudar == 1:
		fmt.Println("DevSuperior")
	case cursoParaEstudar == 2:
		fmt.Println("Curso de Go da Ellen Korbes")
	case cursoParaEstudar == 3:
		fmt.Println("Docker")
	}
}
