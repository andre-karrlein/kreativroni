package main

type sectionListing struct {
	Id    int    `json:"shop_section_id"`
	Title string `json:"title"`
}

type section struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type sectionData struct {
	Count   int              `json:"count"`
	Results []sectionListing `json:"results"`
}
