package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/andre-karrlein/kreativroni/app/api/utils"
	"github.com/andre-karrlein/kreativroni/app/model"
)

func ProductVariationsHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]

	app_key := os.Getenv("APP_KEY")

	if !ok || len(keys[0]) < 1 || keys[0] != app_key {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keys, ok = r.URL.Query()["lang"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	language := keys[0]

	keys, ok = r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := keys[0]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	productsJSON, err := json.Marshal(loadVariations(language, id))
	if err != nil {
		log.Fatal(err)
	}

	w.Write(productsJSON)
}

func loadVariations(language string, id string) []model.Variation {
	b, err := utils.Etsy_request("https://openapi.etsy.com/v3/application/shops/31340310/listings/" + id + "/variation-images")
	if err != nil {
		log.Println(err)
		return nil
	}
	sb := string(b)

	var variationData model.VariationData
	json.Unmarshal([]byte(sb), &variationData)

	var variations []model.Variation

	for _, listingVariation := range variationData.Results {
		variations = append(variations, model.Variation(listingVariation))
	}

	return variations
}
