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
	Convey("empty input should result in empty map", t, func() {
		em := make(map[string]interface{})
		r := Flatten(em)

		So(len(r), ShouldEqual, 0)
	})

	Convey("basic types should be handled correctly", t, func() {
		r := FlattenFromJSON(`
{ "string": "123",
"int": 1,
"float": 3.75,
"bool": true
}
		`)

		So(len(r), ShouldEqual, 4)
		So(r["string"], ShouldEqual, "123")
		So(r["int"], ShouldEqual, "1")
		So(r["float"], ShouldEqual, "3.75")
		So(r["bool"], ShouldEqual, "true")
	})

	Convey("list numbering should be handled correctly", t, func() {
		r := FlattenFromJSON(`
{ "list": [
	"first",
	"2nd",
	3,
	4.3874
]
}
		`)

		So(len(r), ShouldEqual, 4)
		So(r["list.0"], ShouldEqual, "first")
		So(r["list.1"], ShouldEqual, "2nd")
		So(r["list.2"], ShouldEqual, "3")
		So(r["list.3"], ShouldEqual, "4.3874")

		r = FlattenFromJSON(`
		{ "list": [
		]
		}
		`)

		So(len(r), ShouldEqual, 0)

	})

	Convey("nested maps should be handled correctly", t, func() {
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

		So(len(r), ShouldEqual, 4)
		So(r["level1.a"], ShouldEqual, "others")
		So(r["level1.level2.level3.a"], ShouldEqual, "1")
		So(r["level1.level2.here"], ShouldEqual, "too")
		So(r["level1.level2.level3.b"], ShouldEqual, "false")

		r = FlattenFromJSON(`
		{ "empty_map": {}
		}
		`)

		So(len(r), ShouldEqual, 0)

	})

	Convey("mixed example should be handled correctly", t, func() {
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

		So(len(r), ShouldEqual, 5)
		So(r["level1.a"], ShouldEqual, "others")
		So(r["level1.level2.here"], ShouldEqual, "too")
		So(r["level1.level2.level3.0"], ShouldEqual, "a")
		So(r["level1.level2.level3.1"], ShouldEqual, "false")
		So(r["level1.level2.level3.2.other"], ShouldEqual, "one")

		r = FlattenFromJSON(`
		{ "empty_map": {}
		}
		`)

		So(len(r), ShouldEqual, 0)

	})
}
