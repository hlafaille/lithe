package lsp

import "reflect"

func StructToMap(s interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		out[field.Name] = v.Field(i).Interface()
	}
	return out
}
