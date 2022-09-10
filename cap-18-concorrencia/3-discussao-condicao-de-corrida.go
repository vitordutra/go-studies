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
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	wg.Done()
}
