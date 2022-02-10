package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andre-karrlein/kreativroni/app/api"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.Route("/shop", &shop{})
	app.Route("/about_me", &about{})
	app.Route("/legal_notice", &legal_notice{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "kreativroni.de",
		Title:       "KreatiVroni",
		Description: "Kreatives von Vroni",
		Icon: app.Icon{
			Default:    "/web/images/logo_192.png", // Specify default favicon.
			Large:      "/web/images/logo_512.png", // Specify large favicon
			AppleTouch: "/web/images/logo_192.png", // Specify icon on IOS devices.
		},
		Styles: []string{
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
			"/web/css/main.css",
			"https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css",
		},
		ThemeColor: "#3a8277",
		Env: app.Environment{
			"PRODUCTS_KEY": os.Getenv("PRODUCTS_KEY"),
		},
	})

	http.HandleFunc("/api/v1/product", api.ProductsHandler)
	http.HandleFunc("/api/v1/order", api.OrderHandler)
	http.HandleFunc("/api/v1/news", api.NewsHandler)
	http.HandleFunc("/api/v1/sections", api.SectionsHandler)
	http.HandleFunc("/api/v1/customer", api.CustomerHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
