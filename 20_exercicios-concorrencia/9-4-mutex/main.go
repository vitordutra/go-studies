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
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

const quantidadeDeGoroutines = 100

var contador int

func main() {
	contador = 0
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de Goroutines:", quantidadeDeGoroutines, "\nTotal do contador", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
