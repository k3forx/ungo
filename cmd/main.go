package main

import (
	"log"
	"os"

	"github.com/k3forx/ungo"
)

func main() {
	app := ungo.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
