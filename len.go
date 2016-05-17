package viewhelpers

import "reflect"

// Len - size of array
func Len(element interface{}) int {
	if element == nil {
		return 0
	}
	return reflect.ValueOf(element).Len()
}
