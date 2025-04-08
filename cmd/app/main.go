package main

import (
	"log"

	"github.com/Wefdzen/Gravitum/api/router"
)

func main() {
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
