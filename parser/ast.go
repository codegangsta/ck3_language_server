package parser

import "math/big"

// Node is an interface that every element of the AST needs to implement in
// order to support proper parsing
type Node interface {
	// Range returns the reletive range of the given node
	Range() Range
}

// Script represents an entire script file
type Script struct {
	Commands []Command
	SrcRange Range
}

// A Block is a value object that represents a nested block in the script structure:
//
// ex: { always = yes }
// or: { foo = { bar = bat } }
// Blocks contain 0 or more Commands that will be used to construct the block.
type Block struct {
	Commands []Command
	SrcRange Range
}

// Command represents a command that is called with a particular value.
// ID represents the name of the command, while Value contains the left hand
// side value og the command (ID and Value are separated by a '=')
//
// ex: always = yes
type Command struct {
	ID       ID
	Value    Node
	SrcRange Range
}

// ID represents an identifier value. Typically used for strings and command
// IDs
type ID struct {
	Name     string
	SrcRange Range
}

// Macro represents a script macro value
// ex: $MY_MACRO$
type Macro struct {
	Value    string
	SrcRange Range
}

// Boolean represents a boolean value
type Boolean struct {
	Value    bool
	SrcRange Range
}

// Number represents a Number value
type Number struct {
	Value    *big.Float
	SrcRange Range
}
