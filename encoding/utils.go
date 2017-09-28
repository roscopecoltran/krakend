package encoding

import (
	"encoding/json"
	"errors"
	"fmt"

	// "github.com/fatih/camelcase"

	"github.com/clbanning/mxj"
	dtree "github.com/re-pe/go-dtree"
)

/*
	Refs:
	- https://github.com/diguinhorocks/api-handler/blob/master/util/response/response.go
*/

// splitted := camelcase.Split("GolangPackage")
// fmt.Println(splitted[0], splitted[1]) // prints: "Golang", "Package"

type Map map[string]interface{}

func To(t string, content map[string]interface{}) ([]byte, error) {
	if t == "xml" {
		m, err := mxj.AnyXml(content)
		if err != nil {
			panic(err)
		}
		return m, err
	} else if t == "json" {
		return json.Marshal(content)
	}
	return nil, nil
}

type EncodingFormats struct {
	JSON dtree.JsonHandler
	XML  dtree.XMLHandler
}

// func GetHandler

func GetNode(t string, node string, content map[string]interface{}) (dtree.DTree, error) {
	var tree = &EncodingFormats{}
	if t == "xml" {
		fmt.Println("\nReading xml file:\n")
		tree.XML.Content = []byte(`{"Other" : {"a" : 0, "b" : 1, "c" : 2, "i0" : 3}}`)
		if err := tree.XML.Decode(); err != nil {
			fmt.Println("err:", err)
			return dtree.DTree{}, err
			fmt.Println("tree.Value:", tree.XML.Value)
		}
		fmt.Printf("File \"%v\" exists and was read.\n\n", tree.XML.FileName)
		fmt.Printf("Final tree.Value: \"%v\"", tree.XML.Value)
		return tree.XML.Get(node), nil
	} else if t == "json" {
		fmt.Println("\nReading json file:\n")
		tree.JSON.Content = []byte(`{"Other" : {"a" : 0, "b" : 1, "c" : 2, "i0" : 3}}`)
		if err := tree.JSON.Decode(); err != nil {
			fmt.Println("err:", err)
			return dtree.DTree{}, err
		}
		fmt.Println("tree.Value:", tree.JSON.Value)
		fmt.Printf("File \"%v\" exists and was read.\n\n", tree.JSON.FileName)
		fmt.Printf("Final tree.Value: \"%v\"", tree.JSON.Value)
		return tree.JSON.Get(node), nil
	}
	return dtree.DTree{}, errors.New("Unkown encoding format submitted.")
}

func GetJSONValuesForPath(path string) {
	fmt.Println("path:", path)
}

func GetXMLValuesForPath(path string) {
	fmt.Println("path:", path)
}

//func SetNode(t string, content map[string]interface{}) ([]byte, error) {
//	result = tree.Set("NewArr.2.1", tree.NewValue(`"new_value"`))
//}

/*
func ConvertJSON2XML() {
	jasondata := []byte(`[
		{ "somekey":"somevalue" },
		"string",
		3.14159265,
		true
	]`)
	var i interface{}
	err := json.Unmarshal(jsondata, &i)
	if err != nil {
		// do something
	}
	x, err := anyxml.XmlIndent(i, "", "  ", "mydoc")
	if err != nil {
		// do something else
	}
	fmt.Println(string(x))
}
*/
