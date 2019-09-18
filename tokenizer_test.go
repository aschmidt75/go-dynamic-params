package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTokenizerValidInputs(t *testing.T) {
	Convey("Empty string should return empty token list", t, func() {
		t := NewTokenizerFromString("")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 0)

		t = NewTokenizer([]byte{})
		tokens, err = t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 0)
	})

	Convey("Non-Param string should return a single param token", t, func() {
		t := NewTokenizerFromString("just static stuff")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 1)
		So(tokens[0].part, ShouldResemble, []byte("just static stuff"))
		So(tokens[0].tkType, ShouldEqual, typeStaticPart)

		b := []byte{'J', 'u', 's', 't', ' ', 'b', 'y', 't', 'e', 's'}
		t = NewTokenizer(b)
		tokens, err = t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 1)
		So(tokens[0].part, ShouldResemble, b)
		So(tokens[0].tkType, ShouldEqual, typeStaticPart)
	})

	Convey("Non-Param string with param markup should return a single param token", t, func() {
		t := NewTokenizerFromString("{just} $static {} stuff}")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 1)
		So(string(tokens[0].part), ShouldEqual, "{just} $static {} stuff}")
		So(tokens[0].tkType, ShouldEqual, typeStaticPart)
	})

	Convey("Param-only string should return a single token", t, func() {
		t := NewTokenizerFromString("${single.param}")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 1)
		So(string(tokens[0].part), ShouldEqual, "single.param")
		So(tokens[0].tkType, ShouldEqual, typeParamPart)
	})

	Convey("Static/Param string should return two tokens", t, func() {
		t := NewTokenizerFromString("Just a static part and a ${single.param}")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 2)
		So(string(tokens[0].part), ShouldEqual, "Just a static part and a ")
		So(tokens[0].tkType, ShouldEqual, typeStaticPart)
		So(string(tokens[1].part), ShouldEqual, "single.param")
		So(tokens[1].tkType, ShouldEqual, typeParamPart)
	})

	Convey("Param/Static string should return two tokens", t, func() {
		t := NewTokenizerFromString("${single.param} followed by static")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 2)
		So(string(tokens[1].part), ShouldEqual, " followed by static")
		So(tokens[1].tkType, ShouldEqual, typeStaticPart)
		So(string(tokens[0].part), ShouldEqual, "single.param")
		So(tokens[0].tkType, ShouldEqual, typeParamPart)
	})

	Convey("Param/Static/Param string should return three tokens", t, func() {
		t := NewTokenizerFromString("${single.param} followed by static and ${another.param}")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 3)
		So(string(tokens[1].part), ShouldEqual, " followed by static and ")
		So(tokens[1].tkType, ShouldEqual, typeStaticPart)
		So(string(tokens[0].part), ShouldEqual, "single.param")
		So(tokens[0].tkType, ShouldEqual, typeParamPart)
		So(string(tokens[2].part), ShouldEqual, "another.param")
		So(tokens[2].tkType, ShouldEqual, typeParamPart)
	})

	Convey("Static/Param/Static string should return three tokens", t, func() {
		t := NewTokenizerFromString("just a ${single} param")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 3)
		So(string(tokens[0].part), ShouldEqual, "just a ")
		So(tokens[0].tkType, ShouldEqual, typeStaticPart)
		So(string(tokens[1].part), ShouldEqual, "single")
		So(tokens[1].tkType, ShouldEqual, typeParamPart)
		So(string(tokens[2].part), ShouldEqual, " param")
		So(tokens[2].tkType, ShouldEqual, typeStaticPart)
	})

	Convey("A single nested Param should be tokenized correctly", t, func() {
		t := NewTokenizerFromString("${this is ${a} param}")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 1)
		So(string(tokens[0].part), ShouldEqual, "this is ${a} param")
		So(tokens[0].tkType, ShouldEqual, typeParamPart)
		So(tokens[0].withNestedParam, ShouldBeTrue)

		t = NewTokenizerFromString("${this is ${a${nothe${r}}} param}")
		tokens, err = t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 1)
		So(string(tokens[0].part), ShouldEqual, "this is ${a${nothe${r}}} param")
		So(tokens[0].tkType, ShouldEqual, typeParamPart)
		So(tokens[0].withNestedParam, ShouldBeTrue)
	})

	Convey("Imbalanced params look strange but are ok", t, func() {
		t := NewTokenizerFromString("this is ${strange $but}} valid")
		tokens, err := t.Tokenize()

		So(err, ShouldBeNil)
		So(len(tokens), ShouldEqual, 3)
		So(string(tokens[0].part), ShouldEqual, "this is ")
		So(tokens[0].tkType, ShouldEqual, typeStaticPart)
		So(string(tokens[1].part), ShouldEqual, "strange $but")
		So(tokens[1].tkType, ShouldEqual, typeParamPart)
		So(string(tokens[2].part), ShouldEqual, "} valid")
		So(tokens[2].tkType, ShouldEqual, typeStaticPart)

	})

}

func TestTokenizerInvalidInputs(t *testing.T) {
	Convey("Empty params should return an error", t, func() {
		t := NewTokenizerFromString("this is empty: ${}")
		tokens, err := t.Tokenize()

		So(err, ShouldNotBeNil)
		So(len(tokens), ShouldEqual, 1)
	})

	Convey("Unclosed params should return an error", t, func() {
		t := NewTokenizerFromString("this is ${not valid")
		_, err := t.Tokenize()

		So(err, ShouldNotBeNil)
	})

	Convey("Unclosed params should return an error", t, func() {
		t := NewTokenizerFromString("this is ${also ${not} valid")
		_, err := t.Tokenize()

		So(err, ShouldNotBeNil)
	})

}
