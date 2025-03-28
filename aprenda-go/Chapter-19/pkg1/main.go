package main

import (
	"aprenda-go/Chapter-19/pkg2"
	"fmt"
)

func main() {
	fmt.Println("Essa é a função main.")
	Outro()
	pkg2.Pkg2()
}

// RUN: go run ./
