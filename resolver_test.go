package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResolverValidInputs(t *testing.T) {
	Convey("Non-param strings should be resolved", t, func() {
		res, err := ResolveFromString("", ResolverChain{})
	
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "")

		res, err = ResolveFromString("", ResolverChain{NewMapResolver()})
	
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "")

		res, err = ResolveFromString("no.params.here", ResolverChain{NewMapResolver()})
	
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "no.params.here")
	})

	Convey("Sample resolve testcases", t, func() {
		res, err := ResolveFromString("key=${value}", ResolverChain{NewMapResolver().With(map[string]string{
			"value":    "123",
		})})
	
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "key=123")

		res, err = ResolveFromString("key=${${value}}", ResolverChain{NewMapResolver().With(map[string]string{
			"value":    "another.value",
			"another.value": "123",
		})})
	
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "key=123")

		res, err = ResolveFromString("${_key}=${_value}", ResolverChain{NewMapResolver().With(map[string]string{
			"_value":    "another.value",
			"_key": "another.key",
		})})
	
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "another.key=another.value")

		multilineTemplate := `
test:
		key: ${value}
		other-key: other-value
		`
		multilineResult := `
test:
		key: 123
		other-key: other-value
		`
		res, err = ResolveFromString(multilineTemplate, ResolverChain{NewMapResolver().With(map[string]string{
			"value":    "123",
		})})
	
		So(err, ShouldBeNil)
		So(res, ShouldEqual, multilineResult)

	})

}

func TestResolverInvalidInputs(t *testing.T) {
	Convey("Resolve with empty keys should return errors", t, func() {
		_, err := ResolveFromString("key=${}", ResolverChain{NewMapResolver().With(map[string]string{
			"another.value":    "123",
		})})
	
		So(err, ShouldNotBeNil)

		_, err = ResolveFromString("key=${${value}}", ResolverChain{NewMapResolver().With(map[string]string{
			"value":    "",
		})})
	
		So(err, ShouldNotBeNil)
	})

	Convey("Resolve with nonexisting keys should return errors", t, func() {
		_, err := ResolveFromString("key=${no.such.value}", ResolverChain{NewMapResolver().With(map[string]string{
			"another.value":    "123",
		})})
	
		So(err, ShouldNotBeNil)
	})

	Convey("Resolve with empty resolver chain should return errors", t, func() {
		_, err := ResolveFromString("key=${no.resolvers.in.chain}", ResolverChain{})
	
		So(err, ShouldNotBeNil)
	})
}

func TestResolverValidInputsWithChains(t *testing.T) {
	Convey("Resolve should pick distinct parameters from the whole chain correctly", t, func() {
		rc := ResolverChain{
			NewMapResolver().With(map[string]string{
				"_value":    "another.value",
			}),
			NewMapResolver().With(map[string]string{
				"_key": "another.key",
			}),
		}
		res, err := ResolveFromString("${_key}=${_value}", rc)

		So(err, ShouldBeNil)
		So(res, ShouldEqual, "another.key=another.value")
	})

	Convey("Resolve should pick identical parameters from the whole chain correctly", t, func() {
		rc := ResolverChain{
			NewMapResolver().With(map[string]string{
				"_value":    "another.value",
			}),
			NewMapResolver().With(map[string]string{
				"_value":    "another.value.but.hidden.because.first.chainelement.already.has.it",
				"_key": "another.key",
			}),
		}
		res, err := ResolveFromString("${_key}=${_value}", rc)

		So(err, ShouldBeNil)
		So(res, ShouldEqual, "another.key=another.value")
	})
}
