package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type about struct {
	app.Compo
}

func (about *about) Render() app.UI {
	return app.Div().Class("min-h-screen w-full bg-gradient-to-r from-green-200 to-green-500 p-0 pt-5 md:p-8 md:py-10").Body(
		&navbar{},
		app.Div().Class("pt-0 md:pt-20").Body(
			app.Figure().Class("bg-white bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-xl p-4 md:flex md:p-30").Body(
				app.Div().Class("pt-6 md:p-8 text-center md:text-left space-y-4").Body(
					app.H2().Class("text-5xl font-bold title").Text("Über mich"),
					app.Blockquote().Class("pt-5").Body(
						app.P().Class("text-lg font-semibold").Text(
							"Ich bin Vroni, 27 Jahre alt und wohne in Bayern im Berchtesgadener Land. Ich komme eigentlich aus Unterfranken, in der Nähe von Würzburg, habe dort Pädagogik und Philosophie studiert und arbeite jetzt in der Kinder- und Jugendarbeit. Den Online-Shop führe ich nebenberuflich und erfülle mir damit einen Traum, den ich schon seit mehreren Jahren habe. Ich liebe es kreativ zu sein und kann dabei die Zeit und alle Sorgen vergessen.",
						),
					),
				),
				app.Img().Class("w-32 h-32 md:w-48 md:h-auto md:rounded-none rounded-full mx-auto").Src("https://storage.googleapis.com/kreativroni-web/Schaedel.png"),
			),
		),
	)
}
