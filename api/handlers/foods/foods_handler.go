package foods

import (
	"net/http"

	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/data/db"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/data/sql-scripts/queries/foods"
	logging "github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/logger"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/models"
	"github.com/gin-gonic/gin"
)

var logger = logging.NewLogger(config.GetConfig())

func GetAllFoods(c *gin.Context) {
	dbClient := db.GetPostgresDB()
	foodCategory := c.DefaultQuery("foodCategory", "")
	foodSqlScript := foods.GET_ALL_FOOD
	args := []interface{}{}
	if foodCategory != "" {
		foodSqlScript += " WHERE fc.category_name = $1"
		args = append(args, foodCategory)
	}
	rows, err := dbClient.Query(foodSqlScript, args...)
	defer rows.Close()
	if err != nil {
		logger.ErrorF("Error getting foods from database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting foods",
		})
		return
	}
	foodsList := make([]models.FoodItem, 0)
	for rows.Next() {
		var foodItem models.FoodItem
		if err = rows.Scan(
			&foodItem.FoodName,
			&foodItem.Description,
			&foodItem.Calories,
			&foodItem.ImagePath,
			&foodItem.UnitPrice,
			&foodItem.Inventory,
			&foodItem.FoodCategory); err != nil {
			logger.ErrorF("Error getting foods from database: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error getting foods",
			})
			return
		}
		foodsList = append(foodsList, foodItem)
	}
	response := models.FoodItemsListResponse{Data: foodsList}
	c.JSON(http.StatusOK, gin.H{
		"foods": response,
	})

}
