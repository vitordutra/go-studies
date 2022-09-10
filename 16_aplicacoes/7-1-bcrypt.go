/*
- É uma maneira de encriptar senhas utilizando hashes.
- x/crypto/bcrypt
    - GenerateFromPassword
    - CompareHashAndPassword
- Sem Go Playground!
    - go get golang.org/x/crypto/bcrypt
- Arquivo: https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/16_aplicacao/bcrypt/main.go
*/

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "20julho1980"
	// senhaErrada := "20julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

}
