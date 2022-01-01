package main

type order struct {
	ID           string `json:"id" firestore:"id,omitempty"`
	User       string `json:"user firestore:"user,omitempty"`
	Payment      string `json:"payment" firestore:"payment,omitempty"`
	ProductName  string `json:"productName" firestore:"productName,omitempty"`
	ProductId    string `json:"productId" firestore:"productId,omitempty"`
	Price        string `json:"price" firestore:"price,omitempty"`
	Email        string `json:"email" firestore:"email,omitempty"`
	Name         string `json:"name" firestore:"name,omitempty"`
	AddressLine1 string `json:"addressline1" firestore:"addressline1,omitempty"`
	AddressLine2 string `json:"addressline2" firestore:"addressline2,omitempty"`
	City         string `json:"city" firestore:"city,omitempty"`
}
