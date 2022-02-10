package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/andre-karrlein/kreativroni/app/api/utils"
	"github.com/andre-karrlein/kreativroni/app/model"
)

func SectionsHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["appkey"]

	app_key := os.Getenv("APP_KEY")

	if !ok || len(keys[0]) < 1 || keys[0] != app_key {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keys, ok = r.URL.Query()["lang"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	language := keys[0]

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sectionsJSON, err := json.Marshal(loadSections(language))
	if err != nil {
		log.Fatal(err)
	}

	w.Write(sectionsJSON)
}

func loadSections(language string) []model.Section {
	b, err := utils.Etsy_request("https://openapi.etsy.com/v3/application/shops/31340310/sections?language=" + language)
	if err != nil {
		log.Println(err)
		return nil
	}
	sb := string(b)

	var sectionListings model.SectionData
	json.Unmarshal([]byte(sb), &sectionListings)

	var sections []model.Section
	for _, sectionListing := range sectionListings.Results {
		sections = append(sections, model.Section(sectionListing))
	}

	return sections
}
