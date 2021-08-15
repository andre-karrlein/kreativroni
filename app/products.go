package main

type product struct {
	Name  string `firestore:"name,omitempty"`
	Link  string `firestore:"link,omitempty"`
	Image string `firestore:"image,omitempty"`
}

func getProducts() []product {
	products := []product{
		{"Product Name 1", "#", "https://storage.googleapis.com/kreativroni-web/Logo1.PNG"},
		{"Product Name 2", "#", "https://storage.googleapis.com/kreativroni-web/Logo1.PNG"},
		{"Product Name 3", "#", "https://storage.googleapis.com/kreativroni-web/Logo1.PNG"},
		{"Product Name 4", "#", "https://storage.googleapis.com/kreativroni-web/Logo1.PNG"},
		{"Product Name 5", "#", "https://storage.googleapis.com/kreativroni-web/Logo1.PNG"},
	}

	return products
}
