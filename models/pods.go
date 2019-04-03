package models

//Coffee Pod represent the single coffee pod product
type Pod struct {
	PodID      int    `json:"pod_id,omitempty"`
	SizeID     int    `json:"size_id,omitempty"`
	SizeName   string `json:"size_name,omitempty"`
	FlavorID   int    `json:"flavor_id,omitempty"`
	FlavorName string `json:"flavor_name,omitempty"`
	SKU        string `json:"sku,omitempty"`
	Quantity   int    `json:"quantity,omitempty"`
}
