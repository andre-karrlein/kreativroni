package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type shop struct {
	app.Compo
}

func (shop *shop) Render() app.UI {
	products := getProducts()

	return app.Body().Class("shop-body").Body(
		app.Div().Class("shop-header").Body(
			&navbar{},
		),
		app.Div().Class("container").Body(
			app.Range(products).Slice(func(i int) app.UI {
				return app.Div().Class("card").Body(
					app.Div().Class("imgBx").Body(
						app.Img().Src(products[i].Image),
						app.H2().Text(products[i].Name),
					),
					app.Div().Class("content").Body(
						app.A().Href(products[i].Link).Text("Buy Now"),
					),
				)
			}),
		),
	)
}
