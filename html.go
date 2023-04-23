package template

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	mhtml "github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
	"github.com/valyala/fasttemplate"
)

var minifier *minify.M

func init() {
	minifier = minify.New()
	minifier.AddFunc("text/css", css.Minify)
	minifier.AddFunc("text/html", mhtml.Minify)
	minifier.AddFunc("image/svg+xml", svg.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
}

func prepareDynamicValues(template string) string {
	re := regexp.MustCompile(`((?m){{.*}})`)
	template = re.ReplaceAllStringFunc(template, func(s string) string {
		s = strings.ReplaceAll(s, "{{", "~")
		s = strings.ReplaceAll(s, "}}", "~")
		s = strings.ReplaceAll(s, " ", "")
		return s
	})
	return template
}

func minifyTemplate(template string) string {
	template, err := minifier.String("text/html", template)
	if err != nil {
		panic(err)
	}
	return template
}

func replaceTagsWithIndexes(template string) (string, map[string]int) {
	index := 0
	m := map[string]int{}
	template = fasttemplate.ExecuteFuncString(template, "~", "~", func(w io.Writer, tag string) (int, error) {
		i, ok := m[tag]
		if ok {
			return io.WriteString(w, "~"+strconv.Itoa(i)+"~")
		} else {
			m[tag] = index
			i, err := io.WriteString(w, "~"+strconv.Itoa(index)+"~")
			index++
			return i, err
		}
	})
	return template, m
}

func convTagToIndexToIndexToTag(m map[string]int) map[string]string {
	itt := map[string]string{}
	for tag, index := range m {
		itt[strconv.Itoa(index)] = tag
	}
	return itt
}
