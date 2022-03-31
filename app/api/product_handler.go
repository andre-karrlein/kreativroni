package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/andre-karrlein/kreativroni/app/api/utils"
	"github.com/andre-karrlein/kreativroni/app/model"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]

	app_key := os.Getenv("APP_KEY")
	products_key := os.Getenv("PRODUCTS_KEY")

	if !ok || len(keys[0]) < 1 || (keys[0] != app_key && keys[0] != products_key) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keys, ok = r.URL.Query()["lang"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	language := keys[0]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	productsJSON, err := json.Marshal(loadProducts(language))
	if err != nil {
		log.Fatal(err)
	}

	w.Write(productsJSON)
}

func loadProducts(language string) []model.Product {
	b, err := utils.Etsy_request("https://openapi.etsy.com/v3/application/shops/31340310/listings/active?limit=100")
	if err != nil {
		log.Println(err)
		return nil
	}
	sb := string(b)

	var listings model.ListingData
	json.Unmarshal([]byte(sb), &listings)

	var ids []string

	for _, listing := range listings.Results {
		ids = append(ids, strconv.Itoa(listing.Id))
	}

	id_string := strings.Join(ids[:], ",")

	url := "https://openapi.etsy.com/v3/application/listings/batch?limit=100&includes=images&language=" + language + "&listing_ids=" + id_string
	b, err = utils.Etsy_request(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	sb = string(b)

	var etsyData model.EtsyProductData
	json.Unmarshal([]byte(sb), &etsyData)

	var products []model.Product
	for _, listingProduct := range etsyData.Results {
		products = append(products, model.Product(listingProduct))
	}

	return products
}
