package main

import (
	"fmt"
	"os"

	"github.com/attic-labs/noms/go/spec"
	"github.com/attic-labs/noms/go/types"
)

func main() {
	sp, err := spec.ForDataset("http://localhost:8000::people")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create dataset: %s\n", err)
		return
	}
	defer sp.Close()

	if headValue, ok := sp.GetDataset().MaybeHeadValue(); !ok {
		fmt.Fprintf(os.Stdout, "head is empty\n")
	} else {
		// type assertion to convert Head to List
		personList := headValue.(types.List)
		// type assertion to convert List Value to Struct
		personStruct := personList.Get(0).(types.Struct)
		// prints: Rickon
		fmt.Fprintf(os.Stdout, "given: %v\n", personStruct.Get("given"))
	}
}
