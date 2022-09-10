/*
- GOOS
- GOARCH
- Compilando para Linux:
- `GOOS=linux GOARCH=amd64 go build test.go`
- Compilando para Windows:
- `GOOS=windows GOARCH=amd64 go build test.go`
- Compilando para Mac:
- `GOOS=darwin GOARCH=amd64 go build test.go`
- https://godoc.org/runtime#pkg-constants
- git push
- git clone
- go get
- Arquivos: https://github.com/vkorbes/aprendago/tree/master/c%C3%B3digo/19_seu-ambiente-de-desenvolvimento/compilacaocruzada

*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Esse é programa do exercício de compilação cruzada. Foi compilado num linux/amd64, e agora está rodando num sistema:", runtime.GOARCH, runtime.GOOS)

	time.Sleep(time.Second)
}
