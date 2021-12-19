package main

type listing struct {
	Id int `json:"listing_id"`
}

type listingProduct struct {
	Id          int            `json:"listing_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Tags        []string       `json:"tags"`
	Price       price          `json:"price"`
	Url         string         `json:"url"`
	Images      []productImage `json:"images"`
}
type product struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Tags        []string       `json:"tags"`
	Price       price          `json:"price"`
	Url         string         `json:"url"`
	Images      []productImage `json:"images"`
}

type productImage struct {
	Url_full    string `json:"url_fullxfull"`
	Url_75x75   string `json:"url_75x75"`
	Url_170x135 string `json:"url_170x135"`
	Url_570xN   string `json:"url_570xN"`
}

type price struct {
	Amount        int    `json:"amount"`
	Divisor       int    `json:"divisor"`
	Currency_Code string `json:"currency_code"`
}

type productData struct {
	Product []product `json:"products"`
}

type etsyProductData struct {
	Count   int              `json:"count"`
	Results []listingProduct `json:"results"`
}

type listingData struct {
	Count   int       `json:"count"`
	Results []listing `json:"results"`
}
