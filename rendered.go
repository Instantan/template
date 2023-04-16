package template

import "strings"

type RenderedNodes []RenderedNode

type RenderedNode struct {
	Tag        string
	Attributes []string
	Children   RenderedNodes
	Data       string
}

func (n RenderedNodes) Equals(nodes RenderedNodes) bool {
	if len(n) != len(nodes) {
		return false
	}
	for i := range n {
		if !n[i].Equals(nodes[i]) {
			return false
		}
	}
	return true
}

func (n RenderedNodes) String() string {
	builder := &strings.Builder{}
	n.buildString(builder)
	return builder.String()
}

func (n RenderedNodes) buildString(builder *strings.Builder) {
	for i := range n {
		n[i].buildString(builder)
	}
}

func (n RenderedNode) Equals(node RenderedNode) bool {
	return n.Tag == node.Tag &&
		n.Data == node.Data &&
		attributesAreEqual(n.Attributes, node.Attributes) &&
		n.Children.Equals(node.Children)
}

func (n RenderedNode) buildString(builder *strings.Builder) {
	builder.WriteString("<" + n.Tag)
	for i := range n.Attributes {
		builder.WriteString(" " + n.Attributes[i])
	}
	if n.Data != "" {
		builder.WriteString(">" + n.Data)
		builder.WriteString("</" + n.Tag + ">")
		return
	}
	if len(n.Children) == 0 {
		builder.WriteString("/>")
		return
	}
	builder.WriteString(">")
	n.Children.buildString(builder)
	builder.WriteString("</" + n.Tag + ">")
}
