package main

import (
	"code_competence/routes"
)

func main() {

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
