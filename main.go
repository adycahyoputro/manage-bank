package main

import (
	"log"

	"github.com/adycahyoputro/manage-bank/server"
)

func main() {
	// run the server
	if err := server.Run(); err != nil {
		log.Fatalf("error running the server: %s", err)
	}
}
