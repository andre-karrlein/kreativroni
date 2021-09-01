package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type privacy_policy struct {
	app.Compo
}

func (privacy_policy *privacy_policy) Render() app.UI {
	return app.Body().Body(
		app.Header().Body(
			&navbar{},
			app.Div().Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
						app.H2().Class("title").Text("Datenschutz"),
						app.P().Text("Some text"),
					),
				),
			),
		),
	)
}
