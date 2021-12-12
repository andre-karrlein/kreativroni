package main

type etsyProduct struct {
	Id    int    `json:"listing_id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type etsyProductData struct {
	Count   int           `json:"count"`
	Results []etsyProduct `json:"results"`
}

type productImage struct {
	Url string `json:"url_fullxfull"`
}

type etsyImageData struct {
	Count   int            `json:"count"`
	Results []productImage `json:"results"`
}

type product struct {
	Name  string
	Link  string
	Image string
}

type productData struct {
	Product []product `json:"products"`
}
