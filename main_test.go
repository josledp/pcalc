package main

import "testing"

func TestIsFloat(t *testing.T) {
	ttable := []struct {
		entry    string
		expected bool
	}{
		{
			entry:    "3.14",
			expected: true,
		},
		{
			entry:    "3.14",
			expected: true,
		},
		{
			entry:    "3,14",
			expected: false,
		},
		{
			entry:    "314e10",
			expected: false,
		},
	}
	for _, tc := range ttable {
		t.Run(tc.entry, func(t *testing.T) {
			if isFloat(tc.entry) != tc.expected {
				t.Errorf("expected '%v' got '%v' for '%s'", tc.expected, isFloat(tc.entry), tc.entry)
			}
		})
	}
}

func TestParse(t *testing.T) {
	p = new(pile)
	tseq := []struct {
		entry     string
		expString string
		expError  bool
	}{
		{entry: "3", expString: "", expError: false},
		{entry: "p", expString: "3", expError: false},
		{entry: "n", expString: "3", expError: false},
		{entry: "2", expString: "", expError: false},
		{entry: "3", expString: "", expError: false},
		{entry: "*", expString: "", expError: false},
		{entry: "p", expString: "6", expError: false},
		{entry: "2", expString: "", expError: false},
		{entry: "-", expString: "", expError: false},
		{entry: "p", expString: "4", expError: false},
		{entry: "v", expString: "", expError: false},
		{entry: "p", expString: "2", expError: false},
		{entry: "2", expString: "", expError: false},
		{entry: "^", expString: "", expError: false},
		{entry: "p", expString: "4", expError: false},
		{entry: "2", expString: "", expError: false},
		{entry: "/", expString: "", expError: false},
		{entry: "p", expString: "2", expError: false},
		{entry: "x", expString: "", expError: true},
	}
	for _, tc := range tseq {
		str, err := parse(tc.entry)
		if str != tc.expString {
			t.Errorf("expected '%s' got '%s'", tc.expString, str)
		}
		if err != nil && !tc.expError {
			t.Errorf("not error expected but got: %v", err)
		} else if err == nil && tc.expError {
			t.Errorf("error expected, but got none")
		}

	}
}
