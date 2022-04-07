package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/andre-karrlein/kreativroni/app/model"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type shop struct {
	app.Compo

	products []model.Product
}

func (shop *shop) OnNav(ctx app.Context) {
	// Launching a new goroutine:
	ctx.Async(func() {
		app_key := app.Getenv("PRODUCTS_KEY")
		r, err := http.Get("https://kreativroni.de/api/v1/product?lang=de&appkey=" + app_key)
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

		var products []model.Product
		json.Unmarshal([]byte(sb), &products)

		shop.products = products
		shop.Update()
	})
}

func (shop *shop) Render() app.UI {
	return app.Div().Class("bg-gradient-to-r from-green-200 to-green-500 p-0  pt-5 md:p-8 md:py-10").Body(
		app.Div().Body(
			&navbar{},
		),
		app.Div().Class("flex flex-wrap gap-8 justify-center items-center min-h-screen w-full").Body(
			app.Range(shop.products).Slice(func(i int) app.UI {
				price := fmt.Sprintf("%.2f", (float64(shop.products[i].Price.Amount) / float64(shop.products[i].Price.Divisor)))
				price = strings.Replace(price, ".", ",", -1)

				return app.Div().Class("flex max-w-md h-60 p-6 bg-white bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-xl").Body(
					app.Div().Class("flex-none w-44 relative").Body(
						app.Img().Class("absolute inset-0 w-full h-full object-cover rounded-lg").Src(shop.products[i].Images[0].Url_570xN),
					),
					app.Div().Class("flex-auto max-h-full pl-6").Body(
						app.Div().Class("flex flex-wrap items-baseline").Body(
							app.H1().Class("w-full flex-none font-semibold mb-2.5 word break-words").Text(html.UnescapeString(shop.products[i].Title)),
							app.Div().Class("w-full flex-none mt-2 order-1 text-3xl font-bold text-pink-400").Text(price+" "+strings.Replace(shop.products[i].Price.Currency_Code, "EUR", "€", 1)),
						),
						app.Div().Class("flex space-x-3 mb-0 text-sm font-semibold").Body(
							app.Div().Class("flex-auto flex space-x-3").Body(
								app.A().Class("h-8 w-full flex items-center justify-center rounded-full text-black bg-pink-400").Href(shop.products[i].Url).Text("Zum Produkt"),
							),
						),
					),
				)
			}),
		),
	)
}
