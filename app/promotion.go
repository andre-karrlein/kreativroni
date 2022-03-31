package main

import (
	"html"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type promo struct {
	app.Compo
}

type promoProduct struct {
	title string
	image string
	url   string
}

func (promo *promo) Render() app.UI {
	promoProducts := []promoProduct{
		{
			title: "Grüne Grüße (Saatgut + Anhänger)",
			image: "https://i.etsystatic.com/31340310/r/il/651228/3823833011/il_570xN.3823833011_4row.jpg",
			url:   "https://etsy.me/36YSzFw",
		},
		{
			title: "Grüne Grüße „Du bist toll“ (+Anhänger)",
			image: "https://i.etsystatic.com/31340310/r/il/eb6b14/3776342372/il_570xN.3776342372_i0q5.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1206800341/grune-grusse-du-bist-toll-anhanger?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747642000",
		},
		{
			title: "Saatgut Gruß „Schmetterlinge im Bauch“",
			image: "https://i.etsystatic.com/31340310/r/il/c06ba1/3824197841/il_570xN.3824197841_91bf.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1192905340/saatgut-gruss-schmetterlinge-im-bauch?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747684000",
		},
		{
			title: "Saatgut Grüße „Schmetterlinge im Bauch“",
			image: "https://i.etsystatic.com/31340310/r/il/1712a7/3824215467/il_570xN.3824215467_l9mq.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1192909570/saatgut-grusse-schmetterlinge-im-bauch?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747754000",
		},
		{
			title: "Saatgut Grüße „Bee Happy“",
			image: "https://i.etsystatic.com/31340310/r/il/251ade/3824378685/il_570xN.3824378685_lge7.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1192951462/saatgut-grusse-bee-happy?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747784000",
		},
		{
			title: "Saatgut Gruß „Bienenwiese“",
			image: "https://i.etsystatic.com/31340310/r/il/62c51f/3824394185/il_570xN.3824394185_5wzy.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1192953466/saatgut-gruss-bienenwiese?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747812000",
		},
		{
			title: "Saatgut Gruß „Liebesblume“",
			image: "https://i.etsystatic.com/31340310/r/il/95413b/3776804908/il_570xN.3776804908_6a70.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1206914031/saatgut-gruss-liebesblume?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747854000",
		},
		{
			title: "Floraler Saatgut Gruß",
			image: "https://i.etsystatic.com/31340310/r/il/338e82/3776817118/il_570xN.3776817118_6a40.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1192953466/saatgut-gruss-bienenwiese?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747812000",
		},
		{
			title: "Anhänger Filzblume mit Karte",
			image: "https://i.etsystatic.com/31340310/r/il/ae9c12/3674017498/il_570xN.3674017498_6dki.jpg",
			url:   "https://www.etsy.com/de/kreatiVroni/listing/1206980547/anhanger-filzblume-mit-karte?utm_campaign=Share&utm_medium=social_organic&utm_source=MSMT&utm_term=so.smt&share_time=1648747995000",
		},
	}
	promoText := "Mit diesem Produkt kannst du nachhaltige Blumengrüße verschicken. Neben der liebevollen von Hand gestalteten Schachtel ist ein Döschen mit Saatgut für eine Bienenwiese bzw. einen Schmetterlingstreff enthalten. Dadurch bereiten die Blumen nicht nur eine kurze Freude wie bei einem Blumenstrauß, sondern es gibt das ganze Jahr über etwas zu bestaunen."
	promoTitle := "Blumengrüße 2.0"

	return app.Body().Class("bg-gradient-to-r from-green-200 to-green-500 p-0 pt-5 md:p-8 md:py-10").Body(
		&navbar{},
		app.Div().Class("pt-0 md:pt-20").Body(
			app.Figure().Class("bg-white bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-xl p-4 md:flex").Body(
				app.Div().Class("pt-6 md:p-8 text-center md:text-left space-y-4").Body(
					app.H2().Class("text-5xl font-bold title").Text(promoTitle),
					app.Blockquote().Class("pt-5").Body(
						app.P().Class("text-lg font-semibold").Text(
							promoText,
						),
					),
				),
			),
		),
		app.Div().Class("flex flex-wrap gap-8 justify-center items-center w-full pt-2").Body(
			app.Range(promoProducts).Slice(func(i int) app.UI {
				return app.Div().Class("flex max-w-md h-60 p-6 bg-white bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-xl").Body(
					app.Div().Class("flex-none w-44 relative").Body(
						app.Img().Class("absolute inset-0 w-full h-full object-cover rounded-lg").Src(promoProducts[i].image),
					),
					app.Div().Class("flex-auto max-h-full pl-6").Body(
						app.Div().Class("flex flex-wrap items-baseline").Body(
							app.H1().Class("w-full flex-none font-semibold mb-2.5 word break-words").Text(html.UnescapeString(promoProducts[i].title)),
						),
						app.Div().Class("flex space-x-3 mb-0 text-sm font-semibold").Body(
							app.Div().Class("flex-auto flex space-x-3").Body(
								app.A().Class("h-8 w-full flex items-center justify-center rounded-full text-black bg-pink-400").Href(promoProducts[i].url).Text("Zum Produkt"),
							),
						),
					),
				)
			}),
		),
	)
}

func (promo *promo) OnMount(ctx app.Context) {
	if ctx.AppUpdateAvailable {
		ctx.Reload()
	}
}
