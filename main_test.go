package main

import (
	"reflect"
	"testing"
)

type testpsSplitter struct {
	values  string
	average []string
}

type testToPig struct {
	values []string
	average string
}

var testsSplit = []testpsSplitter{
	{"hello world", []string{"hello", "world"}},
}

var testsPig = []testToPig{
	{[]string{"hello", "world"}, "ellohay orldway"},
	{[]string{"hello,", "world"}, "ellohay, orldway"},
	{[]string{"hello,", "world?"}, "ellohay, orldway?"},
	{[]string{"hello", "world!"}, "ellohay orldway!"},
	{[]string{"I"}, "Iyay"},
}

func TestSplitter(t *testing.T) {
	for _, pair := range testsSplit {
		v := Splitter(pair.values)
		if !reflect.DeepEqual(v, pair.average) {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestToPig(t *testing.T) {
	for _, pair := range testsPig {
		v := ToPig(pair.values)
		if v != pair.average {
            t.Error(
                "For", pair.values, 
                "expected", pair.average,
                "got", v,
            )
        }
	}
}
