/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
- Solução: https://play.golang.org/p/qLY-q3vffQ
*/

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

type figura interface {
	Area() float64
}

func (q quadrado) Area() float64 {
	return q.lado * q.lado
}

func (c circulo) Area() float64 {
	return math.Pi * c.raio * c.raio
}

func info(f figura) float64 {
	return f.Area()
}

func main() {
	q := quadrado{
		lado: 15,
	}

	c := circulo{
		raio: 15,
	}

	fmt.Println(info(q))
	fmt.Println(info(c))
}
