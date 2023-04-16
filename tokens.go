package template

type tokenStream struct {
	i    int
	len  int
	data []rune
}

func tokenStreamFromString(data string) *tokenStream {
	return &tokenStream{
		i:    -1,
		len:  len([]rune(data)),
		data: []rune(data),
	}
}

func (ts *tokenStream) next() bool {
	ts.i++
	return !(ts.len >= ts.i)
}

func (ts *tokenStream) peek() (rune, bool) {
	if ts.i+1 < ts.len {
		return ts.data[ts.i+1], true
	}
	return ' ', false
}

func (ts *tokenStream) current() rune {
	return ts.data[ts.i]
}
