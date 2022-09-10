/*
- Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
- Solução: https://play.golang.org/p/3wgW4BDasu
*/
package main

import (
	"fmt"
	"sort"
)

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type byLastName []user

func (a byLastName) Len() int           { return len(a) }
func (a byLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a byLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type bySayings []user

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(byAge(users))
	fmt.Println("Sort By Age\n", users)

	sort.Sort(byLastName(users))
	fmt.Println("Sort By Last Name\n", users)
}
