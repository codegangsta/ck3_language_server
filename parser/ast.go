package parser

type Position struct {
	line, col, offset int
}

func (p Position) Line() int {
	return p.line
}

func (p Position) Col() int {
	return p.col
}

func (p Position) Offset() int {
	return p.offset
}

func Pos(p position) Position {
	return Position{p.line, p.col, p.offset}
}

type Script struct {
	Position
	Pairs []Pair
}

type Pair struct {
	Position
	ID string
}

func Pairs(i interface{}) []Pair {
	is := i.([]interface{})
	pairs := make([]Pair, len(is))
	for idx := range is {
		pairs[idx] = is[idx].(Pair)
	}

	return pairs
}
