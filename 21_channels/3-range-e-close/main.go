/*
- Range:
    - gofunc com for loop com send e close(chan)
    - recebe com range chan

- Código: https://play.golang.org/p/_g5IEjSkh1
*/

package main

import "fmt"

// Não roda!
func main() {
	c := make(chan int)
    go myLoop(10, c)
    prints(c)
}

func myLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
}
