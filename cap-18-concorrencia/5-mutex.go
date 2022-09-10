/*
- Agora vamos resolver a race condition do programa anterior utilizando mutex.
- Mutex é mutual exclusion, exclusão mútua.
- Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
- Na prática:
    - type Mutex
        - func (m *Mutex) Lock()
        - func (m *Mutex) Unlock()
- RWMutex
- Código: https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/18_concorrencia/06_mutex/main.go
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
	var mu sync.Mutex

	wg.Add(totalDeGoroutines)

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			mu.Lock()

			v := variavelCompartilhada
			runtime.Gosched()
			v++
			variavelCompartilhada = v

			mu.Unlock()

			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor Final:", variavelCompartilhada)
}
