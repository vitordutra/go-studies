/*
- Agora vamos dar um mergulho na documentação:
    - https://go.dev/doc/effective_go#concurrency
    - https://pt.wikipedia.org/wiki/Multiplexador
    - O que é yield? runtime.Gosched()
- Race condition:
        Função 1       var     Função 2
         Lendo: 0   →   0
         Yield          0   →   Lendo: 0
         var++: 1               Yield
         Grava: 1   →   1       var++: 1
                        1   ←   Grava: 1
         Lendo: 1   ←   1
         Yield          1   →   Lendo: 1
         var++: 2               Yield
         Grava: 2   →   2       var++: 2
                        2   ←   Grava: 2
- E é por isso que vamos ver mutex, atomic e, por fim, channels.

- Aqui vamos replicar a race condition mencionada no artigo anterior.
    - time.Sleep(time.Second) vs. runtime.Gosched()
		runtime.Gosched() -> para de executar essa thread, vai executar outra, volta
		pra ca depois
- go help → go help build → go run -race main.go
- Como resolver? Mutex.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	variavelCompartilhada := 0
	totalDeGoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totalDeGoroutines)

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			v := variavelCompartilhada
			runtime.Gosched()
			v++
			variavelCompartilhada = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor Final:", variavelCompartilhada)
}
