package main

import (
	"errors"

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

	res := YAMLResolver{yamlIn: []byte(yamlIn)}
	var f interface{}

	err := yaml.Unmarshal(res.yamlIn, &f)
	if err != nil {
		return res, err
	}

	// f is map[interface{}]interface{} but needs to be map[string]interface{}
	// transform, otherwise flatten won't work.
	m := make(map[string]interface{})
	for k, v := range f.(map[interface{}]interface{}) {
		switch kk := k.(type) {
		case string:
			m[kk] = v
		default:
			return res, errors.New("Keys need to be strings")
		}
	}

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
