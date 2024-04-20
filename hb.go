package hb

import (
	"io"
	"slices"
	"strings"
)

type (
	Element struct {
		Attributes [][2]string
		Children   []Element
		Element    string
		attributes string
	}
)

var (
	None = Element{}

	voidElements = []string{
		"area",
		"base",
		"br",
		"col",
		"command",
		"embed",
		"hr",
		"img",
		"input",
		"keygen",
		"link",
		"meta",
		"param",
		"source",
		"track",
		"wbr",
	}
)

func (el *Element) Render(w io.Writer) error {
	return nil
}

func El(el, attribs string, children ...Element) Element {
	return Element{
		Element:    el,
		attributes: attribs,
		Children:   children,
	}
}

func Lazy(el, attribs string, children ...Element) func() Element {
	return func() Element {
		return Element{
			Element:    el,
			attributes: attribs,
			Children:   children,
		}
	}
}

func Classes(names ...string) string {
	return `class="` + strings.Join(names, " ") + `"`
}

func If(cond bool, then, otherwise func() Element) Element {
	if cond {
		return then()
	}
	return otherwise()
}

func When(cond bool, then func() Element) Element {
	return If(cond, then, func() Element {
		return None
	})
}

func isVoid(el string) bool {
	_, found := slices.BinarySearch(voidElements, el)
	return found
}
