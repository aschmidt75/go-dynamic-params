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

	Convey("Nested structures should be resolved correctly", t, func() {
		var err error
		var r YAMLResolver
		var v string

		r, err = NewYAMLResolverFromString(`
first:
  string: "123"
  int: 1
  float: 3.75
  bool: true
second:
  string: "456"
  int: 2
  float: 4.75
  bool: false`)
		So(err, ShouldBeNil)

		v, err = r.Resolve("first.string")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "123")

		v, err = r.Resolve("first.int")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "1")

		v, err = r.Resolve("first.float")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "3.75")

		v, err = r.Resolve("first.bool")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "true")

		v, err = r.Resolve("second.string")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "456")

		v, err = r.Resolve("second.int")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "2")

		v, err = r.Resolve("second.float")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "4.75")

		v, err = r.Resolve("second.bool")

		So(err, ShouldBeNil)
		So(v, ShouldEqual, "false")
	})

}
