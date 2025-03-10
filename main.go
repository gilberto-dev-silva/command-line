package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gilberto-dev-silva/command-line.git/app"
)

func main() {
	fmt.Println("Projeto consulta domínios em Go")

	aplication := app.Generate()
	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
