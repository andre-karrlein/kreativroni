package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func saveHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	auth := strings.Split(authHeader, " ")
	password := auth[1]

	if !CheckAuth(password) {
		w.WriteHeader(http.StatusUnauthorized)
	}

	if r.Method == "POST" {
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		image := r.FormValue("image")
		link := r.FormValue("link")
		log.Println("Name: " + name + " Image: " + image + " Link: " + link)

		article := product{
			Name:  name,
			Link:  link,
			Image: image,
		}

		saveProduct(article)

		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		deleteHandler(w, r)
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

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := r.URL.Query()
	id, present := query["id"]
	if !present || len(id) == 0 {
		fmt.Println("id not present")
	}

	log.Println(id[0])

	deleteProduct(id[0])

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func loadProducts() []productWithId {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var products []productWithId

	iter := client.Collection("products").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var p product
		id := doc.Ref.ID
		doc.DataTo(&p)

		productStructure := productWithId{ID: id, Product: p}

		products = append(products, productStructure)
	}

	return products
}

func saveProduct(product product) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, _, err := client.Collection("products").Add(ctx, product)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

func deleteProduct(id string) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.Collection("products").Doc(id).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
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