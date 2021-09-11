package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type legal_notice struct {
	app.Compo
}

func (legal_notice *legal_notice) Render() app.UI {
	return app.Body().Body(
		app.Header().Body(
			&navbar{},
			app.Div().Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
						app.H2().Class("title").Text("Impressum"),
						app.P().Text("kreatiVroni"),
						app.P().Text("Kreatives von Vroni"),
						app.Br(),
						app.H3().Text("Betreiber der Website:"),
						app.P().Text("Veronika Karrlein"),
						app.Br(),
						app.H3().Text("Adresse:"),
						app.P().Text("Tittmoninger Str. 2"),
						app.P().Text("83410 Laufen"),
						app.Br(),
						app.H3().Text("Kontakt:"),
						app.P().Text("Telefon: 0049 151 11512019"),
						app.P().Text("Email: kreatiVroni@gmail.com"),
						app.P().Text("Internet: kreatiVroni.de"),
						app.Br(),
						app.H3().Text("Umstatzsteuer-ID:"),
						app.P().Text("DE345010837"),
					),
				),
			),
		),
	)
}
