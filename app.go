package main

import (
	"github.com/fiscaluno/aiolia/institution"
	"github.com/fiscaluno/aiolia/server"
)

func main() {
	institution.Migrate()
	server.Listen()
}
