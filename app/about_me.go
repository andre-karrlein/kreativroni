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
						app.P().Text("Ich bin Vroni, 27 Jahre alt und wohne in Bayern im Berchtesgadener Land. Ich komme eigentlich aus Unterfranken, in der Nähe von Würzburg, habe dort Pädagogik und Philosophie studiert und arbeite jetzt in der Kinder- und Jugendarbeit. Den Online-Shop führe ich nebenberuflich und erfülle mir damit einen Traum, den ich schon seit mehreren Jahren habe. Ich liebe es kreativ zu sein und kann dabei die Zeit und alle Sorgen vergessen."),
					),
					app.Div().Class("img-box").Body(
						app.Img().Src("https://storage.googleapis.com/kreativroni-web/Schaedel.png"),
					),
				),
			),
		),
	)
}
