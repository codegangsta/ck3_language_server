package parser

import (
	"reflect"
	"testing"
)

func TestBasicParsing(t *testing.T) {
	_, err := ParseFile("test_cases/example.txt")

	if err != nil {
		t.Fatal(err)
	}
}

func TestValues(t *testing.T) {
	// todo: switch comparison to full struct value
	values := map[string]string{
		"foo = 1": "foo:3 1:5",
	}

	for in, out := range values {
		script, err := Parse("f", []byte(in))
		if err != nil {
			t.Errorf("Unable to process '%s': %v", in, err)
		}

		pair := script.(Script).Pairs[0]
		expect(t, pair.String(), out)
	}
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %+v (type %v) - Got %+v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
