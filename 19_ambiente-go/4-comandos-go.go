/*
- go version
- go env
- go help
- go fmt
    - ./…
- go run
    - go run [file name]
    - go run *.go
- go build
    - para um executável:
        - gera o arquivo binário
        - informa caso tenham havido erros
        - caso não hajam erros, cria um executável e salva-o no diretório atual
    - para um pacote:
        - gera o arquivo
        - informa caso tenham havido erros
        - descarta o executável
- go install
    - para um executável:
        - faz o build
        - nomeia o arquivo com o nome do diretório atual
        - salva o arquivo binário em $GOPATH/bin
    - para um pacote:
        - faz o build
        - salva o arquivo binário em $GOPATH/pkg
        - cria archive files (arquivo.a), os arquivos pré-compilados utilizados pelos imports
- flags
    - "-race"
*/

package main

func main() {

}
