package hb

import (
	"html/template"
	"io"
	"slices"
	"strings"
)

type (
	Element struct {
		children   []Element
		attributes string
		element    string
		typ        int
	}

	Attributes struct {
		attr strings.Builder
	}
)

const SlotTag = "<|slot|>"

const (
	typeEl = iota
	typeText
	typeGroup
	typeTemplate
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
	switch el.typ {
	case typeEl:
		if el.element == "" {
			return nil
		}

		writeString(w, "<", el.element, el.attributes)
		if isVoid(el.element) {
			writeString(w, " />")
			return nil
		}
		writeString(w, ">")

		for _, child := range el.children {
			child.Render(w)
		}

		writeString(w, "</", el.element, ">")
	case typeText:
		text := template.HTMLEscapeString(el.element)
		writeString(w, text)
	case typeGroup:
		for _, child := range el.children {
			child.Render(w)
		}
	case typeTemplate:
		if len(el.children) < 1 {
			writeString(w, el.element)
		} else {
			parts := strings.Split(el.element, SlotTag)
			if len(parts) != 2 {
				panic("invalid template: only one slot is available and required per template")
			}
			writeString(w, parts[0])
			for _, child := range el.children {
				child.Render(w)
			}
			writeString(w, parts[1])
		}
	}

	return nil
}

func writeString(w io.Writer, parts ...string) {
	for _, s := range parts {
		io.WriteString(w, s)
	}
}

func E(el, attribs string, children ...Element) Element {
	pad := ""
	if attribs != "" {
		pad = " "
	}

	res := Element{
		element:    el,
		children:   children,
		attributes: pad + attribs,
	}
	return res
}

func Group(children ...Element) Element {
	return Element{
		children: children,
		typ:      typeGroup,
	}
}

func Text(text string) Element {
	return Element{
		element: template.HTMLEscapeString(text),
		typ:     typeText,
	}
}

func Lazy(el, attribs string, children ...Element) func() Element {
	return func() Element {
		return E(el, attribs, children...)
	}
}

func Template(raw string, children ...Element) Element {
	return Element{
		element:  raw,
		children: children,
		typ:      typeTemplate,
	}
}

func If(cond bool, then Element, otherwise ...Element) Element {
	if cond {
		return then
	}
	if len(otherwise) < 1 {
		return None
	}
	return otherwise[0]
}

func When(cond bool, then func() Element, otherwise ...func() Element) Element {
	if cond {
		return then()
	}
	if len(otherwise) < 1 {
		return None
	}
	return otherwise[0]()
}

func isVoid(el string) bool {
	_, found := slices.BinarySearch(voidElements, el)
	return found
}

func Attr(init string) *Attributes {
	a := &Attributes{}
	if init != "" {
		a.write(init)
	}
	return a
}

func (attr *Attributes) Attr(key, value string, cond ...bool) *Attributes {
	if key == "" {
		return attr
	}
	k := template.HTMLEscapeString(key)

	for _, c := range cond {
		if !c {
			return attr
		}
	}

	if value == "" {
		attr.write(k)
		return attr
	}

	v := template.HTMLEscapeString(value)
	attr.write(k, `="`, v, `"`)
	return attr
}

func (attr *Attributes) Flag(key string, cond bool) *Attributes {
	if cond {
		attr.write(template.HTMLEscapeString(key))
	}
	return attr
}

func (attr *Attributes) Case(key string, m map[string]bool) *Attributes {
	i := 0
	attr.write(template.HTMLEscapeString(key), `="`)
	for val, cond := range m {
		if cond {
			if i > 0 {
				attr.writePlain(" ")
			}
			v := template.HTMLEscapeString(val)
			attr.writePlain(v)
			i++
		}
	}
	attr.writePlain(`"`)
	return attr
}

func (attr *Attributes) Switch(key string, idx int, values []string) *Attributes {
	if len(values) < 1 {
		panic("switch error: values cannot be empty")
	}
	if idx < 0 || idx >= len(values) {
		panic("switch error: idx value out of bounds")
	}

	attr.write(template.HTMLEscapeString(key), `="`)
	v := template.HTMLEscapeString(values[idx])
	attr.writePlain(v, `"`)
	return attr
}

func (attr *Attributes) Raw(key, value string, cond ...bool) *Attributes {
	for _, c := range cond {
		if !c {
			return attr
		}
	}

	attr.write(key, `="`, value, `"`)
	return attr
}

func (attr *Attributes) String() string {
	return attr.attr.String()
}

func (attr *Attributes) write(parts ...string) {
	if attr.attr.Len() > 0 {
		attr.attr.WriteString(" ")
	}
	for _, s := range parts {
		attr.attr.WriteString(s)
	}
}

func (attr *Attributes) writePlain(parts ...string) {
	for _, s := range parts {
		attr.attr.WriteString(s)
	}
}
