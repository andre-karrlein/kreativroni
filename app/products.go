package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type products struct {
	app.Compo
}

func (products *products) Render() app.UI {
	return app.Body().Body(
		app.Header().Class("products-header").Body(
			&navbar{},
		),
		app.Div().Class("main").Body(
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
						app.Text("TESTING PRODUCT PAGE"),
					),
				),
			),
		),
	)
}
