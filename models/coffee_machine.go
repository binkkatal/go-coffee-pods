package models

// CoffeeMachine represents a single Coffee Machine product.
type CoffeeMachine struct {
	CoffeeMachineID int    `json:"coffee_machine_id,omitempty"`
	SizeID          int    `json:"size_id,omitempty"`
	SizeName        string `json:"size_name,omitempty"`
	SKU             string `json:"sku,omitempty"`
	ModelID         int    `json:"model_id,omitempty"`
	ModelName       string `json:"model_name,omitempty"`
	WaterLine       bool   `json:"water_line,omitempty"`
}
