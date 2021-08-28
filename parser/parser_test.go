package parser

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestBasicParsing(t *testing.T) {
	_, err := ParseFile("test_cases/example.txt")

	if err != nil {
		t.Fatal(err)
	}
}

func TestValues(t *testing.T) {
	values := map[string]*Command{
		// Integers
		"foo = 1": {
			ID: &ID{
				Value:    "foo",
				Position: &Pos{1, 1, 0, 3},
			},
			Value: &Number{
				Value:    1.0,
				Position: &Pos{1, 7, 6, 1},
			},
			Position: &Pos{1, 1, 0, 7},
		},

		// Floats
		"pi = 3.14": {
			ID: &ID{
				Value:    "pi",
				Position: &Pos{1, 1, 0, 2},
			},
			Value: &Number{
				Value:    3.14,
				Position: &Pos{1, 6, 5, 4},
			},
			Position: &Pos{1, 1, 0, 9},
		},

		// Negative numbers
		"negative_number = -18": {
			ID: &ID{
				Value:    "negative_number",
				Position: &Pos{1, 1, 0, 15},
			},
			Value: &Number{
				Value:    -18,
				Position: &Pos{1, 19, 18, 3},
			},
			Position: &Pos{1, 1, 0, 21},
		},

		// boolean yes
		"always = yes": {
			ID: &ID{
				Value:    "always",
				Position: &Pos{1, 1, 0, 6},
			},
			Value: &Boolean{
				Value:    true,
				Position: &Pos{1, 10, 9, 3},
			},
			Position: &Pos{1, 1, 0, 12},
		},

		// boolean no
		"always = no": {
			ID: &ID{
				Value:    "always",
				Position: &Pos{1, 1, 0, 6},
			},
			Value: &Boolean{
				Value:    false,
				Position: &Pos{1, 10, 9, 2},
			},
			Position: &Pos{1, 1, 0, 11},
		},

		// macro
		"macro = $MY_MACRO$": {
			ID: &ID{
				Value:    "macro",
				Position: &Pos{1, 1, 0, 5},
			},
			Value: &Macro{
				Value:    "$MY_MACRO$",
				Position: &Pos{1, 9, 8, 10},
			},
			Position: &Pos{1, 1, 0, 18},
		},

		//  scopes
		"scope = var:magic_power": {
			ID: &ID{
				Value:    "scope",
				Position: &Pos{1, 1, 0, 5},
			},
			Value: &ID{
				Value:    "var:magic_power",
				Position: &Pos{1, 9, 8, 15},
			},
			Position: &Pos{1, 1, 0, 23},
		},

		// dot notation
		"nested.id.001 = 1": {
			ID: &ID{
				Value:    "nested.id.001",
				Position: &Pos{1, 1, 0, 13},
			},
			Value: &Number{
				Value:    1.0,
				Position: &Pos{1, 17, 16, 1},
			},
			Position: &Pos{1, 1, 0, 17},
		},

		// blocks
		"nested.id.001 = { always = no }": {
			ID: &ID{
				Value:    "nested.id.001",
				Position: &Pos{1, 1, 0, 13},
			},
			Value: &Block{
				Commands: []*Command{
					{
						ID: &ID{
							Value:    "always",
							Position: &Pos{1, 19, 18, 6},
						},
						Value: &Boolean{
							Value:    false,
							Position: &Pos{1, 28, 27, 2},
						},
						Position: &Pos{1, 18, 17, 13},
					},
				},
				Position: &Pos{1, 17, 16, 15},
			},
			Position: &Pos{1, 1, 0, 31},
		},

		// TODO string literals

		// comments
		"nested.id.001 = 1 # with a comment": {
			ID: &ID{
				Value:    "nested.id.001",
				Position: &Pos{1, 1, 0, 13},
			},
			Value: &Number{
				Value:    1.0,
				Position: &Pos{1, 17, 16, 1},
			},
			Position: &Pos{1, 1, 0, 34},
		},
	}

	for in, want := range values {
		script, err := Parse(in, []byte(in))
		if err != nil {
			t.Error(err)
		}

		got := script.(*Script).Commands[0]
		if diff := deep.Equal(got, want); diff != nil {
			t.Errorf("'%s': \n%v", in, strings.Join(diff, "\n"))
		}
	}
}

func mustEqual(t *testing.T, a interface{}, b interface{}) {
	if diff := deep.Equal(a, b); diff != nil {
		t.Error(strings.Join(diff, ", "))
	}
}
