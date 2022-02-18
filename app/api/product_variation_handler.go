package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

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
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		productsJSON, err := json.Marshal(loadAllVariations(language))
		if err != nil {
			log.Fatal(err)
		}

		w.Write(productsJSON)
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
	index := 1

	for _, listingVariation := range variationData.Results {
		variations = append(variations, model.Variation{
			Id:         index,
			PropertyId: listingVariation.Id,
			ValueId:    listingVariation.ValueId,
			Value:      listingVariation.Value,
			ImageId:    listingVariation.ImageId,
		})

		index++
	}

	return variations
}

func loadAllVariations(language string) []model.Variations {
	b, err := utils.Etsy_request("https://openapi.etsy.com/v3/application/shops/31340310/listings/active")
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

	var variations []model.Variations
	for _, id := range ids {
		b, err := utils.Etsy_request("https://openapi.etsy.com/v3/application/shops/31340310/listings/" + id + "/variation-images")
		if err != nil {
			log.Println(err)
			return nil
		}
		sb := string(b)

		var variationData model.VariationData
		json.Unmarshal([]byte(sb), &variationData)

		var variationElems []model.Variation
		index := 1

		for _, listingVariation := range variationData.Results {
			variationElems = append(variationElems, model.Variation{
				Id:         index,
				PropertyId: listingVariation.Id,
				ValueId:    listingVariation.ValueId,
				Value:      listingVariation.Value,
				ImageId:    listingVariation.ImageId,
			})

			index++
		}

		int_id, _ := strconv.Atoi(id)
		variations = append(variations, model.Variations{
			Id:         int_id,
			Variations: variationElems,
		})
	}

	return variations
}
