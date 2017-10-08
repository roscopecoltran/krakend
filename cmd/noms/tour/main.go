package main

import (
	"fmt"
	"os"

	"github.com/attic-labs/noms/go/spec"
)

func main() {
	sp, err := spec.ForDatabase("http://localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not access database: %s\n", err)
		return
	}
	defer sp.Close()
}
