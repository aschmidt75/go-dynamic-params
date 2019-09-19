package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestYAMLResolverBasics(t *testing.T) {

	Convey("basic values should be resolved correctly", t, func() {
		var err error
		var r YAMLResolver
		var v string

		r, err = NewYAMLResolverFromString(`
---
string: "123"
int: 1
float: 3.75
bool: true`)
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
