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
						app.P().Text("Some text"),
					),
				),
			),
		),
	)
}
