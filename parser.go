package template

type parser struct {
	n        Nodes
	currNode Node
	ts       *tokenStream
}

func (p *parser) parse(template string) {
	p.ts = tokenStreamFromString(template)
	for p.ts.next() {
		curr := p.ts.current()
		switch curr {
		case '<':
			p.parseTag()
		}
	}
}

func (p *parser) parseTag() {
	parsingAttributes := false

	for p.ts.next() {
		curr := p.ts.current()
		switch curr {
		case ' ':
			if p.currNode.Tag == "" || parsingAttributes {
				continue
			}
			parsingAttributes = true
		case '>':
			p.parseChilds()
		default:
			p.currNode.Tag += string(curr)
		}
	}
}

func (p *parser) parseChilds() {

}

func (p *parser) nodes() Nodes {
	return p.n
}
