/*
- Utilize atomic para consertar a condição de corrida do exercício #3.
- Solução: https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/20_exercicios-ninja-9/05/main.go
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

const quantidadeDeGoroutines = 50000

var contador int32

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
			// v := contador
			atomic.AddInt32(&contador, 1)
			runtime.Gosched()
			// v++
			// contador = v
			atomic.LoadInt32(&contador)
			fmt.Println(contador)
			wg.Done()
		}()
	}
}
