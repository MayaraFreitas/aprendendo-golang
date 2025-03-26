package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenarporidade []user
type ordenarporsobrenome []user

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
	fmt.Println("sem ordenacao:\n", users)

	sort.Sort(ordenarporidade(users))
	fmt.Println("\nordenado por idade:")
	for _, v := range users {
		fmt.Println(v.Age, "-", v.First, v.Last)

	}

	sort.Sort(ordenarporsobrenome(users))
	fmt.Println("\nordenado por sobrenome")
	for _, v := range users {
		fmt.Println(v.First, v.Last, "-", v.Age)

	}
}

// Ordenar por idade
func (users ordenarporidade) Len() int {
	return len(users)
}

func (users ordenarporidade) Less(i, j int) bool {
	return users[i].Age < users[j].Age
}

func (users ordenarporidade) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

// Ordenar por sobrename
func (users ordenarporsobrenome) Len() int {
	return len(users)
}

func (users ordenarporsobrenome) Less(i, j int) bool {
	return users[i].Last < users[j].Last
}

func (users ordenarporsobrenome) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}
