package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestReverse struct {
	name   string
	input  string
	ouput  string
	status string
}

func TestGetReverseRuneSuccess(t *testing.T) {
	var testList = []TestReverse{
		{name: "Test positive case 1", input: "golang", ouput: "gnalog", status: "success"},
		{name: "Test positive case 2", input: "are", ouput: "era", status: "success"},
		{name: "Test negative case 1", input: "lol", ouput: "lol", status: "failed"},
		{name: "Test negative case 1", input: "git", ouput: "tig", status: "failed"},
	}
	for _, tc := range testList {
		t.Run(tc.name, func(t *testing.T) {
			outstr, _ := getReverseRune(tc.input)
			assert.Equal(t, outstr, tc.ouput)
		})
	}
}

func TestGetReverseRuneFailed(t *testing.T) {
	var testList = []TestReverse{
		{name: "Test positive case 1", input: "golang", ouput: "golang", status: "success"},
		{name: "Test positive case 2", input: "are", ouput: "are", status: "success"},
		{name: "Test negative case 1", input: "golang", ouput: "golang", status: "success"},
	}
	for _, tc := range testList {
		t.Run(tc.name, func(t *testing.T) {
			outstr, _ := getReverseRune(tc.input)
			assert.Assert(t, outstr != tc.ouput)
		})
	}
}

func TestGetReverseSuccess(t *testing.T) {
	var testList = []TestReverse{
		{name: "Test positive case 1", input: "golang", ouput: "gnalog", status: "success"},
		{name: "Test positive case 2", input: "are", ouput: "era", status: "success"},
		{name: "Test negative case 1", input: "lol", ouput: "lol", status: "failed"},
		{name: "Test negative case 1", input: "git", ouput: "tig", status: "failed"},
	}
	for _, tc := range testList {
		t.Run(tc.name, func(t *testing.T) {
			outstr := getReverse(tc.input)
			assert.Equal(t, outstr, tc.ouput)
		})
	}
}

func TestGetReverseFailed(t *testing.T) {
	var testList = []TestReverse{
		{name: "Test positive case 1", input: "golang", ouput: "golang", status: "success"},
		{name: "Test positive case 2", input: "are", ouput: "are", status: "success"},
		{name: "Test negative case 1", input: "golang", ouput: "golang", status: "success"},
	}
	for _, tc := range testList {
		t.Run(tc.name, func(t *testing.T) {
			outstr := getReverse(tc.input)
			assert.Assert(t, outstr != tc.ouput)
		})
	}
}

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, true, true)
}
