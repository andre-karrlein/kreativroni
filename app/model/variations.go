package model

type Variation struct {
	Id      int    `json:"id"`
	ValueId int    `json:"value_id"`
	Value   string `json:"value"`
	ImageId int    `json:"image_id"`
}

type ListingVariation struct {
	Id      int    `json:"property_id"`
	ValueId int    `json:"value_id"`
	Value   string `json:"value"`
	ImageId int    `json:"image_id"`
}

type VariationData struct {
	Count   int                `json:"count"`
	Results []ListingVariation `json:"results"`
}
