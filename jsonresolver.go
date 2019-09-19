package main

import (
	"encoding/json"
)

// JSONResolver ...
type JSONResolver struct {
	jsonIn []byte
	flat   map[string]string
	origin string
}

// NewJSONResolverFromString creates a new JSONResolver from
// a given input string
func NewJSONResolverFromString(jsonIn string) (JSONResolver, error) {
	res := JSONResolver{jsonIn: []byte(jsonIn)}
	var f interface{}

	err := json.Unmarshal(res.jsonIn, &f)
	if err != nil {
		return res, err
	}

	m := f.(map[string]interface{})

	res.flat = Flatten(m)

	return res, nil
}

// Resolve looks up given key in the flattened JSON input
func (r JSONResolver) Resolve(key string) (string, error) {
	v, ex := r.flat[key]

	if ex == false {
		return "", &ResolveError{what: "not found", key: key}
	}
	return v, nil
}
