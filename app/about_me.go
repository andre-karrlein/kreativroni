package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type about struct {
	app.Compo
}

func (about *about) Render() app.UI {
	return app.Body().Body(
		app.Header().Body(
			&navbar{},
			app.Div().Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
						app.H2().Class("title").Text("Über mich"),
					),
				),
			),
		),
	)
}
