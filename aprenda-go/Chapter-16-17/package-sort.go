package main

import (
	"fmt"
	"sort"
)

func main() {
	sortString()
	sortInts()

}

func sortString() {
	fmt.Println("\nSorting String:")
	x := []string{"azul", "verde", "marrom", "amarelo", "laranja"}
	fmt.Println(x)

	sort.Strings(x)
	fmt.Println(x)
}

func sortInts() {
	fmt.Println("\nSorting Ints:")
	x := []int{54, 12, 78, 300, 5}
	fmt.Println(x)

	sort.Ints(x)
	fmt.Println(x)
}
