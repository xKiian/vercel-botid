package main

import (
	"botid/botid"
	"fmt"
	"log"
)

func main() {
	solver, err := botid.NewBotID("https://botid-testing-pi.vercel.app/149e9513-01fa-4fb0-aad4-566afd725d1b/2d206a39-8ed7-437e-a3be-862e0f06eea3/a-4-a/c.js?i=0&v=3&h=botid-testing-pi.vercel.app")
	if err != nil {
		log.Fatal(err)
	}

	token, err := solver.GenerateToken()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+]", token)

	fmt.Println(solver.Verify(token))
}
