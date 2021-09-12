package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type home struct {
	app.Compo
}

func (home *home) Render() app.UI {
	return app.Body().Body(
		app.Header().Body(
			&navbar{},
			app.Div().Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
						app.H2().Class("title").Text("Herzlich Willkommen"),
						app.P().Text("auf meiner Website kreatiVroni.de. Ich freue mich sehr, dass du hier gelandet bist und dir meine kreativen Projekte ansiehst. Ich habe es im August 2021 endlich gewagt und einen kleinen Onlineshop eröffnet. Egal, ob ein Einkauf in meinem Shop, ein Feedback zur Website oder ein Herz auf Instagram, ich bin dir für jede Unterstützung sehr dankbar!"),
					),
					app.Div().Class("img-box").Body(
						app.Img().Src("https://storage.googleapis.com/kreativroni-web/Haaand.PNG"),
					),
				),
			),
		),
	)
}
