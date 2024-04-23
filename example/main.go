package main

import (
	"net/http"

	"github.com/jlucasnsilva/hb"
	"github.com/jlucasnsilva/hb/ht"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		root := ht.HTML(hb.Attr(`lang="en"`),
			ht.Head(nil,
				ht.CSSLink(
					"https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css",
				),
				ht.Title(nil, hb.Text("HB - HTML Builder")),
			),
			ht.Body(
				hb.Attr(`class="container"`),
				ht.Header(nil, ht.H1(nil, hb.Text("HB Sample Page"))),
				ht.Main(nil,
					ht.H2(nil, hb.Text("Background")),
					ht.P(nil,
						hb.Text(`HB was developed by `),
						ht.A(
							hb.Attr(
								`href="https://github.com/jlucasnsilva" target="_blank"`,
							),
							hb.Text("this guy"),
						),
						hb.Text(` inspired by the Gomponents library.`),
					),
				),
				ht.Footer(nil, hb.Text("Thank you")),
			),
		)

		root.Render(w)
	})

	http.ListenAndServe(":8080", nil)
}
