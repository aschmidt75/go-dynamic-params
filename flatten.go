package main

import (
	"fmt"
)

// Flatten takes a map from string to anything (e.g. from JSON or YAML), flattens it
// and returns a map of strings. All entries are in the top, levels separated by '.'
func Flatten(m map[string]interface{}) map[string]string {
	res := make(map[string]string)
	if m != nil {
		flattenStruct("", &res, m)
	}
	return res
}

func flattenStruct(prefix string, res *map[string]string, m map[string]interface{}) {
	for k, v := range m {
		compoundKey := fmt.Sprintf("%s%s", prefix, k)
		switch vv := v.(type) {
		case string:
			(*res)[compoundKey] = vv
		case float64:
			(*res)[compoundKey] = fmt.Sprintf("%v", vv)
		case bool:
			(*res)[compoundKey] = fmt.Sprintf("%v", vv)
		case []interface{}:
			flattenList(fmt.Sprintf("%s%s.", prefix, k), res, vv)
		case map[string]interface{}:
			flattenStruct(fmt.Sprintf("%s%s.", prefix, k), res, vv)
		}
	}
}

func flattenList(prefix string, res *map[string]string, m []interface{}) {
	for i, v := range m {
		compoundKey := fmt.Sprintf("%s%d", prefix, i)
		switch vv := v.(type) {
		case string:
			(*res)[compoundKey] = vv
		case float64:
			(*res)[compoundKey] = fmt.Sprintf("%v", vv)
		case bool:
			(*res)[compoundKey] = fmt.Sprintf("%v", vv)
		case []interface{}:
			flattenList(compoundKey, res, vv)
		case map[string]interface{}:
			flattenStruct(fmt.Sprintf("%s.", compoundKey), res, vv)
		}
	}
}
