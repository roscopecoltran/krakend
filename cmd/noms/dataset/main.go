package main

import (
  "fmt"
  "os"

  "github.com/attic-labs/noms/go/spec"
)

func main() {
  sp, err := spec.ForDataset("http://localhost:8000::people")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Could not create dataset: %s\n", err)
    return
  }
  defer sp.Close()

  if _, ok := sp.GetDataset().MaybeHeadValue(); !ok {
    fmt.Fprintf(os.Stdout, "head is empty\n")
  }
}
