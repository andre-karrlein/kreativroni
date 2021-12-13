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
	Url_full    string `json:"url_fullxfull"`
	Url_75x75   string `json:"url_75x75"`
	Url_170x135 string `json:"url_170x135"`
	Url_570xN   string `json:"url_570xN"`
}

type productData struct {
	Product []product `json:"products"`
}
