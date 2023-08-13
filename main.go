package main

import (
	"github.com/avtara/boilerplate-go/app"
	"log"
)

func main() {
	instance := app.New()

	err := instance.Start()
	if err != nil {
		log.Panic(err)
	}
}
