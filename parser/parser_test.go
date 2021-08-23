package parser

import "testing"

func TestBasicParsing(t *testing.T) {
	_, err := ParseFile("example.txt")

	if err != nil {
		t.Fatal(err)
	}
}
