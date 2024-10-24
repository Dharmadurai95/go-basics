package main

import (
	"fmt"
	"log"

	"github.com/Dharmadurai95/greetings"
)

func main() {

	log.SetPrefix("greeting: ")
	log.SetFlags(0)
	nameList := []string{"Gokulapriya", "Dharmadurai", "Havish", "Nishvanth"}
	message, err := greetings.Hellos(nameList)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
