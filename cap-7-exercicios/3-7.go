// - Crie um programa que demonstre o funcionamento da declaração if.
// - Solução: https://play.golang.org/p/OGPgTJZbiy

package main

import "fmt"

func main() {
	precisoTerminarGo := false
	precisaEstudarJava := false
	if precisoTerminarGo {
		fmt.Println("Voce precisa terminar o curso de Go, mas vai com calma")
	} else if precisaEstudarJava {
		fmt.Println("Voce vai terminar o curso do DevSuperior, fica de boa")
	} else {
		fmt.Println("Provavelmente voce terminou os cursos, descansa e vai para a proxima jornada")
	}
}
