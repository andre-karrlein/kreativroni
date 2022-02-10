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
