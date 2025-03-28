package main

import "fmt"

type pessoa struct {
	nome string
}

type humanos interface {
	falar()
}

func main() {

	p1 := pessoa{"Mayara"}
	// (&p1).falar()
	dizerAlgumaCoisa(&p1)
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "disse oi")

}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}
