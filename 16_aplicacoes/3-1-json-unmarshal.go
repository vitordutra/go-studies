/*
- E agora o contrário.
- Link: https://cdn.rawgit.com/GoesToEleven/g...
- JSON-to-Go
- Tags
- Marshal/unmarshal vs. encoder/decoder
    - Marshal vai pra uma variável
    - Encoder "vai direto"
- Go Playground: https://play.golang.org/p/l6wbuLu1NS
- Com Encoder: https://play.golang.org/p/Pgwr0O07aL
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	ContaBancaria float64 `json:"ContaBancaria"`
}

func main() {

	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","ContaBancaria":1000000.5}`)
	
	var jamesBond Pessoa

	err := json.Unmarshal(sb, &jamesBond)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(jamesBond)
	fmt.Println(jamesBond.Profissao)
}
