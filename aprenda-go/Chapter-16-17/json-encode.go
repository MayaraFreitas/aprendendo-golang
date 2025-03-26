package main

import (
	"encoding/json"
	"os"
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

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)
}
