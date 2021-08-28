//TODO: Write some methods for ranging over nodes here
package parser

type Pos struct {
	Line, Column, Offset int
}

type Range struct {
	Start, End Pos
}
