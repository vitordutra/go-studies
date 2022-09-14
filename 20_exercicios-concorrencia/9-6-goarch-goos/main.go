/*
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build
    - go install
- Solução: https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/20_exercicios-ninja-9/06/main.go
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
}
