package main

import (
	"flag"
	"log"
	"os"
)

var ioLogger = log.New(os.Stdout, "IO:", log.Ltime)

func main() {
	var algorithm = flag.String("algo", "binary", "specify a search algorithm")

	flag.Parse()

	ioLogger.Printf("receiver user parametes:\n\talgo=%s", *algorithm)
}
