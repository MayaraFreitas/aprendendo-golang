package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {

	jamesbond := Pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 9999999999999.01,
	}

	darthvader := Pessoa{"Anakin", "Skywalker", 50, "Vilão", 50.21}

	j, err := json.Marshal(jamesbond)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(j))
	}

	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(d))
	}
}
