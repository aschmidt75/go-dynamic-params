package main

import "fmt"

// Resolver is able to get a value for a given key
type Resolver interface {
	Resolve(key string) (string, error)
}

// ResolverChain contains multiple resolvers
type ResolverChain []Resolver

// ResolveError has details about resolve errors
type ResolveError struct {
	what string
	key  string
}

func (e *ResolveError) Error() string {
	return fmt.Sprintf("%s (for key=%s)", e.what, e.key)
}

// ResolveFromString takes a string and a resolver, and resolves
// all parameter references using the given resolver
func ResolveFromString(in string, resolvers ResolverChain) (string, error) {
	//fmt.Printf("resolv: in=%s\n", in)

	t := NewTokenizerFromString(in)
	tokens, err := t.Tokenize()
	if err != nil {
		return "", err
	}

	res := ""
	for _, token := range tokens {
		//fmt.Printf("%3d, %s (%d)\n", idx, token.part, token.tkType)

		if token.tkType == typeStaticPart {
			res = fmt.Sprintf("%s%s", res, token.part)
		}
		if token.tkType == typeParamPart {
			x := string(token.part)
			if token.withNestedParam {
				// recurse into resolving the whole thing in
				// case of nested params
				x, err = ResolveFromString(string(token.part), resolvers)
				if err != nil {
					return "", err
				}
			}

			resolveOk := false
			y := ""
			for _, resolver := range resolvers {
				y, err = resolver.Resolve(x)
				if err == nil {
					resolveOk = true
					break
				}
			}

			if resolveOk {
				res = fmt.Sprintf("%s%s", res, y)
			} else {
				return "", &ResolveError{what: "unable to look up parameter", key: x}
			}

		}
	}

	return res, nil
}
