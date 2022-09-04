package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type navbar struct {
	app.Compo
}

func (navbar *navbar) Render() app.UI {
	return app.Nav().Class("container flex justify-around p-10 mx-auto bg-opacity-0 rounded-xl flex-col md:flex-row").Body(
		app.Div().Class("flex items-center max-h-4 max-w-sm").Body(
			app.A().Href("/").Body(
				app.Img().Src("/web/Logo2.PNG"),
			),
		),
		app.Div().Class("items-center py-8 md:py-0 md:space-x-8 flex flex-col md:flex-row").Body(
			app.A().Class("text-black hover:text-pink-300").Href("/shop").Text("Shop"),
			app.A().Class("text-black hover:text-pink-300").Href("/fashion").Text("Fashion"),
			app.A().Class("text-black hover:text-pink-300").Href("https://www.etsy.com/de/shop/kreatiVroni").Text("Etsy"),
			app.A().Class("text-black hover:text-pink-300").Href("https://www.instagram.com/kreativroni/?hl=de").Text("Instagram"),
			app.A().Class("text-black hover:text-pink-300").Href("/about_me").Text("Über mich"),
			app.A().Class("text-black hover:text-pink-300").Href("mailto:kreativroni@gmail.com").Text("Kontakt"),
			app.A().Class("text-black hover:text-pink-300").Href("/legal_notice").Text("Impressum"),
		),
	)
}
