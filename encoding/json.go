package encoding

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	// "github.com/buger/jsonparser"
	"github.com/k0kubun/pp"
	"github.com/roscopecoltran/mxj"
)

/*
XPath                JSONPath                 Result
/store/book/author   $.store.book[*].author   the authors of all books in the store
//author             $..author                all authors
/store/*             $.store.*                all things in store, which are some books and a red bicycle.
/store//price        $.store..price           the price of everything in the store.
//book[3]            $..book[2]               the third book
//book[last()]       $..book[(@.length-1)]
                     $..book[-1:]             the last book in order.
//book[position()<3] $..book[0,1]
                     $..book[:2]              the first two books
//book[isbn]         $..book[?(@.isbn)]       filter all books with isbn number
//book[price<10]     $..book[?(@.price<10)]   filter all books cheapier than 10
//*                  $..*                     all Elements in XML document. All members of JSON structure.
*/

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

// https://golang.org/pkg/encoding/json/#NewDecoder
// JSONDecoder implements the Decoder interface
func JSONDecoder(r io.Reader, v *map[string]interface{}, targets []map[string]string) error {
	fmt.Printf("JSONDecoder(..)\n")
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(v)
}

// JSONCollectionDecoder implements the Decoder interface over a collection
func JSONCollectionDecoder(r io.Reader, v *map[string]interface{}, targets []map[string]string) error {
	fmt.Printf("JSONCollectionDecoder(..)\n")
	var collection []interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": collection}
	return nil
}

// https://github.com/clbanning/mxj/blob/master/examples/leafnodes.go

// Retrieve a Map value from an io.Reader.
//  NOTE: The raw JSON off the reader is buffered to []byte using a ByteReader. If the io.Reader is an
//        os.File, there may be significant performance impact. If the io.Reader is wrapping a []byte
//        value in-memory, however, such as http.Request.Body you CAN use it to efficiently unmarshal
//        a JSON object.
// https://github.com/clbanning/mxj/blob/master/jpath.go
// https://github.com/clbanning/mxj/blob/master/json.go
// https://github.com/clbanning/mxj/issues/21
func JSONDecoderMXJ(r io.Reader, v *map[string]interface{}, targets []map[string]string) error {
	result := make(map[string]interface{}, len(targets)) // -1
	mxj.JsonUseNumber = true
	mv, merr := mxj.NewMapJsonReaderAll(r)
	if merr != nil {
		return merr
	}

	mxj.LeafUseDotNotation()
	l := m.LeafNodes()
	for _, v := range l {
		fmt.Println("path:", v.Path, "value:", v.Value)
	}

	// fmt.Printf("NewMapJsonReader, mv : %#v\n", mv)
	if len(targets) > 0 {
		for _, target := range targets {
			var newFieldName, targetName, targetPath string
			for kv, fv := range target {
				// fmt.Printf("JSONDecoderMXJ(..) > kv=%s, fv=%s\n", kv, fv)

				/*
					if strings.ToLower(kv) == "subkeys" {
						subkeys = fv
						subKeysList = strings.Split(",", subkeys)
					}
				*/
				if strings.ToLower(kv) == "target" {
					targetName = fv
				}
				if strings.ToLower(kv) == "rename" {
					newFieldName = fv
				}
				if strings.ToLower(kv) == "path" {
					targetPath = fv
				}
			}
			if targetName != "" || targetPath != "" {
				// fmt.Printf("targetName: %s, newFieldName: %s, targetPath: %s", targetName, newFieldName, targetPath)
				if newFieldName != "" && targetName != "" {
					err := mv.RenameKey(targetName, newFieldName)
					if err != nil {
						return err
					}
				}
				/*
					if len(subKeysList) > 0 {
						node, err := mv.ValuesForPath(newFieldName, subKeysList...)
						if err != nil {
							return err
						}
					} else {
						node, err := mv.ValuesForPath(newFieldName)
						if err != nil {
							return err
						}
					}
				*/
				node, err := mv.ValuesForPath(newFieldName)
				if err != nil {
					return err
				}
				/*
					newFieldKeys := strings.Split(newFieldName, ".")
					if len(newFieldKeys) > 1 {
						for _, k := range newFieldKeys {
							result[k] =
						}
					} else {
				*/
				result[newFieldName] = node
				//}
			}
		}
		pp.Print(result)
		*(v) = result
		return nil
	}
	*(v) = mv
	return nil
}

/*

// Extract / modify Map values
targets := m.PathsForKey(key)
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
