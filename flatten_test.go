package main

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func FlattenFromJSON(jsonIn string) map[string]string {
	var f interface{}
	json.Unmarshal([]byte(jsonIn), &f)

	m := f.(map[string]interface{})

	return Flatten(m)
}

func TestFlattening(t *testing.T) {
	Convey("empty input should result in empty map", t, func(c C) {
		em := make(map[string]interface{})
		r := Flatten(em)

		c.So(len(r), ShouldEqual, 0)
	})

	Convey("basic types should be handled correctly", t, func(c C) {
		r := FlattenFromJSON(`
{ "string": "123",
"int": 1,
"float": 3.75,
"bool": true
}
		`)

		c.So(len(r), ShouldEqual, 4)
		c.So(r["string"], ShouldEqual, "123")
		c.So(r["int"], ShouldEqual, "1")
		c.So(r["float"], ShouldEqual, "3.75")
		c.So(r["bool"], ShouldEqual, "true")
	})

	Convey("list numbering should be handled correctly", t, func(c C) {
		r := FlattenFromJSON(`
{ "list": [
	"first",
	"2nd",
	3,
	4.3874
]
}
		`)

		c.So(len(r), ShouldEqual, 4)
		c.So(r["list.0"], ShouldEqual, "first")
		c.So(r["list.1"], ShouldEqual, "2nd")
		c.So(r["list.2"], ShouldEqual, "3")
		c.So(r["list.3"], ShouldEqual, "4.3874")

		r = FlattenFromJSON(`
		{ "list": [
		]
		}
		`)

		c.So(len(r), ShouldEqual, 0)

	})

	Convey("nested maps should be handled correctly", t, func(c C) {
		r := FlattenFromJSON(`
{   
	"level1": {
		"level2": {
			"here": "too",
			"level3": {
				"a": 1,
				"b": false,
				"level4": {

				}
			}
		},
		"a": "others"
	}
}
		`)

		c.So(len(r), ShouldEqual, 4)
		c.So(r["level1.a"], ShouldEqual, "others")
		c.So(r["level1.level2.level3.a"], ShouldEqual, "1")
		c.So(r["level1.level2.here"], ShouldEqual, "too")
		c.So(r["level1.level2.level3.b"], ShouldEqual, "false")

		r = FlattenFromJSON(`
		{ "empty_map": {}
		}
		`)

		c.So(len(r), ShouldEqual, 0)

	})

	Convey("mixed example should be handled correctly", t, func(c C) {
		r := FlattenFromJSON(`
{   
	"level1": {
		"level2": {
			"here": "too",
			"level3": [
				"a",
				false,
				{
					"other": "one"
				}
			]
		},
		"a": "others"
	}
}
		`)

		c.So(len(r), ShouldEqual, 5)
		c.So(r["level1.a"], ShouldEqual, "others")
		c.So(r["level1.level2.here"], ShouldEqual, "too")
		c.So(r["level1.level2.level3.0"], ShouldEqual, "a")
		c.So(r["level1.level2.level3.1"], ShouldEqual, "false")
		c.So(r["level1.level2.level3.2.other"], ShouldEqual, "one")

		r = FlattenFromJSON(`
		{ "empty_map": {}
		}
		`)

		c.So(len(r), ShouldEqual, 0)

	})
}
