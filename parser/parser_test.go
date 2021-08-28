package parser

import (
	"math/big"
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
	// todo: switch comparison to full struct value
	values := map[string]*Command{
		"foo = 1": {
			ID: &ID{
				Name:     "foo",
				Position: &Pos{1, 1, 0, 3},
			},
			Value: &Number{
				Value:    big.NewFloat(1.0),
				Position: &Pos{1, 7, 6, 1},
			},
			Position: &Pos{1, 1, 0, 7},
		},
	}

	for in, want := range values {
		script, err := Parse("f", []byte(in))
		if err != nil {
			t.Errorf("Unable to process '%s': %v", in, err)
		}

		got := script.(*Script).Commands[0]
		mustEqual(t, got, want)
	}
}

func mustEqual(t *testing.T, a interface{}, b interface{}) {
	if diff := deep.Equal(a, b); diff != nil {
		t.Error(strings.Join(diff, ", "))
	}
}
