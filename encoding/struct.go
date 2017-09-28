package encoding

/*
import (
	"fmt"
	"github.com/fatih/structs"
)
*/

/*
	Refs:
	- https://github.com/sokool/scraper/blob/architecture/storage/struct.go
*/

/*
type data struct {
	data []map[string]interface{}
}

func (this *data) Add(in map[string]interface{}) {
	this.data = append(this.data, in)
	fmt.Printf(in)
}

func (this *data) Flush(name string) {

}

func STRUCT(params []string) *data {
	return &data{}
}


type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

server := &Server{
	Name:    "gopher",
	ID:      123456,
	Enabled: true,
}

// Convert a struct to a map[string]interface{}
// => {"Name":"gopher", "ID":123456, "Enabled":true}
m := structs.Map(server)

// Convert the values of a struct to a []interface{}
// => ["gopher", 123456, true]
v := structs.Values(server)

// Convert the names of a struct to a []string
// (see "Names methods" for more info about fields)
n := structs.Names(server)

// Convert the values of a struct to a []*Field
// (see "Field methods" for more info about fields)
f := structs.Fields(server)

// Return the struct name => "Server"
n := structs.Name(server)

// Check if any field of a struct is initialized or not.
h := structs.HasZero(server)

// Check if all fields of a struct is initialized or not.
z := structs.IsZero(server)

// Check if server is a struct or a pointer to struct
i := structs.IsStruct(server)


*/
