package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type fashion struct {
	app.Compo
}

func (fashion *fashion) OnNav(ctx app.Context) {
	ctx.Navigate("https://www.spreadshirt.de/shop/user/kreativroni+-+kreatives+von+vroni")
}

func (fashion *fashion) Render() app.UI {
	return app.Div().Class("min-h-screen w-full bg-gradient-to-r from-green-200 to-green-500 p-0 pt-5 md:p-8 md:py-10").Body(
		&navbar{},
		app.Div().Class("pt-0 md:pt-20").Body(),
	)
}
