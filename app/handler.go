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

func imageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		return
	}
	id := keys[0]
	w.WriteHeader(http.StatusOK)

	productsJSON, err := json.Marshal(loadImage(id))
	if err != nil {
		log.Fatal(err)
	}

	w.Write(productsJSON)
}

func loadProducts() []etsyProduct {
	b, err := etsy_request("/listings/active")
	if err != nil {
		log.Println(err)
		return nil
	}
	sb := string(b)

	var etsyProducts etsyProductData
	json.Unmarshal([]byte(sb), &etsyProducts)

	var products []etsyProduct

	for _, product := range etsyProducts.Results {
		products = append(products, product)
	}
	return products
}

func loadImage(id string) productImage {
	b, err := etsy_request("/listings/" + id + "/images")
	if err != nil {
		log.Println(err)
	}
	sb := string(b)

	var imageData etsyImageData
	json.Unmarshal([]byte(sb), &imageData)
	url := imageData.Results[0].Url

	id_int, _ := strconv.Atoi(id)

	return productImage{Id: id_int, Url: url}
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
