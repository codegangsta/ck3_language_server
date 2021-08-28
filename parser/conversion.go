package parser

import (
	"bytes"
	"errors"
	"strconv"
)

func convertPos(c *current) *Pos {
	return &Pos{
		Line:   c.pos.line,
		Column: c.pos.col,
		Offset: c.pos.offset,
		Length: len(c.text),
	}
}

func convertMacro(c *current) (*Macro, error) {
	return &Macro{
		Value:    string(c.text),
		Position: convertPos(c),
	}, nil
}

func convertBoolean(c *current) (*Boolean, error) {
	var val bool

	switch string(c.text) {

	case "yes":
		val = true
	case "no":
		val = false
	default:
		return nil, errors.New("Invalid boolean value")
	}

	return &Boolean{
		Value:    val,
		Position: convertPos(c),
	}, nil
}

func convertNumber(c *current) (*Number, error) {
	f, err := strconv.ParseFloat(string(c.text), 64)
	if err != nil {
		return nil, err
	}

	return &Number{
		Value:    f,
		Position: convertPos(c),
	}, nil
}

func convertID(c *current) (*ID, error) {
	return &ID{
		Value:    string(c.text),
		Position: convertPos(c),
	}, nil
}

func convertString(c *current) (*String, error) {
	c.text = bytes.Replace(c.text, []byte(`\/`), []byte(`/`), -1)

	val, err := strconv.Unquote(string(c.text))
	if err != nil {
		return nil, err
	}
	return &String{
		Value:    val,
		Position: convertPos(c),
	}, nil
}

func convertCommand(c *current, id interface{}, value interface{}) (*Command, error) {
	return &Command{
		ID:       id.(*ID),
		Value:    value.(Node),
		Position: convertPos(c),
	}, nil
}

func convertScript(c *current, commands interface{}) (*Script, error) {
	return &Script{
		Commands: convertCommands(commands),
		Position: convertPos(c),
	}, nil
}

func convertBlock(c *current, commands interface{}) (*Block, error) {
	return &Block{
		Commands: convertCommands(commands),
		Position: convertPos(c),
	}, nil
}

func convertCommands(i interface{}) []*Command {
	is := i.([]interface{})
	commands := make([]*Command, len(is))
	for idx := range is {
		commands[idx] = is[idx].(*Command)
	}

	return commands
}
