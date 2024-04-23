package hb

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	eTestCase struct {
		label    string
		expected string
		el       Element
	}

	attrTestCase struct {
		label    string
		expected string
		attr     *Attributes
	}
)

func TestWriteText(t *testing.T) {
	buf := strings.Builder{}
	text := "Hello world"
	el := Text(text)

	assert.Nil(t, el.Render(&buf))
	assert.Equal(t, text, buf.String())
}

func TestE(t *testing.T) {
	testCases := []eTestCase{
		{
			label:    "simple case should succeed",
			expected: `<h1>Title</h1>`,
			el:       E("h1", nil, Text("Title")),
		},
		{
			label:    "void element should render correctly",
			expected: `<img src="/img/ninja.png" alt="Ninja pic" />`,
			el:       E("img", Attr(`src="/img/ninja.png" alt="Ninja pic"`)),
		},
		{
			label:    "children should render correctly",
			expected: `<ul class="list"><li class="item">A</li><li class="item"><button onclick="bye()">Exit</button></li></ul>`,
			el: E("ul", Attr(`class="list"`),
				E("li", Attr(`class="item"`), Text("A")),
				E("li", Attr(`class="item"`),
					E("button", Attr(`onclick="bye()"`), Text("Exit")),
				),
			),
		},
		{
			label:    "conditional elements should render correctly",
			expected: `<div><p>Hello, world!</p></div>`,
			el: E("div", nil,
				If(true,
					E("p", nil, Text("Hello, world!")),
					E("h1", nil, Text("Hello, world!")),
				),
			),
		},
		{
			label:    "conditional elements should render nothing",
			expected: `<div></div>`,
			el: E("div", nil,
				If(false, E("p", nil, Text("Hello, world!"))),
			),
		},
		{
			label:    "conditional elements should render correctly",
			expected: `<div><h1>Hello, world!</h1></div>`,
			el: E("div", nil,
				If(false,
					E("p", nil, Text("Hello, world!")),
					E("h1", nil, Text("Hello, world!")),
				),
			),
		},
		{
			label:    "conditional lazy elements should render correctly",
			expected: `<div><p>Hello, world!</p></div>`,
			el: E("div", nil,
				When(true,
					Lazy("p", nil, Text("Hello, world!")),
					Lazy("h1", nil, Text("Hello, world!")),
				),
			),
		},
		{
			label:    "conditional lazy elements should render nothing",
			expected: `<div></div>`,
			el: E("div", nil,
				When(false, Lazy("p", nil, Text("Hello, world!"))),
			),
		},
		{
			label:    "conditional lazy elements should render correctly",
			expected: `<div><h1>Hello, world!</h1></div>`,
			el: E("div", nil,
				When(false,
					Lazy("p", nil, Text("Hello, world!")),
					Lazy("h1", nil, Text("Hello, world!")),
				),
			),
		},
		{
			label:    "should render group correctly",
			expected: `<li>A</li><li>B</li><li>C</li><li>D</li>`,
			el: Group(
				E("li", nil, Text("A")),
				E("li", nil, Text("B")),
				E("li", nil, Text("C")),
				E("li", nil, Text("D")),
			),
		},
		{
			label:    "empty template should render correctly",
			expected: "<html><head><title>New page</title></head><body></body></html>",
			el: Template(
				"<html><head><title>New page</title></head><body></body></html>",
			),
		},
		{
			label:    "template with children should render correctly",
			expected: "<html><head><title>New page</title></head><body><ul><li>A</li><li>B</li></ul></body></html>",
			el: Template(
				"<html><head><title>New page</title></head><body><ul>"+SlotTag+"</ul></body></html>",
				E("li", nil, Text("A")),
				E("li", nil, Text("B")),
			),
		},
		{
			label:    "template with children should render correctly",
			expected: "<!DOCTYPE html><html></html>",
			el:       Template("<!DOCTYPE html>"+SlotTag, E("html", nil)),
		},
		{
			label:    "map should succeed",
			expected: `<li>A</li><li>B</li>`,
			el: Map([]string{"A", "B"}, func(s string) Element {
				return E("li", nil, Text(s))
			}),
		},
	}

	for i, tc := range testCases {
		label := fmt.Sprintf("#%v - %v", i, tc.label)
		t.Run(label, func(t *testing.T) {
			buf := strings.Builder{}
			err := tc.el.Render(&buf)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, buf.String())
		})
	}
}

func TestAttr(t *testing.T) {
	testCases := []attrTestCase{
		{
			label:    "should render nothing",
			expected: "",
			attr:     nil,
		},
		{
			label:    "should render nothing",
			expected: "",
			attr:     Attr(""),
		},
		{
			label:    "should render flag",
			expected: "required",
			attr:     Attr("required"),
		},
		{
			label:    "should render class",
			expected: `class="container"`,
			attr:     Attr(`class="container"`),
		},
		{
			label:    "should render class",
			expected: `class="container"`,
			attr:     Attr(`class="container"`),
		},
		{
			label:    "should render class and required and type",
			expected: `class="container" required type="text"`,
			attr:     Attr(`class="container"`).KV("required", "", "type", "text"),
		},
		{
			label:    "should render class and required flag",
			expected: `class="container" required`,
			attr: Attr(`class="container"`).
				Flag("required", true).
				Flag("open", false),
		},
		{
			label:    "should render class and required flag",
			expected: `class="container" required`,
			attr: KV("class", "container").
				Maybe("required", "", true).
				Maybe("open", "", false),
		},
		{
			label:    "should escape sequence",
			expected: `class="&lt;&gt;"`,
			attr:     KV("class", "<>"),
		},
		{
			label:    "should NOT escape sequence",
			expected: `class="<>"`,
			attr:     Raw("class", "<>"),
		},
		{
			label:    "switch should render correctly",
			expected: `type="select"`,
			attr:     Switch("type", 2, []string{"text", "submit", "select"}),
		},
		{
			label:    "comp should render correctly",
			expected: `class="btn btn-outline"`,
			attr:     Comp("class", []string{"btn", "btn-outline"}),
		},
	}

	for i, tc := range testCases {
		label := fmt.Sprintf("#%v - %v", i, tc.label)
		t.Run(label, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.attr.String())
		})
	}
}

func TestCase(t *testing.T) {
	t.Run("should render only correct classes", func(t *testing.T) {
		attr := Attr("required").
			Case("class", map[string]bool{
				"container": true,
				"w-50":      false,
				"w-full":    true,
				"bg-light":  false,
				"bg-dark":   true,
			}).
			String()

		assert.True(t, strings.Contains(attr, "container"))
		assert.False(t, strings.Contains(attr, "w-50"))
		assert.True(t, strings.Contains(attr, "w-full"))
		assert.False(t, strings.Contains(attr, "bg-light"))
		assert.True(t, strings.Contains(attr, "bg-dark"))
	})
}
