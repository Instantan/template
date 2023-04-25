package template_test

import (
	"io"
	"testing"

	"github.com/Instantan/template"
)

func TestParse(t *testing.T) {
	{
		tmpl := template.Parse(`
			<div>
				<h1 class="{{bla}}">
					{{name}}
				</h1>
			</div>
		`)
		t.Log(tmpl.String())
		t.Log(tmpl.ExecuteString(func(w io.Writer, tag string) (int, error) {
			return io.WriteString(w, "Hello")
		}))
	}
	{
		tmpl := template.Parse(largeTemplate)
		t.Logf("Large Template len before compilation: %v", len(largeTemplate))
		t.Logf("Large Template len after compilation: %v", len(tmpl.String()))
	}
	{
		tmpl := template.Parse(largeTemplate)
		t.Log(tmpl.ID())
	}
}
