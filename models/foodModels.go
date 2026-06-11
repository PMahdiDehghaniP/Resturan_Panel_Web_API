package models

type FoodItem struct {
	FoodName     string
	FoodCategory string
	Description  string
	Calories     float64
	ImagePath    string
	UnitPrice    float64
	Inventory    float64
}

type FoodItemsListResponse struct {
	Data []FoodItem
}
