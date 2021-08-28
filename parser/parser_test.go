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
				Name:     "foo",
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
				Name:     "pi",
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
				Name:     "negative_number",
				Position: &Pos{1, 1, 0, 15},
			},
			Value: &Number{
				Value:    -18,
				Position: &Pos{1, 19, 18, 3},
			},
			Position: &Pos{1, 1, 0, 21},
		},

		// TODO boolean yes
		// TODO boolean no
		// TODO macro
		// TODO scopes
		// TODO dot notation
		// TODO ids
		// TODO blocks
		// TODO string literals
		// TODO comments
	}

	for in, want := range values {
		script, err := Parse("f", []byte(in))
		if err != nil {
			t.Errorf("Unable to process '%s': %v", in, err)
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
