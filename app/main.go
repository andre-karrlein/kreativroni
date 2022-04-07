package main

import (
	"log"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.Route("/shop", &shop{})
	app.Route("/about_me", &about{})
	app.Route("/legal_notice", &legal_notice{})
	app.Route("/aktion", &promo{})

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
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
		RawHeaders: []string{
			"<meta name='apple-itunes-app' content='app-id=1601515699, app-argument=myURL'>",
		},
		Env: app.Environment{
			"PRODUCTS_KEY": os.Getenv("PRODUCTS_KEY"),
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
