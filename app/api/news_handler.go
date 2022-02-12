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

func NewsHandler(w http.ResponseWriter, r *http.Request) {
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

		err := utils.DeleteFromCollection(keys[0], "news")
		if err != nil {
			http.Error(w, "Error deleting post", http.StatusBadGateway)
		}

		w.WriteHeader(http.StatusAccepted)
		return
	}
	if r.Method == "POST" {
		// Declare a new news struct.
		var n model.News

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = saveNews(n)
		if err != nil {
			http.Error(w, "Error creating or updating post", http.StatusBadGateway)
		}

		w.WriteHeader(http.StatusCreated)
		return
	}
	if r.Method == "GET" {
		keys_id, ok := r.URL.Query()["id"]
		if !ok || len(keys_id[0]) < 1 {
			newsJSON, err := json.Marshal(loadAllNews())
			if err != nil {
				log.Fatal(err)
			}

			w.Write(newsJSON)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

		id := keys_id[0]

		newsJSON, err := json.Marshal(loadNews(id))
		if err != nil {
			log.Fatal(err)
		}

		w.Write(newsJSON)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func loadAllNews() []model.News {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	var news_items []model.News

	iter := client.Collection("news").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var o model.News
		doc.DataTo(&o)

		news_items = append(news_items, o)
	}

	return news_items
}

func loadNews(id string) model.News {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	var news_item model.News

	iter := client.Collection("news").Where("id", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		var n model.News
		doc.DataTo(&n)

		news_item = n
	}

	return news_item
}

func saveNews(news model.News) error {
	ctx := context.Background()
	client := utils.CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection("news").Doc(news.ID).Set(ctx, news)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}
