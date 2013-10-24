package main

import "testing"

type testpair struct {
	inputString string
	splitString []string
}

var tests = []testpair{
	{"X", []string{"X"}},
	{"X-X", []string{"X", "X"}},
}

func TestSplitInputString(t *testing.T) {
	for _, pair := range tests {
		v := SplitInputString(pair.inputString)

		if len(pair.splitString) != len(v) {
			t.Error(
				"For", pair.inputString,
				"expected", pair.splitString,
				"got", v,
			)
		}
	}
}

type convertTestPair struct {
	input  string
	output int
}

var convertTests = []convertTestPair{
	{"1", 1},
	{"10", 10},
	{"X", 10},
	{"/", 5},
}

func TestConvertStringToValue(t *testing.T) {
	for _, pair := range convertTests {
		v := ConvertStringToValue(pair.input)

		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}

type basicScoreTestPair struct {
	input  string
	output int
}

var basicScoreTests = []basicScoreTestPair{
	{"1-1", 2},
	{"10-10-10-2", 32},
	{"X-0-0", 10},
	{"X-5-2", 24},
	{"5-/-5-0-0-0-0-0", 20},
	{"9-0-9-0-9-0-9-0-9-0-9-0-9-0-9-0-9-0-9-0", 90},
	{"5-/-5-/-5-/-5-/-5-/-5-/-5-/-5-/-5-/-5-/-5", 150},
	{"X-X-X-X-X-X-X-X-X-X-X-X", 300},
}

func TestBasicScore(t *testing.T) {
	for _, pair := range basicScoreTests {
		v := BasicScore(pair.input)

		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}
