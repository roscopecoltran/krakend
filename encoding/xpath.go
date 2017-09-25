package encoding

/*
	refs:
	- https://github.com/antchfx/xquery/tree/master/xml
	- https://github.com/ChrisTrenkamp/goxpath
	- https://github.com/santhosh-tekuri/xpathparser
	- https://github.com/antchfx/xpath

	examples:
	- https://github.com/antchfx/xquery/blob/master/examples/bing/main.go
	- https://github.com/antchfx/xquery/blob/master/examples/amazon/main.go
*/

//import (
//	"github.com/roscopecoltran/krakend/logging"
// 	"gopkg.in/xmlpath.v2"
//)

// XPATH is the key for the xpath encoding
const XPATH = "xpath"

/*
// NewJSONDecoder return the right JSON decoder
func NewXPATHDecoder(isCollection bool) Decoder {
	if isCollection {
		return XPATHCollectionDecoder
	}
	return XPATHDecoder
}
*/

/*
// JSONDecoder implements the Decoder interface
func XPATHDecoder(r io.Reader, v *map[string]interface{}) error {
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(v)
}

// JSONCollectionDecoder implements the Decoder interface over a collection
func XPATHCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	var collection []interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": collection}
	return nil
}

	node, err := xmlpath.ParseHTML(reader)
	if err != nil {
		log.Fatalln(err)
	}

	path, err := xmlpath.Compile(xpath)
	if err != nil {
		log.Fatalln(err)
	}

	// output

	it := path.Iter(node)
	for it.Next() {
		scrapText := it.Node().String()

		fmt.Println(scrapText)
	}

*/
