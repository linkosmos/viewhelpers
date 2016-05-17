package viewhelpers

import "errors"

// -
var (
	ErrDictKeysMustBeString = errors.New("dict keys must be strings")
)

// Dict - allows to pass multiple values to template
func Dict(values ...interface{}) (map[string]interface{}, error) {
	size := len(values)
	if size%2 != 0 {
		panic("Values passed must be even combination")
	}
	dict := make(map[string]interface{}, size/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, ErrDictKeysMustBeString
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
