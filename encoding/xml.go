package encoding

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/k0kubun/pp"
	//	"github.com/roscopecoltran/krakend/logging"
)

// XML is the key for the xml encoding
const XML = "xml"

// NewXMLDecoder return the right XML decoder
func NewXMLDecoderOrig(isCollection bool) Decoder {
	if isCollection {
		return XMLCollectionDecoder
	}
	return XMLDecoder
}

// NewXMLDecoder2 return the right XML decoder
func NewXMLDecoder(mode string) Decoder {
	fmt.Printf("NewXMLDecoder(..), mode: %s\n", mode)
	switch mode {
	case "collection":
		return XMLCollectionDecoder
	case "mxj":
		return XMLDecoderMXJ
	}
	return XMLDecoder
}

// XMLDecoder implements the Decoder interface
func XMLDecoder(r io.Reader, v *map[string]interface{}, paths []map[string]string) error {
	fmt.Printf("XMLDecoder(..), paths:\n")
	pp.Print(paths)
	mv, err := mxj.NewMapXmlReader(r)
	if err != nil {
		return err
	}
	*v = mv
	return nil
}

// XMLCollectionDecoder implements the Decoder interface over a collection
func XMLCollectionDecoder(r io.Reader, v *map[string]interface{}, paths []map[string]string) error {
	fmt.Printf("XMLCollectionDecoder(..), paths:\n")
	pp.Print(paths)
	mv, err := mxj.NewMapXmlReader(r)
	if err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": mv}
	return nil
}

// XMLCollectionDecoder implements the Decoder interface over a collection
func XMLDecoderMXJ(r io.Reader, v *map[string]interface{}, paths []map[string]string) error {
	fmt.Printf("XMLDecoderMXJ(..), paths:\n")
	pp.Print(paths)
	result := make(map[string]interface{}, len(paths)) //-1)
	// m, merr := NewMapXmlReader(jdata[:len(jdata)-2])
	mv, err := mxj.NewMapXmlReader(r)
	if err != nil {
		fmt.Println("error while init NewMapJsonReader,  err:\n", err)
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
	// Attributes are keys with prepended hyphen, '-'.
	values, err := m.ValuesForPath("doc.some_tag.geoInfo.country.-name")
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	for _, v := range values {
		fmt.Println("v:", v)
	}
*/

/*
	// extract the 'list3-1-1-1' node - there'll be just 1?
	// NOTE: attribute keys are prepended with '-'
	lstVal, lerr := m.ValuesForPath("*.*.*.*.*", "-name:list3-1-1-1")
	if lerr != nil {
		fmt.Println("ierr:", lerr.Error())
		return
	}

	// assuming just one value returned - create a new Map
	mv := mxj.Map(lstVal[0].(map[string]interface{}))

	// extract the 'int' values by 'name' attribute: "-name"
	// interate over list of 'name' values of interest
	var names = []string{"field1", "field2", "field3", "field4", "field5"}
	for _, n := range names {
		vals, verr := mv.ValuesForKey("int", "-name:"+n)
		if verr != nil {
			fmt.Println("verr:", verr.Error(), len(vals))
			return
		}
		if len(vals) == 0 { // good to check to avoid PANIC
			continue
		}

		// values for simple elements have key '#text'
		// NOTE: there can be only one value for key '#text'
		fmt.Println(n, ":", vals[0].(map[string]interface{})["#text"])
	}
*/

/*
func XMLPathDecoder(r io.Reader, path string, v *map[string]interface{}) error {
	list, err := m.ValuesForPath(path)
	if err != nil {
		fmt.Println("book author err:", err)
		return
	}
	fmt.Println("authors:", list)
}
*/

/*
	Refs:
	- https://github.com/sokool/scraper/blob/architecture/storage/xml.go
	- https://github.com/dspencerr/Go-Open-XML-Parser/blob/master/xmlParser/mxjUnwrap.go
	- https://github.com/stormasm/nomstack/blob/master/samples/go/xml-import/xml_importer.go
	- https://github.com/stormasm/nomstack/blob/master/samples/go/noms-merge/main.go
	- https://github.com/andystanton/proxybastard/blob/master/proxy/maven.go
	- https://github.com/nonda/crowdinsheets/blob/master/xml2csv.go
	-

*/

/*
type xml struct {
	name string
	list []map[string]interface{}
}

func (this *xml) Add(o map[string]interface{}) {
	this.list = append(this.list, o)

}

func (this *xml) Count() int {
	return len(this.list)
}

func (this *xml) Flush(name string) {
	handler, err := os.Create(name + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	fmt.Fprintf(handler, "<%s>\n", "objects")
	for _, data := range this.list {
		o := mxj.Map(data)
		bytes, _ := o.Xml()
		fmt.Fprintf(handler, "%s\n", string(bytes))
	}
	fmt.Fprintf(handler, "</%s>\n", "objects")

}

func XML(filename []string) Storage {
	return &xml{
		name: filename[0],
	}
}
*/
