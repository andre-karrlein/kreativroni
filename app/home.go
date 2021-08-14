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
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
					),
				),
			),
		),
	)
}
