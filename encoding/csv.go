package encoding

import (
	// "encoding/xml"
	"fmt"
	"strings"

	"github.com/agrison/go-tablib"
)

func escapeTranslationString(src string) string {
	escaped := strings.Replace(src, "\n", "\\n", -1)
	escaped = strings.Replace(escaped, "\"", "\\\"", -1)
	return escaped
}

// Csv2Strings convert input CSV content to iOS strings file
func Csv2Strings(csvSource []byte) (string, error) {
	csv, err := tablib.LoadCSV(csvSource)
	if err != nil {
		return "", err
	}

	var results []string
	var currentRow map[string]interface{}

	var i int
	for {
		currentRow, err = csv.Row(i)
		if err != nil {
			break
		}

		source, ok := currentRow["Source"].(string)
		translation, ok := currentRow["English"].(string)
		if !ok {
			i++
			continue
		}

		source = escapeTranslationString(source)
		translation = escapeTranslationString(translation)
		line := fmt.Sprintf("\"%s\" = \"%s\";", source, translation)
		results = append(results, line)
		i++
	}
	return strings.Join(results, "\n"), nil
}

/*

type root struct {
	AndroidStrings      []AndroidString      `xml:"string"`
	AndroidStringsArray []AndroidStringArray `xml:"string-array"`
	AndroidIntArray     []AndroidIntArray    `xml:"integer-array"`
}

// AndroidString contains a single translation of a string
type AndroidString struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

// AndroidArrayItem is just an item in a string-array or int-array
type AndroidArrayItem struct {
	Value string `xml:",innerxml"`
}

// AndroidStringArray contains multiple strings that represent the same key
type AndroidStringArray struct {
	Key   string             `xml:"name,attr"`
	Items []AndroidArrayItem `xml:"item"`
}

// AndroidIntArray contains multiple int items that represent the same key
type AndroidIntArray struct {
	Key   string             `xml:"name,attr"`
	Items []AndroidArrayItem `xml:"item"`
}

// https://github.com/nonda/crowdinsheets/blob/master/xml2csv.go
// Xml2Csv converts android.xml to android.csv
func Xml2Csv(xmlData []byte) ([]byte, error) {
	var xmlRoot root

	if err := xml.Unmarshal(xmlData, &xmlRoot); err != nil {
		return []byte{}, err
	}

	csv := tablib.NewDataset([]string{"key", "translation"})

	for _, s := range xmlRoot.AndroidStrings {
		csv.AppendValues(s.Key, s.Value)
	}

	csvOutput, err := csv.CSV()
	if err != nil {
		return []byte{}, err
	}

	return csvOutput.Bytes(), nil
}
*/
