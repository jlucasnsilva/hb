package ht

import "github.com/jlucasnsilva/hb"

func HTML5(
	title string,
	lang string,
	head hb.Element,
	bodyAttr *hb.Attributes,
	body hb.Element,
) hb.Element {
	return HTMLDoctype(
		HTML(
			hb.Attr("", "lang", lang).String(),
			Head("",
				Meta(`name="viewport" content="width=device-width, initial-scale=1.0"`),
				head,
				Title("", hb.Text(title)),
			),
			Body(bodyAttr.String(), body),
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
	return Link(hb.Attr(`rel="stylesheet"`, "href", filepath).String())
}

func ScriptFile(deferFlag bool, filepath string) hb.Element {
	return Script(
		hb.Attr("").
			Attr("defer", "", deferFlag).
			Attr("src", filepath).
			String(),
	)
}

func A(attribs string, children ...hb.Element) hb.Element {
	return hb.E("a", attribs, children...)
}

func Abbr(attribs string, children ...hb.Element) hb.Element {
	return hb.E("abbr", attribs, children...)
}

func Acronym(attribs string, children ...hb.Element) hb.Element {
	return hb.E("acronym", attribs, children...)
}

func Address(attribs string, children ...hb.Element) hb.Element {
	return hb.E("address", attribs, children...)
}

func Applet(attribs string, children ...hb.Element) hb.Element {
	return hb.E("applet", attribs, children...)
}

func Area(attribs string, children ...hb.Element) hb.Element {
	return hb.E("area", attribs, children...)
}

func Article(attribs string, children ...hb.Element) hb.Element {
	return hb.E("article", attribs, children...)
}

func Aside(attribs string, children ...hb.Element) hb.Element {
	return hb.E("aside", attribs, children...)
}

func Audio(attribs string, children ...hb.Element) hb.Element {
	return hb.E("audio", attribs, children...)
}

func B(attribs string, children ...hb.Element) hb.Element {
	return hb.E("b", attribs, children...)
}

func Base(attribs string, children ...hb.Element) hb.Element {
	return hb.E("base", attribs, children...)
}

func Basefont(attribs string, children ...hb.Element) hb.Element {
	return hb.E("basefont", attribs, children...)
}

func Bdi(attribs string, children ...hb.Element) hb.Element {
	return hb.E("bdi", attribs, children...)
}

func Bdo(attribs string, children ...hb.Element) hb.Element {
	return hb.E("bdo", attribs, children...)
}

func Big(attribs string, children ...hb.Element) hb.Element {
	return hb.E("big", attribs, children...)
}

func Blockquote(attribs string, children ...hb.Element) hb.Element {
	return hb.E("blockquote", attribs, children...)
}

func Body(attribs string, children ...hb.Element) hb.Element {
	return hb.E("body", attribs, children...)
}

func Br(attribs string, children ...hb.Element) hb.Element {
	return hb.E("br", attribs, children...)
}

func Button(attribs string, children ...hb.Element) hb.Element {
	return hb.E("button", attribs, children...)
}

func Canvas(attribs string, children ...hb.Element) hb.Element {
	return hb.E("canvas", attribs, children...)
}

func Caption(attribs string, children ...hb.Element) hb.Element {
	return hb.E("caption", attribs, children...)
}

func Center(attribs string, children ...hb.Element) hb.Element {
	return hb.E("center", attribs, children...)
}

func Cite(attribs string, children ...hb.Element) hb.Element {
	return hb.E("cite", attribs, children...)
}

func Code(attribs string, children ...hb.Element) hb.Element {
	return hb.E("code", attribs, children...)
}

func Col(attribs string, children ...hb.Element) hb.Element {
	return hb.E("col", attribs, children...)
}

func ColGroup(attribs string, children ...hb.Element) hb.Element {
	return hb.E("colgroup", attribs, children...)
}

func Data(attribs string, children ...hb.Element) hb.Element {
	return hb.E("data", attribs, children...)
}

func Datalist(attribs string, children ...hb.Element) hb.Element {
	return hb.E("datalist", attribs, children...)
}

func Dd(attribs string, children ...hb.Element) hb.Element {
	return hb.E("dd", attribs, children...)
}

func Del(attribs string, children ...hb.Element) hb.Element {
	return hb.E("del", attribs, children...)
}

func Details(attribs string, children ...hb.Element) hb.Element {
	return hb.E("details", attribs, children...)
}

func Dfn(attribs string, children ...hb.Element) hb.Element {
	return hb.E("dfn", attribs, children...)
}

func Dialog(attribs string, children ...hb.Element) hb.Element {
	return hb.E("dialog", attribs, children...)
}

func Dir(attribs string, children ...hb.Element) hb.Element {
	return hb.E("dir", attribs, children...)
}

func Div(attribs string, children ...hb.Element) hb.Element {
	return hb.E("div", attribs, children...)
}

func Dl(attribs string, children ...hb.Element) hb.Element {
	return hb.E("dl", attribs, children...)
}

func Dt(attribs string, children ...hb.Element) hb.Element {
	return hb.E("dt", attribs, children...)
}

func Em(attribs string, children ...hb.Element) hb.Element {
	return hb.E("em", attribs, children...)
}

func Embed(attribs string, children ...hb.Element) hb.Element {
	return hb.E("embed", attribs, children...)
}

func FieldSet(attribs string, children ...hb.Element) hb.Element {
	return hb.E("fieldset", attribs, children...)
}

func FigCaption(attribs string, children ...hb.Element) hb.Element {
	return hb.E("figcaption", attribs, children...)
}

func Figure(attribs string, children ...hb.Element) hb.Element {
	return hb.E("figure", attribs, children...)
}

func Font(attribs string, children ...hb.Element) hb.Element {
	return hb.E("font", attribs, children...)
}

func Footer(attribs string, children ...hb.Element) hb.Element {
	return hb.E("footer", attribs, children...)
}

func Form(attribs string, children ...hb.Element) hb.Element {
	return hb.E("form", attribs, children...)
}

func Frame(attribs string, children ...hb.Element) hb.Element {
	return hb.E("frame", attribs, children...)
}

func FrameSet(attribs string, children ...hb.Element) hb.Element {
	return hb.E("frameset", attribs, children...)
}

func H1(attribs string, children ...hb.Element) hb.Element {
	return hb.E("h1", attribs, children...)
}

func H2(attribs string, children ...hb.Element) hb.Element {
	return hb.E("h2", attribs, children...)
}

func H3(attribs string, children ...hb.Element) hb.Element {
	return hb.E("h3", attribs, children...)
}

func H4(attribs string, children ...hb.Element) hb.Element {
	return hb.E("h4", attribs, children...)
}

func H5(attribs string, children ...hb.Element) hb.Element {
	return hb.E("h5", attribs, children...)
}

func H6(attribs string, children ...hb.Element) hb.Element {
	return hb.E("h6", attribs, children...)
}

func Head(attribs string, children ...hb.Element) hb.Element {
	return hb.E("head", attribs, children...)
}

func Header(attribs string, children ...hb.Element) hb.Element {
	return hb.E("header", attribs, children...)
}

func HGroup(attribs string, children ...hb.Element) hb.Element {
	return hb.E("hgroup", attribs, children...)
}

func Hr(attribs string, children ...hb.Element) hb.Element {
	return hb.E("hr", attribs, children...)
}

func HTML(attribs string, children ...hb.Element) hb.Element {
	return hb.E("html", attribs, children...)
}

func I(attribs string, children ...hb.Element) hb.Element {
	return hb.E("i", attribs, children...)
}

func Iframe(attribs string, children ...hb.Element) hb.Element {
	return hb.E("iframe", attribs, children...)
}

func Img(attribs string, children ...hb.Element) hb.Element {
	return hb.E("img", attribs, children...)
}

func Input(attribs string, children ...hb.Element) hb.Element {
	return hb.E("input", attribs, children...)
}

func Ins(attribs string, children ...hb.Element) hb.Element {
	return hb.E("ins", attribs, children...)
}

func Kbd(attribs string, children ...hb.Element) hb.Element {
	return hb.E("kbd", attribs, children...)
}

func Label(attribs string, children ...hb.Element) hb.Element {
	return hb.E("label", attribs, children...)
}

func Legend(attribs string, children ...hb.Element) hb.Element {
	return hb.E("legend", attribs, children...)
}

func Li(attribs string, children ...hb.Element) hb.Element {
	return hb.E("li", attribs, children...)
}

func Link(attribs string, children ...hb.Element) hb.Element {
	return hb.E("link", attribs, children...)
}

func Main(attribs string, children ...hb.Element) hb.Element {
	return hb.E("main", attribs, children...)
}

func Map(attribs string, children ...hb.Element) hb.Element {
	return hb.E("map", attribs, children...)
}

func Mark(attribs string, children ...hb.Element) hb.Element {
	return hb.E("mark", attribs, children...)
}

func Menu(attribs string, children ...hb.Element) hb.Element {
	return hb.E("menu", attribs, children...)
}

func Meta(attribs string, children ...hb.Element) hb.Element {
	return hb.E("meta", attribs, children...)
}

func Meter(attribs string, children ...hb.Element) hb.Element {
	return hb.E("meter", attribs, children...)
}

func Nav(attribs string, children ...hb.Element) hb.Element {
	return hb.E("nav", attribs, children...)
}

func NoFrames(attribs string, children ...hb.Element) hb.Element {
	return hb.E("noframes", attribs, children...)
}

func NoScript(attribs string, children ...hb.Element) hb.Element {
	return hb.E("noscript", attribs, children...)
}

func Object(attribs string, children ...hb.Element) hb.Element {
	return hb.E("object", attribs, children...)
}

func Ol(attribs string, children ...hb.Element) hb.Element {
	return hb.E("ol", attribs, children...)
}

func OptGroup(attribs string, children ...hb.Element) hb.Element {
	return hb.E("optgroup", attribs, children...)
}

func Option(attribs string, children ...hb.Element) hb.Element {
	return hb.E("option", attribs, children...)
}

func Output(attribs string, children ...hb.Element) hb.Element {
	return hb.E("output", attribs, children...)
}

func P(attribs string, children ...hb.Element) hb.Element {
	return hb.E("p", attribs, children...)
}

func Param(attribs string, children ...hb.Element) hb.Element {
	return hb.E("param", attribs, children...)
}

func Picture(attribs string, children ...hb.Element) hb.Element {
	return hb.E("picture", attribs, children...)
}

func Pre(attribs string, children ...hb.Element) hb.Element {
	return hb.E("pre", attribs, children...)
}

func Progress(attribs string, children ...hb.Element) hb.Element {
	return hb.E("progress", attribs, children...)
}

func Q(attribs string, children ...hb.Element) hb.Element {
	return hb.E("q", attribs, children...)
}

func Rp(attribs string, children ...hb.Element) hb.Element {
	return hb.E("rp", attribs, children...)
}

func Rt(attribs string, children ...hb.Element) hb.Element {
	return hb.E("rt", attribs, children...)
}

func Ruby(attribs string, children ...hb.Element) hb.Element {
	return hb.E("ruby", attribs, children...)
}

func S(attribs string, children ...hb.Element) hb.Element {
	return hb.E("s", attribs, children...)
}

func Samp(attribs string, children ...hb.Element) hb.Element {
	return hb.E("samp", attribs, children...)
}

func Script(attribs string, children ...hb.Element) hb.Element {
	return hb.E("script", attribs, children...)
}

func Search(attribs string, children ...hb.Element) hb.Element {
	return hb.E("search", attribs, children...)
}

func Section(attribs string, children ...hb.Element) hb.Element {
	return hb.E("section", attribs, children...)
}

func Select(attribs string, children ...hb.Element) hb.Element {
	return hb.E("select", attribs, children...)
}

func Small(attribs string, children ...hb.Element) hb.Element {
	return hb.E("small", attribs, children...)
}

func Source(attribs string, children ...hb.Element) hb.Element {
	return hb.E("source", attribs, children...)
}

func Span(attribs string, children ...hb.Element) hb.Element {
	return hb.E("span", attribs, children...)
}

func Strike(attribs string, children ...hb.Element) hb.Element {
	return hb.E("strike", attribs, children...)
}

func Strong(attribs string, children ...hb.Element) hb.Element {
	return hb.E("strong", attribs, children...)
}

func Style(attribs string, children ...hb.Element) hb.Element {
	return hb.E("style", attribs, children...)
}

func Sub(attribs string, children ...hb.Element) hb.Element {
	return hb.E("sub", attribs, children...)
}

func Summary(attribs string, children ...hb.Element) hb.Element {
	return hb.E("summary", attribs, children...)
}

func Sup(attribs string, children ...hb.Element) hb.Element {
	return hb.E("sup", attribs, children...)
}

func SVG(attribs string, children ...hb.Element) hb.Element {
	return hb.E("svg", attribs, children...)
}

func Table(attribs string, children ...hb.Element) hb.Element {
	return hb.E("table", attribs, children...)
}

func TBody(attribs string, children ...hb.Element) hb.Element {
	return hb.E("tbody", attribs, children...)
}

func Td(attribs string, children ...hb.Element) hb.Element {
	return hb.E("td", attribs, children...)
}

func Template(attribs string, children ...hb.Element) hb.Element {
	return hb.E("template", attribs, children...)
}

func TextArea(attribs string, children ...hb.Element) hb.Element {
	return hb.E("textarea", attribs, children...)
}

func TFoot(attribs string, children ...hb.Element) hb.Element {
	return hb.E("tfoot", attribs, children...)
}

func Th(attribs string, children ...hb.Element) hb.Element {
	return hb.E("th", attribs, children...)
}

func THead(attribs string, children ...hb.Element) hb.Element {
	return hb.E("thead", attribs, children...)
}

func Time(attribs string, children ...hb.Element) hb.Element {
	return hb.E("time", attribs, children...)
}

func Title(attribs string, children ...hb.Element) hb.Element {
	return hb.E("title", attribs, children...)
}

func Tr(attribs string, children ...hb.Element) hb.Element {
	return hb.E("tr", attribs, children...)
}

func Track(attribs string, children ...hb.Element) hb.Element {
	return hb.E("track", attribs, children...)
}

func Tt(attribs string, children ...hb.Element) hb.Element {
	return hb.E("tt", attribs, children...)
}

func U(attribs string, children ...hb.Element) hb.Element {
	return hb.E("u", attribs, children...)
}

func Ul(attribs string, children ...hb.Element) hb.Element {
	return hb.E("ul", attribs, children...)
}

func Var(attribs string, children ...hb.Element) hb.Element {
	return hb.E("var", attribs, children...)
}

func Video(attribs string, children ...hb.Element) hb.Element {
	return hb.E("video", attribs, children...)
}

func Wbr(attribs string, children ...hb.Element) hb.Element {
	return hb.E("wbr", attribs, children...)
}
