package parser

// Node is an interface that every element of the AST needs to implement.
type Node interface {
	Range() Range
}

type Script struct {
	Attributes []Attribute
	Range
}

type Block struct {
	Attributes []Attribute
	Range
}

type Attribute struct {
	ID       ID
	Value    Node
	SrcRange Range
}

type ID struct {
	Name     string
	SrcRange Range
}

type Lookup struct {
	Value    string
	SrcRange Range
}

type Macro struct {
	Value    string
	SrcRange Range
}

type Boolean struct {
	Value    bool
	SrcRange Range
}

type Number struct {
	Value    bool
	SrcRange Range
}
