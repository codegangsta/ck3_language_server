package parser

func convertPos(p position) *Pos {
	return &Pos{p.line, p.col, p.offset}
}

func convertMacro(c *current) (*Macro, error) {
	return &Macro{}, nil
}

func convertBoolean(c *current) (*Boolean, error) {
	return &Boolean{}, nil
}

func convertNumber(c *current) (*Number, error) {
	return &Number{}, nil
}

func convertID(c *current) (*ID, error) {
	return &ID{}, nil
}

func convertCommand(c *current, id interface{}, value interface{}) (*Command, error) {
	return &Command{}, nil
}

func convertScript(c *current, commands interface{}) (*Script, error) {
	return &Script{}, nil
}

func convertBlock(c *current, commands interface{}) (*Block, error) {
	return &Block{}, nil
}

func convertCommands(i interface{}) []Command {
	is := i.([]interface{})
	pairs := make([]Command, len(is))
	for idx := range is {
		pairs[idx] = is[idx].(Command)
	}

	return pairs
}
