package main

import (
	"log"

	"github.com/39shin52/todoAPI/app/interfaces"
)

func main() {
	r := interfaces.NewServer()
	err := r.Init()
	if err != nil {
		log.Fatal(err)
	}

	r.Router.Run(":9000")
}
