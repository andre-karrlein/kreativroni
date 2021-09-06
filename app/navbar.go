package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type navbar struct {
	app.Compo
}

func (navbar *navbar) Render() app.UI {
	return app.Nav().Class("navbar").Body(
		app.Div().Class("logo").Body(
			app.A().Href("/").Body(
				app.Img().Height(200).Width(300).Src("https://storage.googleapis.com/kreativroni-web/Logo2.PNG"),
			),
		),
		app.Ul().Class("menu").Body(
			app.Li().Body(
				app.A().Href("/about_me").Text("Über mich"),
			),
			app.Li().Body(
				app.A().Href("/shop").Text("Shop"),
			),
			app.Li().Body(
				app.A().Target("_blank").Href("https://www.instagram.com/kreativroni/?hl=de").Text("Instagram"),
			),
			app.Li().Body(
				app.A().Href("https://www.etsy.com/de/shop/kreatiVroni").Text("Etsy"),
			),
			app.Li().Body(
				app.A().Href("/legal_notice").Text("Impressum"),
			),
		),
	)
}
