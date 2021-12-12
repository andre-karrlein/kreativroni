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

	var etsyProducts etsyProductData
	json.Unmarshal([]byte(sb), &etsyProducts)

	var products []product

	for _, etsy_product := range etsyProducts.Results {
		b, err = etsy_request("/listings/" + strconv.Itoa(etsy_product.Id) + "/images")
		if err != nil {
			log.Println(err)
			return nil
		}
		sb := string(b)

		var imageData etsyImageData
		json.Unmarshal([]byte(sb), &imageData)
		url := ""
		if (imageData.Count == 0) {
			url = imageData.Results[0].Url
		}

		products = append(products, product{
			Name:  etsy_product.Title,
			Link:  etsy_product.Url,
			Image: url,
		})
	}
	return products
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
