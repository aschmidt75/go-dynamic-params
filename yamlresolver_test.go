package dynp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestYAMLResolverBasics(t *testing.T) {

	Convey("basic values should be resolved correctly", t, func(c C) {
		var err error
		var r YAMLResolver
		var v string

		r, err = NewYAMLResolverFromString(`
---
string: "123"
int: 1
float: 3.75
bool: true`)
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

	Convey("Nested structures should be resolved correctly", t, func(c C) {
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
		c.So(err, ShouldBeNil)

		v, err = r.Resolve("first.string")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "123")

		v, err = r.Resolve("first.int")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "1")

		v, err = r.Resolve("first.float")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "3.75")

		v, err = r.Resolve("first.bool")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "true")

		v, err = r.Resolve("second.string")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "456")

		v, err = r.Resolve("second.int")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "2")

		v, err = r.Resolve("second.float")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "4.75")

		v, err = r.Resolve("second.bool")

		c.So(err, ShouldBeNil)
		c.So(v, ShouldEqual, "false")
	})

}
