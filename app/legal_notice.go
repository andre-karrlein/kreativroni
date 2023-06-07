package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type legal_notice struct {
	app.Compo
}

func (legal_notice *legal_notice) Render() app.UI {
	return app.Div().Class("min-h-screen w-full bg-gradient-to-r from-green-200 to-green-500 p-0 pt-5 md:p-8 md:py-10").Body(
		&navbar{},
		app.Div().Class("pt-0 md:pt-20").Body(
			app.Figure().Class("bg-white bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-xl p-4 md:flex md:p-30").Body(
				app.Div().Class("pt-6 md:p-8 text-center md:text-left space-y-4").Body(
					app.H2().Class("text-5xl font-bold title").Text("Impressum"),
					app.Blockquote().Class("pt-5").Body(
						app.P().Text("kreatiVroni"),
						app.P().Text("Kreatives von Vroni"),
						app.Br(),
						app.H3().Text("Betreiber der Website:"),
						app.P().Text("Veronika Karrlein"),
						app.Br(),
						app.H3().Text("Adresse:"),
						app.P().Text("Münchener Str. 25F"),
						app.P().Text("83395 Freilassing"),
						app.Br(),
						app.H3().Text("Kontakt:"),
						app.P().Text("Telefon: 0043 664 75140497"),
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
