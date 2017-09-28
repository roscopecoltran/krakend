package encoding

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/k0kubun/pp"
)

// JSON is the key for the json encoding
const JSON = "json"

// NewJSONDecoder return the right JSON decoder
func NewJSONDecoderOrig(isCollection bool) Decoder {
	if isCollection {
		return JSONCollectionDecoder
	}
	return JSONDecoderMXJ
}

// NewJSONDecoder2 return the right JSON decoder
func NewJSONDecoder(mode string) Decoder {
	fmt.Printf("NewJSONDecoder(..), mode: %s:\n", mode)
	switch mode {
	case "collection":
		return JSONCollectionDecoder
	case "mxj":
		return JSONDecoderMXJ
	}
	return JSONDecoder
}

// JSONDecoder implements the Decoder interface
func JSONDecoder(r io.Reader, v *map[string]interface{}, paths []map[string]string) error {
	fmt.Println("JSONDecoder(..), paths:\n")
	pp.Print(paths)
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(v)
}

// JSONCollectionDecoder implements the Decoder interface over a collection
func JSONCollectionDecoder(r io.Reader, v *map[string]interface{}, paths []map[string]string) error {
	fmt.Println("JSONCollectionDecoder(..), paths:\n")
	pp.Print(paths)
	var collection []interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": collection}
	return nil
}

// https://github.com/clbanning/mxj/blob/master/jpath.go
// https://github.com/clbanning/mxj/blob/master/json.go
func JSONDecoderMXJ(r io.Reader, v *map[string]interface{}, paths []map[string]string) error {
	fmt.Printf("JSONDecoderMXJ(..), paths:\n")
	pp.Print(paths)
	result := make(map[string]interface{}, len(paths)) // -1
	// m, merr := NewMapJsonReader(jdata[:len(jdata)-2])
	mv, merr := mxj.NewMapJsonReader(r)
	if merr != nil && merr != io.EOF {
		panic(merr)
		fmt.Printf("error while init merr: %s\n", merr.Error())
	}
	for _, mapping := range paths {
		var fieldName, pathStr string
		for _, fv := range mapping {
			pp.Printf(fv)
			if strings.ToLower(fv) == "name" {
				fieldName = fv
			}
			if strings.ToLower(fv) == "path" {
				pathStr = fv
			}
		}
		if fieldName == "" || pathStr == "" {
			return errors.New("field name or data path are missing for this path extraction")
		}
		node, err := mv.ValuesForPath(pathStr)
		if err != nil {
			panic(merr)
			fmt.Println("error while init mv.ValuesForPath(...),  err:\n", err)
			return err
		}
		result[fieldName] = node
		fmt.Println(fieldName, ":", result)
	}
	*(v) = mv
	return nil
}

/*

// Extract / modify Map values
paths := m.PathsForKey(key)
path := mv.PathForKeyShortest(key)
values, err := mv.ValuesForKey(key, subkeys)
values, err := mv.ValuesForPath(path, subkeys)
count, err := mv.UpdateValuesForPath(newVal, path, subkeys)

// Get everything at once, irrespective of path depth:
leafnodes := mv.LeafNodes()
leafvalues := mv.LeafValues()

// Map
newMap, err := mv.NewMap("oldKey_1:newKey_1", "oldKey_2:newKey_2", ..., "oldKey_N:newKey_N")
newXml, err := newMap.Xml()   // for example
newJson, err := newMap.Json() // ditto



*/

/*
	Refs:
	- https://github.com/clbanning/mxj/blob/master/examples/jpath.go

	// $.store.book[*].author   the authors of all books in the store
	list, err := m.ValuesForPath("store.book.author")
	if err != nil {
		fmt.Println("book author err:", err)
		return
	}
	fmt.Println("authors:", list)

	// $..author                all authors
	list, err = m.ValuesForKey("author")
	if err != nil {
		fmt.Println("author err:", err)
		return
	}
	fmt.Println("authors:", list)

	// $.store.*                all things in store, which are some books and a red bicycle.
	list, err = m.ValuesForKey("store")
	if err != nil {
		fmt.Println("store things err:", err)
	}
	fmt.Println("store things:", list)

	// /store//price        $.store..price           the price of everything in the store.
	list, err = m.ValuesForKey("price")
	if err != nil {
		fmt.Println("price of things err:", err)
	}
	fmt.Println("price of things:", list)

	// $..book[2]               the third book
	v, err := m.ValueForPath("store.book[2]")
	if err != nil {
		fmt.Println("price of things err:", err)
	}
	fmt.Println("3rd book:", v)

	// $..book[-1:]             the last book in order
	list, err = m.ValuesForPath("store.book")
	if err != nil {
		fmt.Println("list of books err:", err)
	}
	if len(list) <= 1 {
		fmt.Println("last book:", list)
	} else {
		fmt.Println("last book:", list[len(list)-1:])
	}

	// $..book[:2]              the first two books
	list, err = m.ValuesForPath("store.book")
	if err != nil {
		fmt.Println("list of books err:", err)
	}
	if len(list) <= 2 {
		fmt.Println("1st 2 books:", list)
	} else {
		fmt.Println("1st 2 books:", list[:2])
	}

	// $..book[?(@.isbn)]       filter all books with isbn number
	list, err = m.ValuesForPath("store.book", "isbn:*")
	if err != nil {
		fmt.Println("list of books err:", err)
	}
	fmt.Println("books with isbn:", list)

	// $..book[?(@.price<10)]   filter all books cheapier than 10
	list, err = m.ValuesForPath("store.book")
	if err != nil {
		fmt.Println("list of books err:", err)
	}
	var n int
	for _, v := range list {
		if v.(map[string]interface{})["price"].(float64) >= 10.0 {
			continue
		}
		list[n] = v
		n++
	}
	list = list[:n]
	fmt.Println("books with price < $10:", list)

	// $..*                     all Elements in XML document. All members of JSON structure.
	// 1st where values are not complex elements
	list = m.LeafValues()
	fmt.Println("list of leaf values:", list)

	// $..*                     all Elements in XML document. All members of JSON structure.
	// 2nd every value - even complex elements
	path := "*"
	list = make([]interface{}, 0)
	for {
		v, _ := m.ValuesForPath(path)
		if len(v) == 0 {
			break
		}
		list = append(list, v...)
		path = path + ".*"
	}
	fmt.Println("list of all values:", list)

*/
/*
func ArbitrayDecoder(r io.Reader, v *map[string]interface{}) error {
	mv, err := mxj.NewMapStruct(r)
	if err := mv.Struct(structPointer); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"results": result}
	return nil
}
*/

/*
	Refs:
	- https://github.com/sokool/scraper/blob/architecture/storage/json.go
	-
*/
/*
type json struct {
	name    string
	objects []map[string]interface{}
}

func (this *json) Add(o map[string]interface{}) {
	this.objects = append(this.objects, o)
}

func (this *json) Count() int {
	return len(this.objects)
}

func (this *json) Flush(name string) {
	handler, err := os.Create(name + ".json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	fmt.Fprintf(handler, "[")
	for _, object := range this.objects {
		bytes, _ := mxj.Map(object).Json()
		fmt.Fprintf(handler, "%s,\n", string(bytes))
	}
	fmt.Fprintf(handler, "]")

	//reset objects list?
	this.objects = []map[string]interface{}{}
}

func JSON(filename []string) Storage {
	return &json{
		name: filename[0],
	}
}
*/
