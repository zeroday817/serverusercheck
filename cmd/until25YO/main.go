package main

import (
	"GOLANG_PROJECTS/internal/pkg/app"
	"log"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
