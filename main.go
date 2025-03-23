package main

import (
	"log"
	"os"

	"github.com/gilberto-dev-silva/command-line.git/app"
)

func main() {
	aplication := app.Generate()
	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
