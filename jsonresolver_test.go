package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJSONResolverBasics(t *testing.T) {
	Convey("empty input should result in error", t, func(c C) {
		var err error
		var r JSONResolver

		_, err = NewJSONResolverFromString("")

		c.So(err, ShouldNotBeNil)

		r, err = NewJSONResolverFromString("{}")
		_, err = r.Resolve("no.such.value")

		c.So(err, ShouldNotBeNil)

	})

	Convey("basic values should be resolved correctly", t, func(c C) {
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
		c.So(err, ShouldBeNil)

		v, err = r.Resolve("string")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "123")

		v, err = r.Resolve("int")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "1")

		v, err = r.Resolve("float")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "3.75")

		v, err = r.Resolve("bool")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "true")
	})
}
