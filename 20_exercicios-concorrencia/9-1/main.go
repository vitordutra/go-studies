/*
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
- Solução:
    - https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/20_exercicios-ninja-9/01_foda/main_foda.go
    - https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/20_exercicios-ninja-9/01_moleza/main_moleza.go
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novasGoroutines(100)
	wg.Wait()
}

func novasGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine numero:", i)
			wg.Done()
		}(x)
	}
}
