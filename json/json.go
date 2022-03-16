package json

import (
	"encoding/json"
	"reflect"

	"github.com/go-python/gpython/py"
)

const dumps_doc = `dumps() -> str

Serialize obj to a JSON formatted str
`

const module_doc = `

Interaction with JSON.

`

// serialize StringDict into JSON string
func json_dumps(self py.Object, args py.Tuple) (py.Object, error) {
	if len(args) == 0 {
		return nil, py.ExceptionNewf(py.IndexError, "Too few arguments")
	} else {
		if reflect.TypeOf(args[0]).String() == "py.StringDict" {
			data, err := json.Marshal(args[0])
			if err != nil {
				return nil, py.ExceptionNewf(py.BaseException, "Failed to marshal JSON.")
			}
			return py.String(string(data)), nil
		} else {
			return nil, py.ExceptionNewf(py.TypeError, "Wrong type for json.dumps")
		}
	}
}

// Initialise the module
func init() {
	methods := []*py.Method{
		py.MustNewMethod("dumps", json_dumps, 0, dumps_doc),
	}

	py.RegisterModule(&py.ModuleImpl{
		Info: py.ModuleInfo{
			Name: "json",
			Doc:  module_doc,
		},
		Methods: methods,
		Globals: py.StringDict{},
	})

}
