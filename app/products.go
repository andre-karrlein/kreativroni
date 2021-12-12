package main

type listing struct {
	Id int `json:"listing_id"`
}
type product struct {
	Id     int            `json:"listing_id"`
	Title  string         `json:"title"`
	Url    string         `json:"url"`
	Images []productImage `json:"images"`
}

type etsyProductData struct {
	Count   int       `json:"count"`
	Results []product `json:"results"`
}

type listingData struct {
	Count   int       `json:"count"`
	Results []listing `json:"results"`
}

type productImage struct {
	Url string `json:"url_fullxfull"`
}

type productData struct {
	Product []product `json:"products"`
}
