package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	senha := "que_senha_mais_fraca"
	senhaerrada := ":D"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))
}
