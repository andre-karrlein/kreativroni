package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/andre-karrlein/kreativroni/app/api/utils"
	"github.com/andre-karrlein/kreativroni/app/model"
	"google.golang.org/api/iterator"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]
	app_key := os.Getenv("APP_KEY")

	if !ok || len(keys[0]) < 1 || keys[0] != app_key {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method == "POST" {
		// Declare a new order struct.
		var o model.Order

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

func saveOrder(order model.Order) error {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("orders").Doc(order.ID+"--"+order.ProductId).Set(ctx, order)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func loadOrders(id string) []model.Order {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	var orders []model.Order

	iter := client.Collection("orders").Where("user", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var o model.Order
		doc.DataTo(&o)

		orders = append(orders, o)
	}

	return orders
}

func loadAllOrders() []model.Order {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	var orders []model.Order

	iter := client.Collection("orders").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var o model.Order
		doc.DataTo(&o)

		orders = append(orders, o)
	}

	return orders
}
