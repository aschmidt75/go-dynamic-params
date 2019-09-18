package main

type MapResolver struct {
	m map[string]string
}

func NewMapResolver() MapResolver {
	res := MapResolver{}
	res.m = make(map[string]string)
	return res
}

func (r MapResolver) With(in map[string]string) MapResolver {
	for k, v := range in {
		r.m[k] = v
	}
	return r
}

func (r MapResolver) Resolve(key string) (string, error) {
	v, ex := r.m[key]

	if ex == false {
		return "", &ResolveError{what: "not found", key: key}
	}
	return v, nil
}
