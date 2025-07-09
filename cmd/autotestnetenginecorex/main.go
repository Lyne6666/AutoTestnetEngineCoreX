// cmd/autotestnetenginecorex/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"autotestnetenginecorex/internal/autotestnetenginecorex"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := autotestnetenginecorex.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
