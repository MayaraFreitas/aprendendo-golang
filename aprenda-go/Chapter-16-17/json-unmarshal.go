package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	NomeDiferente string `json:"Nome"`
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {

	sb := []byte(
		`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":9999999999999.01}`,
		//`{"Nome":"Anakin","Sobrenome":"Skywalker","Idade":50,"Profissao":"Vilão","Contabancaria":50.21}`,
	)

	var jamesbond Pessoa

	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(jamesbond)
	fmt.Println(jamesbond.Profissao)
}
