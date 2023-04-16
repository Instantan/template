package template

type Value any

type Nodes []Node

type Node struct {
	Tag        string
	Attributes []Attribute
	Children   []Node
	Scope      string
	Data       Value
}

type Attribute struct {
	Key   string
	Value Value
}

func Parse(template string) Nodes {
	p := &parser{}
	p.parse(template)
	return p.nodes()
}

func (n Nodes) Apply(data any) RenderedNodes {
	return []RenderedNode{}
}

func (n Node) HasChildren() bool {
	return len(n.Children) != 0
}
