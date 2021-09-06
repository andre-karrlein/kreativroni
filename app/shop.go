package main

import (
	"encoding/json"
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

		var productsData []productWithId
		json.Unmarshal([]byte(sb), &productsData)

		var products []product

		for _, product := range productsData {
			products = append(products, product.Product)
		}

		shop.products = products
		shop.Update()
	})
}

func (shop *shop) Render() app.UI {
	return app.Body().Class("shop-body").Body(
		app.Div().Class("shop-header").Body(
			&navbar{},
		),
		app.Div().Class("container").Body(
			app.Range(shop.products).Slice(func(i int) app.UI {
				return app.Div().Class("card").Body(
					app.Div().Class("imgBx").Body(
						app.Img().Src(shop.products[i].Image),
						app.H2().Text(shop.products[i].Name),
					),
					app.Div().Class("content").Body(
						app.A().Href(shop.products[i].Link).Text("Buy Now"),
					),
				)
			}),
		),
	)
}
