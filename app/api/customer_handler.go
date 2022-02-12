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

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]

	app_key := os.Getenv("APP_KEY")

	if !ok || len(keys[0]) < 1 || keys[0] != app_key {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method == "DELETE" {
		keys, ok := r.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := utils.DeleteFromCollection(keys[0], "customer")
		if err != nil {
			http.Error(w, "Error deleting post", http.StatusBadGateway)
		}

		w.WriteHeader(http.StatusAccepted)
		return
	}
	if r.Method == "POST" {
		// Declare a new customer struct.
		var c model.Customer

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = saveCustomer(c)
		if err != nil {
			http.Error(w, "Error creating or updating customer", http.StatusBadGateway)
		}

		w.WriteHeader(http.StatusCreated)
		return
	}
	if r.Method == "GET" {
		keys_id, ok := r.URL.Query()["id"]
		if !ok || len(keys_id[0]) < 1 {
			customersJSON, err := json.Marshal(loadAllCustomers())
			if err != nil {
				log.Fatal(err)
			}

			w.Write(customersJSON)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

		id := keys_id[0]

		customerJSON, err := json.Marshal(loadCustomer(id))
		if err != nil {
			log.Fatal(err)
		}

		w.Write(customerJSON)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func loadAllCustomers() []model.Customer {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	var customers []model.Customer

	iter := client.Collection("customer").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var c model.Customer
		doc.DataTo(&c)

		customers = append(customers, c)
	}

	return customers
}

func loadCustomer(id string) model.Customer {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	var customer_item model.Customer

	iter := client.Collection("customer").Where("id", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var c model.Customer
		doc.DataTo(&c)

		customer_item = c
	}

	return customer_item
}

func saveCustomer(customer_item model.Customer) error {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("customer").Doc(customer_item.ID).Set(ctx, customer_item)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}
