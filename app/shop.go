package main

import (
	"encoding/json"
	"html"
	"io/ioutil"
	"net/http"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type shop struct {
	app.Compo

	products []product
}

func (shop *shop) OnNav(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		r, err := http.Get("/api/v1/product")
		if err != nil {
			app.Log(err)
			return
		}
		defer r.Body.Close()

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			app.Log(err)
			return
		}

		sb := string(b)

		var productsData productData
		json.Unmarshal([]byte(sb), &productsData)

		var products []product

		products = append(products, productsData.Product...)

		shop.products = products
		shop.Update()
	})
}

func (shop *shop) Render() app.UI {
	return app.Body().Class("bg-gradient-to-r from-green-200 to-green-500 p-0  pt-5 md:p-8 md:py-10").Body(
		app.Div().Body(
			&navbar{},
		),
		app.Div().Class("flex flex-wrap gap-8 justify-center items-center min-h-screen w-full").Body(
			app.Range(shop.products).Slice(func(i int) app.UI {
				return app.Div().Class("flex max-w-md h-60 p-6 bg-white bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-xl").Body(
					app.Div().Class("flex-none w-44 relative").Body(
						app.Img().Class("absolute inset-0 w-full h-full object-cover rounded-lg").Src(shop.products[i].Images[0].Url_170x135),
					),
					app.Div().Class("flex-auto max-h-full pl-6").Body(
						app.Div().Class("flex flex-wrap items-baseline").Body(
							app.H1().Class("w-full flex-none font-semibold mb-2.5 word break-words").Text(html.UnescapeString(shop.products[i].Title)),
						),
						app.Div().Class("flex space-x-3 mb-0 text-sm font-semibold").Body(
							app.Div().Class("flex-auto flex space-x-3").Body(
								app.A().Class("h-8 w-full flex items-center justify-center rounded-full text-black bg-pink-300").Href(shop.products[i].Url).Text("Zum Produkt"),
							),
						),
					),
				)
			}),
		),
	)
}
