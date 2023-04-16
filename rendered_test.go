package template

import "testing"

func TestStringifyRenderedNodes(t *testing.T) {
	n := RenderedNodes{
		{
			Tag: "p",
			Attributes: []string{
				`class="adwjdkanwdj awjkd nakwd"`,
			},
			Children: RenderedNodes{
				{
					Tag:  "h1",
					Data: "bla",
				},
			},
		},
	}
	t.Log(n.String())
}

func TestEqualityRenderedNodes(t *testing.T) {
	n1 := RenderedNodes{
		{
			Tag: "p",
			Attributes: []string{
				`class="adwjdkanwdj awjkd nakwd"`,
			},
			Children: RenderedNodes{
				{
					Tag:  "h1",
					Data: "bla",
				},
			},
		},
	}
	n2 := RenderedNodes{
		{
			Tag: "p",
			Attributes: []string{
				`class="adwjdkanwdj awjkd nakwd"`,
			},
			Children: RenderedNodes{
				{
					Tag:  "h1",
					Data: "bla",
				},
			},
		},
	}
	t.Log(n1.Equals(n2))
}
