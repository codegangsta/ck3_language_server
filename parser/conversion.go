package parser

func convertPos(p position) Pos {
	return Pos{p.line, p.col, p.offset}
}

func convertPairs(i interface{}) []Pair {
	is := i.([]interface{})
	pairs := make([]Pair, len(is))
	for idx := range is {
		pairs[idx] = is[idx].(Pair)
	}

	return pairs
}
