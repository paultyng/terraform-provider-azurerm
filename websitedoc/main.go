package main

import (
	"go/token"
	"log"
)

func main() {
	g := &generator{
		fset:     token.NewFileSet(),
		provider: "azurerm",
	}

	err := g.Generate("azurerm")
	if err != nil {
		log.Fatal(err)
	}
}
