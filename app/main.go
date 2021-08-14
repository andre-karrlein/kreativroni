package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.Route("/products", &products{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "kreativroni.de",
		Title:       "KreatiVroni",
		Description: "Kreatives von Vroni",
		Icon: app.Icon{
			Default:    "/web/images/ak_logo_192.png", // Specify default favicon.
			Large:      "/web/images/ak_logo_512.png", // Specify large favicon
			AppleTouch: "/web/images/ak_logo_192.png", // Specify icon on IOS devices.
		},
		Styles: []string{
			//"https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
			"/web/css/main.css",
		},
		ThemeColor: "#3a8277",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
