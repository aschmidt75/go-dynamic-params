package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJSONResolverBasics(t *testing.T) {
	Convey("empty input should result in error", t, func() {
		var err error
		var r JSONResolver

		_, err = NewJSONResolverFromString("")

		So(err, ShouldNotBeNil)

		r, err = NewJSONResolverFromString("{}")
		_, err = r.Resolve("no.such.value")

		So(err, ShouldNotBeNil)

	})

	Convey("basic values should be resolved correctly", t, func() {
		var err error
		var r JSONResolver
		var v string

		r, err = NewJSONResolverFromString(`
		{ "string": "123",
		"int": 1,
		"float": 3.75,
		"bool": true
		}
		`)
		So(err, ShouldBeNil)

		v, err = r.Resolve("string")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "123")

		v, err = r.Resolve("int")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "1")

		v, err = r.Resolve("float")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "3.75")

		v, err = r.Resolve("bool")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "true")
	})
}
