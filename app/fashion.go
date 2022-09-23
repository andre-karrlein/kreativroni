package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type fashion struct {
	app.Compo
}

func (fashion *fashion) OnPreRender(ctx app.Context) {
	ctx.Navigate("https://www.spreadshirt.de/shop/user/kreativroni+-+kreatives+von+vroni")
}

func (fashion *fashion) Render() app.UI {
	return app.Div().Class("min-h-screen w-full bg-gradient-to-r from-green-200 to-green-500 p-0 pt-5 md:p-8 md:py-10").Body(
		&navbar{},
		app.Div().Class("pt-0 md:pt-20").Body(
			app.Figure().Class("bg-white bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-xl p-4 md:flex md:p-30").Body(
				app.Div().Class("pt-6 md:p-8 text-center md:text-left space-y-4").Body(
					app.H2().Class("text-5xl font-bold title").Text("Herzlich Willkommen"),
					app.Blockquote().Class("pt-5").Body(
						app.P().Class("text-lg font-semibold").Text(
							"auf meiner Website kreatiVroni.de. Ich freue mich sehr, dass du hier gelandet bist und dir meine kreativen Projekte ansiehst. Ich habe es im August 2021 endlich gewagt und einen kleinen Onlineshop eröffnet. Egal, ob ein Einkauf in meinem Shop, ein Feedback zur Website oder ein Herz auf Instagram, ich bin dir für jede Unterstützung sehr dankbar!",
						),
					),
				),
				app.Img().Class("w-32 h-32 md:w-48 md:h-auto md:rounded-none rounded-full mx-auto").Src("/web/Haaand.png"),
			),
		),
	)
}
