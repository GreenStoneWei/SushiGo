package restaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	restaurantModel "stone.com/sushigo/internal/model/restaurant"
)

func GetAllRestaurants(c *gin.Context) {
	result := restaurantModel.GetAllRestaurants()
	c.IndentedJSON(http.StatusOK, result)
}
