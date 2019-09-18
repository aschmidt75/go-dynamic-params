package main

import "fmt"

type Resolver interface {
	Resolve(key string) (string, error)
}

// ResolveError is
type ResolveError struct {
	what string
	key  string
}

func (e *ResolveError) Error() string {
	return fmt.Sprintf("%s (for key=%s)", e.what, e.key)
}

func ResolveFromString(in string, r Resolver) (string, error) {
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
				x, err = ResolveFromString(string(token.part), r)
				if err != nil {
					return "", err
				}
			}
			y, err := r.Resolve(x)
			if err != nil {
				return "", err
			}
			res = fmt.Sprintf("%s%s", res, y)

		}
	}

	return res, nil
}
