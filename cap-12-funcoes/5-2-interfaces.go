/*
- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo e também o tipo da interface.

- Exemplos:
    - Os tipos profissão1 e profissão2 contem o tipo pessoa
    - Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
    - Implementam a interface gente
    - Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
    - Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }*
    - Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
    https://play.golang.org/p/zGKr7cvTPF
    - Go Playground ao vivo:
    https://play.golang.org/p/njiKbTT20Cr
- Onde se utiliza?
    - Área de formas geométricas (gobyexample.com)
    - Sort
    - DB
    - Writer interface: arquivos locais, http request/response
- Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type arquiteto struct {
	pessoa
	especialização  string
	obrasConcluídas int
}

type dentista struct {
	pessoa
	especialização   string
	dentesArrancados int
}

type gente interface {
	OiBomDia()
}

func (d dentista) OiBomDia() {
	fmt.Printf("Oi, bom dia! Meu nome é %v e eu já arranquei %d dentes!\n", d.nome, d.dentesArrancados)
}

func (a arquiteto) OiBomDia() {
	fmt.Printf("Oi, bom dia! Meu nome é %v e eu já concluí %d obras!\n", a.nome, a.obrasConcluídas)
}

func demonstrarPolimorfismo(g gente) {
	fmt.Println("--- func demonstrarPolimorfismo ---")

	switch g.(type) {
	case arquiteto:
		fmt.Println(g.(arquiteto).nome, "é uma pessoa que trabalha com", g.(arquiteto).especialização)
	case dentista:
		fmt.Println(g.(dentista).nome, "é uma pessoa que trabalha com", g.(dentista).especialização)
	}

	fmt.Println("Ele diz bom dia!")
	g.OiBomDia()
}

func main() {
	alfredo := arquiteto{
		pessoa: pessoa{
			nome:      "Alfredo",
			sobrenome: "da Silva",
			idade:     30,
		},
		especialização:  "Galpões",
		obrasConcluídas: 10,
	}

	rogério := dentista{
		pessoa: pessoa{
			nome:      "Alfredo",
			sobrenome: "da Silva",
			idade:     30,
		},
		especialização:   "Galpões",
		dentesArrancados: 1124,
	}

	fmt.Println(alfredo)
	alfredo.OiBomDia()
	demonstrarPolimorfismo(alfredo)
	fmt.Println("-----")
	fmt.Println(rogério)
	rogério.OiBomDia()
	demonstrarPolimorfismo(rogério)
}
