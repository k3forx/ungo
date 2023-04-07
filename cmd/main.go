package main

import (
	"log"
	"os"

	"github.com/k3forx/ungo"
)

func main() {
	if err := ungo.NewApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
