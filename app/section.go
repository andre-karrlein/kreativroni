package main

type section struct {
	Id    int    `json:"shop_section_id"`
	Title string `json:"title"`
}

type sectionData struct {
	Count   int       `json:"count"`
	Results []section `json:"results"`
}
