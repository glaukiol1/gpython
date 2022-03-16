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

var (
	JSON_DICT_TYPE = py.NewType("JsonDict", "JSON Key-Value Object")
)

// serialize StringDict into JSON string
func json_dumps(self py.Object, args py.Tuple) (py.Object, error) {
	if len(args) == 0 {
		return nil, py.ExceptionNewf(py.TypeError, "Too few arguments")
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

// de-serialize JSON formatted string into JsonDict
func json_loads(self py.Object, args py.Tuple) (py.Object, error) {
	if len(args) == 0 {
		return nil, py.ExceptionNewf(py.TypeError, "Too few arguments")
	} else {
		if reflect.TypeOf(args[0]).String() == "py.String" {
			var result map[string]interface{}
			arg, err := py.ReprAsString(args[0])
			if err != nil {
				return nil, py.ExceptionNewf(py.TypeError, "Failed to parse argument into Go string.")
			}
			err = json.Unmarshal([]byte(arg), result)
			if err != nil {
				return nil, py.ExceptionNewf(py.KeyError, "Failed to parse JSON")
			}
			parsedDict := JSON_DICT_TYPE.Alloc()
			for k, v := range result {
				switch c := v.(type) {
				case string:
					parsedDict.Dict[k] = py.String(c)
				}
			}
			return parsedDict, nil
		} else {
			return nil, py.ExceptionNewf(py.KeyError, "Argument must be of type string")
		}
	}
}

// Initialise the module
func init() {
	methods := []*py.Method{
		py.MustNewMethod("dumps", json_dumps, 0, dumps_doc),
		py.MustNewMethod("loads", json_loads, 0, ""),
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
