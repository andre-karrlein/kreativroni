package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	productsJSON, err := json.Marshal(productData{
		Product: loadProducts(),
	})
	if err != nil {
		log.Fatal(err)
	}

	w.Write(productsJSON)
}

func loadProducts() []product {
	b, err := etsy_request("/listings/active")
	if err != nil {
		log.Println(err)
		return nil
	}
	sb := string(b)

	var listings listingData
	json.Unmarshal([]byte(sb), &listings)

	var ids []string

	for _, listing := range listings.Results {
		ids = append(ids, strconv.Itoa(listing.Id))
	}

	id_string := strings.Join(ids[:], ",")

	url := "https://openapi.etsy.com/v3/application/listings/batch?includes=images&listing_ids=" + id_string
	b, err = etsy_request(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	sb = string(b)

	var etsyData etsyProductData
	json.Unmarshal([]byte(sb), &etsyData)

	return etsyData.Results
}

func etsy_request(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://openapi.etsy.com/v3/application/shops/31340310"+path, nil)
	if err != nil {
		return nil, err
	}

	api_key := os.Getenv("API-KEY")
	req.Header.Add("x-api-key", api_key)

	lowerCaseHeader := make(http.Header)

	for key, value := range req.Header {
		lowerCaseHeader[strings.ToLower(key)] = value
	}

	req.Header = lowerCaseHeader
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
