/*
- O sort que eu quero não existe. Quero fazer o meu.
- Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
    - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
- Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
- E aí posso fazer do jeito que eu quiser.
- Exemplo:
    - struct carros: nome, consumo, potencia
    - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
    - tipo ordenarPorPotencia
    - tipo ordenarPorConsumo
    - tipo ordenarPorLucroProDonoDoPosto
- Go Playground: https://play.golang.org/p/KOIhAsE3OK
*/

package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

// Ordernar Por Potencia
type ordenarPorPotencia []carro

func (c ordenarPorPotencia) Len() int           { return len(c) }
func (c ordenarPorPotencia) Less(i, j int) bool { return c[i].potencia < c[j].potencia }
func (c ordenarPorPotencia) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// Ordernar por consumo
type ordernarPorConsumo []carro

func (c ordernarPorConsumo) Len() int           { return len(c) }
func (c ordernarPorConsumo) Less(i, j int) bool { return c[i].consumo > c[i].consumo }
func (c ordernarPorConsumo) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func main() {
	carros := []carro{
		{"Chevete", 50, 8},
		{"Porsche", 300, 5},
		{"Fusca", 20, 30},
	}

	fmt.Println("Inicial\n", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Potencia\n", carros)

	sort.Sort(ordernarPorConsumo(carros))

	fmt.Println("Consumo\n", carros)

}
