package foods

import (
	"net/http"

	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/data/db"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/data/sql-scripts/mutations/foods"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/models"
	"github.com/gin-gonic/gin"
)

func HandleCreateFood(c *gin.Context) {
	dbClient := db.GetPostgresDB()
	var request = models.CreateFoodItemRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.ErrorF("Failed to bind json body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var foodId int64
	err := dbClient.QueryRow(foods.CREATE_FOOD, request.FoodName,
		request.Description,
		request.Calories,
		request.ImagePath,
		request.UnitPrice,
		request.Inventory,
		request.FoodCategory).Scan(&foodId)
	if err != nil {
		logger.ErrorF("Error creating food: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating food",
		})
		return
	}
	response := models.CreateFoodItemResponse{
		FoodId:  foodId,
		Message: "Created food",
	}
	c.JSON(http.StatusCreated, response)
}
