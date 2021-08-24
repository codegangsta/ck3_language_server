package parser

import "testing"

func TestBasicParsing(t *testing.T) {
	_, err := ParseFile("example.txt")

	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkBasicParsing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFile("example.test")
	}
}
