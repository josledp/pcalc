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
			expected: true,
		},
		{
			entry:    "-314.10e10",
			expected: true,
		},
		{
			entry:    "31.4e-10",
			expected: true,
		},
		{
			entry:    "-314.22e-10",
			expected: true,
		},
		{
			entry:    "-",
			expected: false,
		},
		{
			entry:    "-e",
			expected: false,
		},
		{
			entry:    "e10",
			expected: false,
		},
		{
			entry:    "e",
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
		str, err := parseEntry(tc.entry)
		if str != tc.expString {
			t.Errorf("expected '%s' got '%s' for %s", tc.expString, str, tc.entry)
		}
		if err != nil && !tc.expError {
			t.Errorf("not error expected but got: %v for %s", err, tc.entry)
		} else if err == nil && tc.expError {
			t.Errorf("error expected, but got none for %s", tc.entry)
		}

	}
}
