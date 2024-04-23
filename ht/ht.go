package ht

import "github.com/jlucasnsilva/hb"

func MetaViewport() hb.Element {
	return Meta(
		hb.Attr(
			`name="viewport" content="width=device-width, initial-scale=1.0"`,
		),
	)
}

func HTMLDoctype(children ...hb.Element) hb.Element {
	return hb.Template(
		"<!DOCTYPE html>"+hb.SlotTag,
		children...,
	)
}

func CSSLink(filepath string) hb.Element {
	return Link(hb.KV(
		"rel", "stylesheet",
		"href", filepath,
	))
}

func ScriptFile(deferFlag bool, filepath string) hb.Element {
	return Script(
		hb.Flag("defer", deferFlag).
			KV("src", filepath),
	)
}

func A(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("a", attribs, children...)
}

func Abbr(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("abbr", attribs, children...)
}

func Acronym(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("acronym", attribs, children...)
}

func Address(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("address", attribs, children...)
}

func Applet(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("applet", attribs, children...)
}

func Area(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("area", attribs, children...)
}

func Article(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("article", attribs, children...)
}

func Aside(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("aside", attribs, children...)
}

func Audio(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("audio", attribs, children...)
}

func B(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("b", attribs, children...)
}

func Base(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("base", attribs, children...)
}

func Basefont(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("basefont", attribs, children...)
}

func Bdi(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("bdi", attribs, children...)
}

func Bdo(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("bdo", attribs, children...)
}

func Big(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("big", attribs, children...)
}

func Blockquote(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("blockquote", attribs, children...)
}

func Body(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("body", attribs, children...)
}

func Br(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("br", attribs, children...)
}

func Button(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("button", attribs, children...)
}

func Canvas(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("canvas", attribs, children...)
}

func Caption(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("caption", attribs, children...)
}

func Center(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("center", attribs, children...)
}

func Cite(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("cite", attribs, children...)
}

func Code(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("code", attribs, children...)
}

func Col(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("col", attribs, children...)
}

func ColGroup(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("colgroup", attribs, children...)
}

func Data(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("data", attribs, children...)
}

func Datalist(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("datalist", attribs, children...)
}

func Dd(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("dd", attribs, children...)
}

func Del(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("del", attribs, children...)
}

func Details(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("details", attribs, children...)
}

func Dfn(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("dfn", attribs, children...)
}

func Dialog(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("dialog", attribs, children...)
}

func Dir(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("dir", attribs, children...)
}

func Div(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("div", attribs, children...)
}

func Dl(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("dl", attribs, children...)
}

func Dt(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("dt", attribs, children...)
}

func Em(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("em", attribs, children...)
}

func Embed(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("embed", attribs, children...)
}

func FieldSet(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("fieldset", attribs, children...)
}

func FigCaption(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("figcaption", attribs, children...)
}

func Figure(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("figure", attribs, children...)
}

func Font(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("font", attribs, children...)
}

func Footer(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("footer", attribs, children...)
}

func Form(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("form", attribs, children...)
}

func Frame(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("frame", attribs, children...)
}

func FrameSet(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("frameset", attribs, children...)
}

func H1(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("h1", attribs, children...)
}

func H2(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("h2", attribs, children...)
}

func H3(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("h3", attribs, children...)
}

func H4(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("h4", attribs, children...)
}

func H5(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("h5", attribs, children...)
}

func H6(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("h6", attribs, children...)
}

func Head(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("head", attribs, children...)
}

func Header(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("header", attribs, children...)
}

func HGroup(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("hgroup", attribs, children...)
}

func Hr(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("hr", attribs, children...)
}

func HTML(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("html", attribs, children...)
}

func I(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("i", attribs, children...)
}

func Iframe(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("iframe", attribs, children...)
}

func Img(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("img", attribs, children...)
}

func Input(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("input", attribs, children...)
}

func Ins(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("ins", attribs, children...)
}

func Kbd(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("kbd", attribs, children...)
}

func Label(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("label", attribs, children...)
}

func Legend(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("legend", attribs, children...)
}

func Li(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("li", attribs, children...)
}

func Link(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("link", attribs, children...)
}

func Main(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("main", attribs, children...)
}

func Map(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("map", attribs, children...)
}

func Mark(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("mark", attribs, children...)
}

func Menu(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("menu", attribs, children...)
}

func Meta(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("meta", attribs, children...)
}

func Meter(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("meter", attribs, children...)
}

func Nav(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("nav", attribs, children...)
}

func NoFrames(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("noframes", attribs, children...)
}

func NoScript(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("noscript", attribs, children...)
}

func Object(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("object", attribs, children...)
}

func Ol(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("ol", attribs, children...)
}

func OptGroup(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("optgroup", attribs, children...)
}

func Option(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("option", attribs, children...)
}

func Output(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("output", attribs, children...)
}

func P(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("p", attribs, children...)
}

func Param(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("param", attribs, children...)
}

func Picture(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("picture", attribs, children...)
}

func Pre(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("pre", attribs, children...)
}

func Progress(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("progress", attribs, children...)
}

func Q(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("q", attribs, children...)
}

func Rp(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("rp", attribs, children...)
}

func Rt(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("rt", attribs, children...)
}

func Ruby(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("ruby", attribs, children...)
}

func S(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("s", attribs, children...)
}

func Samp(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("samp", attribs, children...)
}

func Script(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("script", attribs, children...)
}

func Search(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("search", attribs, children...)
}

func Section(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("section", attribs, children...)
}

func Select(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("select", attribs, children...)
}

func Small(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("small", attribs, children...)
}

func Source(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("source", attribs, children...)
}

func Span(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("span", attribs, children...)
}

func Strike(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("strike", attribs, children...)
}

func Strong(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("strong", attribs, children...)
}

func Style(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("style", attribs, children...)
}

func Sub(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("sub", attribs, children...)
}

func Summary(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("summary", attribs, children...)
}

func Sup(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("sup", attribs, children...)
}

func SVG(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("svg", attribs, children...)
}

func Table(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("table", attribs, children...)
}

func TBody(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("tbody", attribs, children...)
}

func Td(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("td", attribs, children...)
}

func Template(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("template", attribs, children...)
}

func TextArea(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("textarea", attribs, children...)
}

func TFoot(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("tfoot", attribs, children...)
}

func Th(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("th", attribs, children...)
}

func THead(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("thead", attribs, children...)
}

func Time(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("time", attribs, children...)
}

func Title(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("title", attribs, children...)
}

func Tr(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("tr", attribs, children...)
}

func Track(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("track", attribs, children...)
}

func Tt(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("tt", attribs, children...)
}

func U(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("u", attribs, children...)
}

func Ul(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("ul", attribs, children...)
}

func Var(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("var", attribs, children...)
}

func Video(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("video", attribs, children...)
}

func Wbr(attribs *hb.Attributes, children ...hb.Element) hb.Element {
	return hb.E("wbr", attribs, children...)
}
