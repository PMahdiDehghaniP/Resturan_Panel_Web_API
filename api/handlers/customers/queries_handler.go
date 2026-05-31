package customers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/data/db"
	customers "github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/data/sql-scripts/queries/cutomers"
	logging "github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/logger"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/models"
	"github.com/gin-gonic/gin"
)

var logger = logging.NewLogger(config.GetConfig())

func GetAllCustomers(c *gin.Context) {
	var dbClient = db.GetPostgresDB()
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	var (
		total      int64
		countQuery = "SELECT count(*) FROM customer"
	)
	if err := dbClient.QueryRow(countQuery).Scan(&total); err != nil {
		logger.ErrorF("Error getting customers: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting customers count",
		})
		return
	}
	rows, err := dbClient.Query(customers.GET_ALL_CUSTOMERS, pageSize, offset)
	defer rows.Close()
	if err != nil {
		logger.ErrorF("Error getting customers: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting customers count",
		})
		return
	}
	customersList := make([]models.Customer, 0)
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(
			&customer.CustomerId,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.PhoneNumber,
			&customer.CustomerType,
			&customer.JoinDate,
		); err != nil {
			logger.ErrorF("Error scanning customers: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error scanning customers",
			})
			return
		}
		customersList = append(customersList, customer)
	}
	if err := rows.Err(); err != nil {
		logger.ErrorF("Error reading customers rows: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading customers",
		})
		return
	}
	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))
	response := models.CustomersListResponse{
		Data: customersList,
		Pagination: models.Pagination{
			PageIndex: page,
			PageSize:  pageSize,
			Total:     int(total),
			TotalPage: totalPage,
		},
	}
	c.JSON(http.StatusOK, response)
}
