package main

type product struct {
	Name  string `firestore:"name,omitempty"`
	Link  string `firestore:"link,omitempty"`
	Image string `firestore:"image,omitempty"`
}

type productWithId struct {
	ID      string  `json:"id"`
	Product product `json:"product"`
}
