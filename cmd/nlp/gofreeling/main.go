package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/advancedlogic/go-freeling/lib"
	_ "github.com/advancedlogic/go-freeling/models"
)

func main() {
	document := new(DocumentEntity)
	analyzer := NewAnalyzer()
	document.Content = "Hello World"
	output := analyzer.AnalyzeText(document)

	js := output.ToJSON()
	b, err := json.Marshal(js)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
