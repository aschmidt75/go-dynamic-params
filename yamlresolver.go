package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// YAMLResolver ...
type YAMLResolver struct {
	yamlIn []byte
	flat   map[string]string
	origin string
}

// NewYAMLResolverFromString creates a new YAMLResolver from
// a given input string
func NewYAMLResolverFromString(yamlIn string) (YAMLResolver, error) {
	fmt.Printf("in: %#v\n", yamlIn)

	res := YAMLResolver{yamlIn: []byte(yamlIn)}
	var f interface{}

	err := yaml.Unmarshal(res.yamlIn, &f)
	if err != nil {
		return res, err
	}
	fmt.Printf("%#v\n", f)

	// TODO: find out at runtime if this cast will work out
	m := f.(map[string]interface{})

	fmt.Printf("%#v\n", m)

	res.flat = Flatten(m)

	return res, nil
}

// Resolve looks up given key in the flattened YAML input
func (r YAMLResolver) Resolve(key string) (string, error) {
	v, ex := r.flat[key]

	if ex == false {
		return "", &ResolveError{what: "not found", key: key}
	}
	return v, nil
}
