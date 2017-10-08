package main

import (
	"fmt"
	"os"

	"github.com/attic-labs/noms/go/spec"
	"github.com/attic-labs/noms/go/types"
)

func newPerson(givenName string, male bool) types.Struct {
	return types.NewStruct("Person", types.StructData{
		"given": types.String(givenName),
		"male":  types.Bool(male),
	})
}

func main() {
	sp, err := spec.ForDataset("http://localhost:8000::people")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create dataset: %s\n", err)
		return
	}
	defer sp.Close()

	db := sp.GetDatabase()

	data := types.NewList(db,
		newPerson("Rickon", true),
		newPerson("Bran", true),
		newPerson("Arya", false),
		newPerson("Sansa", false),
	)

	fmt.Fprintf(os.Stdout, "data type: %v\n", types.TypeOf(data).Describe())

	_, err = db.CommitValue(sp.GetDataset(), data)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error commiting: %s\n", err)
	}
}
