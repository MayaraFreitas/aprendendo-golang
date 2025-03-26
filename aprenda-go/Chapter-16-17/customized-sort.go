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

type ordenarPorPotencia []carro
type ordenarPorConsumo []carro
type ordenarPorLucroProDonoDoPosto []carro

func main() {

	carros := []carro{
		carro{"Chevete", 50, 20},
		carro{"Fusca", 500, 10},
		carro{"Porsche", 300, 17},
	}

	fmt.Println(carros)

	fmt.Println("\n Ordenando por potencia:")
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)

	fmt.Println("\n Ordenando por menos consumo:")
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println(carros)

	fmt.Println("\n Ordenando por mais consumo:")
	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)

}

// Ordenando Por Potencia
func (carros ordenarPorPotencia) Len() int {
	return len(carros)
}
func (carros ordenarPorPotencia) Less(i, j int) bool {
	return carros[i].potencia < carros[j].potencia
}
func (carros ordenarPorPotencia) Swap(i, j int) {
	carros[i], carros[j] = carros[j], carros[i]
}

// Ordenando Por Consumo
func (carros ordenarPorConsumo) Len() int {
	return len(carros)
}
func (carros ordenarPorConsumo) Less(i, j int) bool {
	return carros[i].consumo < carros[j].consumo
}
func (carros ordenarPorConsumo) Swap(i, j int) {
	carros[i], carros[j] = carros[j], carros[i]
}

// Ordenando Por Consumo
func (carros ordenarPorLucroProDonoDoPosto) Len() int {
	return len(carros)
}
func (carros ordenarPorLucroProDonoDoPosto) Less(i, j int) bool {
	return carros[i].consumo > carros[j].consumo
}
func (carros ordenarPorLucroProDonoDoPosto) Swap(i, j int) {
	carros[i], carros[j] = carros[j], carros[i]
}
