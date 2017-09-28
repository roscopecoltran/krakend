package encoding

import (
	"fmt"

	tablib "github.com/agrison/go-tablib"
)

func AgnosticLoadJSON() {

	ds, err := tablib.LoadJSON([]byte(`[
	  {"age":90,"firstName":"John","lastName":"Adams"},
	  {"age":67,"firstName":"George","lastName":"Washington"},
	  {"age":83,"firstName":"Henry","lastName":"Ford"}
	]`))
	if err != nil {
		fmt.Println("err:", err)
	}

	xml, _ := ds.XML()
	fmt.Println(xml)

	csv, _ := ds.CSV()
	fmt.Println(csv)

}

func AgnosticYAML() {

	ds, err := tablib.LoadYAML([]byte(`- age: 90
	  firstName: John
	  lastName: Adams
	- age: 67
	  firstName: George
	  lastName: Washington
	- age: 83
	  firstName: Henry
	  lastName: Ford`))
	if err != nil {
		fmt.Println("err:", err)
	}

	json, _ := ds.JSON()
	fmt.Println(json)

	xml, _ := ds.XML()
	fmt.Println(xml)

	csv, _ := ds.CSV()
	fmt.Println(csv)

}

func AgnosticDataBook() {

	db := tablib.NewDatabook()
	// or loading a JSON content
	// db, err := tablib.LoadDatabookJSON([]byte(`...`))

	// or a YAML content
	db, err := tablib.LoadDatabookYAML([]byte(`- age: 90
	  firstName: John
	  lastName: Adams
	- age: 67
	  firstName: George
	  lastName: Washington
	- age: 83
	  firstName: Henry
	  lastName: Ford`))
	if err != nil {
		fmt.Println("err:", err)
	}

	// a dataset of presidents
	presidents, _ := tablib.LoadJSON([]byte(`[
	  {"Age":90,"First name":"John","Last name":"Adams"},
	  {"Age":67,"First name":"George","Last name":"Washington"},
	  {"Age":83,"First name":"Henry","Last name":"Ford"}
	]`))

	// a dataset of cars
	cars := tablib.NewDataset([]string{"Maker", "Model", "Year"})
	cars.AppendValues("Porsche", "991", 2012)
	cars.AppendValues("Skoda", "Octavia", 2011)
	cars.AppendValues("Ferrari", "458", 2009)
	cars.AppendValues("Citroen", "Picasso II", 2013)
	cars.AppendValues("Bentley", "Continental GT", 2003)

	// add the sheets to the Databook
	db.AddSheet("Cars", cars.Sort("Year"))
	db.AddSheet("Presidents", presidents.SortReverse("Age"))

	fmt.Println(db.JSON())

}
