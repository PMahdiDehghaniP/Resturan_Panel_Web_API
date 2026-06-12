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

type CreateFoodItemRequest struct {
	FoodName     string  `json:"foodName" binding:"required"`
	Description  string  `json:"description"`
	Calories     int     `json:"calories"`
	ImagePath    string  `json:"imagePath"`
	UnitPrice    float64 `json:"unitPrice" binding:"required"`
	Inventory    int     `json:"inventory" binding:"required"`
	FoodCategory int     `json:"foodCategory" binding:"required"`
}
type CreateFoodItemResponse struct {
	FoodId  int64
	Message string
}
type FoodItemsListResponse struct {
	Data []FoodItem
}
