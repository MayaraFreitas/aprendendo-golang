package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

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

	fmt.Printf("\n\nENCODE:\n")
	encodeError := json.NewEncoder(os.Stdout).Encode(users)
	if encodeError != nil {
		fmt.Println("Decode Erro: ", encodeError)
	}

	fmt.Printf("\n\nDECODE:\n")

	const jsonUsers = `
			{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]}
			{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]}
			{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}
	`

	decode := json.NewDecoder(strings.NewReader(jsonUsers))
	for {
		var usersDecoded user
		if decodeError := decode.Decode(&usersDecoded); decodeError == io.EOF {
			break
		} else if decodeError != nil {
			fmt.Println("Decode Erro: ", decodeError)
		}
		fmt.Println(usersDecoded)
	}

}
