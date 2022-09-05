/*
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Solução: https://play.golang.org/p/nD3TW8VQmH
*/

package main

import "fmt"

func main() {
    nomeEHobbies := map[string][]string {
        "jessica_veloso": {"medicina", "grey's anatomy"},
        "vitor_dutra": {"programação", "filmes", "series"},
        "aline_gois": {"psicanalise", "espiritismo"},
    }

    fmt.Println(nomeEHobbies)

    for key, values := range nomeEHobbies {
        fmt.Println(key)
        for _, hobby := range values {
            fmt.Printf("\t%v\n", hobby)
        }
    }
}
