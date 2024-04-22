package main

import (
	"net/http"

	"github.com/jlucasnsilva/hb"
	"github.com/jlucasnsilva/hb/ht"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		root := ht.HTML5(
			"HB - HTML Builder",
			"en",
			ht.CSSLink("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"),
			hb.Attr("container"),
			hb.Group(
				ht.Header("", ht.H1("", hb.Text("HB Sample Page"))),
				ht.Main("",
					ht.H2("", hb.Text("Background")),
					ht.P("",
						hb.Text(`HB was developed by `),
						ht.A(`href="https://github.com/jlucasnsilva" target="_blank"`,
							hb.Text("this guy")),
						hb.Text(` inspired by the Gomponents library.`),
					),
				),
				ht.Footer("", hb.Text("Thank you")),
			),
		)

		root.Render(w)
	})

	http.ListenAndServe(":8080", nil)
}
