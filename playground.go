package main

import "fmt"

func main() {
	if x := 100; x > 100 {
		fmt.Println("x > 100")
	} else {
		fmt.Println("x <= 100")
	}
}
