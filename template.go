package template

import (
	"io"
	"sync/atomic"

	"github.com/valyala/fasttemplate"
)

var templateID = atomic.Int64{}

type Template struct {
	id         int
	tmplStr    string
	tmpl       *fasttemplate.Template
	isDynamic  bool
	tagToIndex map[string]int
	indexToTag map[string]string
}

func Parse(template string) *Template {
	templateID.Add(1)
	template = prepareDynamicValues(template)
	template = minifyTemplate(template)
	template, tags := replaceTagsWithIndexes(template)
	return &Template{
		id:         int(templateID.Load()),
		tmplStr:    template,
		tmpl:       fasttemplate.New(template, "~", "~"),
		isDynamic:  len(tags) > 0,
		tagToIndex: tags,
		indexToTag: convTagToIndexToIndexToTag(tags),
	}
}

func (t *Template) String() string {
	return t.tmplStr
}

func (t *Template) HTML() string {
	return ""
}

func (t *Template) ID() int {
	return t.id
}

func (t *Template) IsDynamic() bool {
	return t.isDynamic
}

func (t *Template) DynamicValues() map[string]int {
	return t.tagToIndex
}

func (t *Template) Execute(w io.Writer, fn func(w io.Writer, tag string) (int, error)) (int64, error) {
	return t.tmpl.ExecuteFunc(w, func(w io.Writer, index string) (int, error) {
		return fn(w, t.indexToTag[index])
	})
}

func (t *Template) ExecuteString(fn func(w io.Writer, tag string) (int, error)) string {
	return t.tmpl.ExecuteFuncString(func(w io.Writer, index string) (int, error) {
		return fn(w, t.indexToTag[index])
	})
}
