// cmd/quantumvelocity/main.go
package main

import (
	"flag"
	"log"
	"os"

	"quantumvelocity/internal/quantumvelocity"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := quantumvelocity.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
