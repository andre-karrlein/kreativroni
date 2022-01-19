package main

type news struct {
	ID      string `json:"id" firestore:"id,omitempty"`
	Title   string `json:"title" firestore:"title,omitempty"`
	Message string `json:"message" firestore:"message,omitempty"`
}
