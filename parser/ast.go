package parser

import "math/big"

// Node is an interface that every element of the AST needs to implement in
// order to support proper parsing
type Node interface {
	// Pos returns the reletive position of the given node
	Pos() *Pos
}

// Script represents an entire script file
type Script struct {
	Commands []*Command
	Position *Pos
}

// A Block is a value object that represents a nested block in the script structure:
//
// ex: { always = yes }
// or: { foo = { bar = bat } }
// Blocks contain 0 or more Commands that will be used to construct the block.
type Block struct {
	Commands []*Command
	Position *Pos
}

// Command represents a command that is called with a particular value.
// ID represents the name of the command, while Value contains the left hand
// side value og the command (ID and Value are separated by a '=')
//
// ex: always = yes
type Command struct {
	ID       *ID
	Value    Node
	Position *Pos
}

// ID represents an identifier value. Typically used for strings and command
// IDs
type ID struct {
	Name     string
	Position *Pos
}

// Macro represents a script macro value
// ex: $MY_MACRO$
type Macro struct {
	Value    string
	Position *Pos
}

// Boolean represents a boolean value
type Boolean struct {
	Value    bool
	Position *Pos
}

// Number represents a Number value
type Number struct {
	Value    *big.Float
	Position *Pos
}

// Position Methods for all nodes
func (x *Script) Pos() *Pos  { return x.Position }
func (x *Block) Pos() *Pos   { return x.Position }
func (x *Command) Pos() *Pos { return x.Position }
func (x *ID) Pos() *Pos      { return x.Position }
func (x *Macro) Pos() *Pos   { return x.Position }
func (x *Boolean) Pos() *Pos { return x.Position }
func (x *Number) Pos() *Pos  { return x.Position }
