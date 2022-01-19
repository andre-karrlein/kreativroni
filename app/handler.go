package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]

	app_key := os.Getenv("APP_KEY")

	if !ok || len(keys[0]) < 1 || keys[0] != app_key {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	productsJSON, err := json.Marshal(loadProducts())
	if err != nil {
		log.Fatal(err)
	}

	w.Write(productsJSON)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]

	app_key := os.Getenv("APP_KEY")

	if !ok || len(keys[0]) < 1 || keys[0] != app_key {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	newsJson, err := json.Marshal(loadNews())
	if err != nil {
		log.Fatal(err)
	}

	w.Write(newsJson)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]
	app_key := os.Getenv("APP_KEY")

	if !ok || len(keys[0]) < 1 || keys[0] != app_key {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method == "POST" {
		// Declare a new order struct.
		var o order

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = saveOrder(o)
		if err != nil {
			http.Error(w, "Error creating order", http.StatusBadGateway)
		}

		w.WriteHeader(http.StatusCreated)
		return
	}
	if r.Method == "GET" {
		keys_id, ok := r.URL.Query()["id"]
		if !ok || len(keys_id[0]) < 1 {
			allOrdersJSON, err := json.Marshal(loadAllOrders())
			if err != nil {
				log.Fatal(err)
			}

			w.Write(allOrdersJSON)
			w.WriteHeader(http.StatusOK)
			return
		}

		id := keys_id[0]

		ordersJSON, err := json.Marshal(loadOrders(id))
		if err != nil {
			log.Fatal(err)
		}

		w.Write(ordersJSON)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func loadProducts() []product {
	b, err := etsy_request("https://openapi.etsy.com/v3/application/shops/31340310/listings/active")
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

	url := "https://openapi.etsy.com/v3/application/listings/batch?language=de&includes=images&listing_ids=" + id_string
	b, err = etsy_request(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	sb = string(b)

	var etsyData etsyProductData
	json.Unmarshal([]byte(sb), &etsyData)

	var products []product
	for _, listingProduct := range etsyData.Results {
		products = append(products, product(listingProduct))
	}

	return products
}

func etsy_request(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	api_key := os.Getenv("API_KEY")
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

func saveOrder(order order) error {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.Collection("orders").Doc(order.ID+"--"+order.ProductId).Set(ctx, order)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func loadOrders(id string) []order {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var orders []order

	iter := client.Collection("orders").Where("user", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var o order
		doc.DataTo(&o)

		orders = append(orders, o)
	}

	return orders
}

func loadAllOrders() []order {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var orders []order

	iter := client.Collection("orders").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var o order
		doc.DataTo(&o)

		orders = append(orders, o)
	}

	return orders
}

func loadNews() []news {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var news_items []news

	iter := client.Collection("news").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var o news
		doc.DataTo(&o)

		news_items = append(news_items, o)
	}

	return news_items
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "kreativroni" // os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
