package utils

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func CreateClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "kreativroni" // os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func DeleteFromCollection(id string, collection string) error {
	ctx := context.Background()
	client := CreateClient(ctx)
	defer client.Close()

	_, err := client.Collection(collection).Doc(id).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	return nil
}
