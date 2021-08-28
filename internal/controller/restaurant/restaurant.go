package restaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	restaurantModel "stone.com/sushigo/internal/model/restaurant"
)

func GetAllRestaurants(c *gin.Context) {
	result := restaurantModel.GetAllRestaurants()
	c.IndentedJSON(http.StatusOK, result)
}

func GetRestaurantById(c *gin.Context) {
	restaurantId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	result := restaurantModel.GetRestaurantById(restaurantId)
	c.IndentedJSON(http.StatusOK, result)
}

func PostWaitingList(c *gin.Context) {
	restaurantId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	var requestBody WaitingListRequest
	reqBodyErr := c.BindJSON(&requestBody)
	if reqBodyErr != nil {
		panic(reqBodyErr)
	}
	userId, err := strconv.Atoi(requestBody.UserId)

	waitingListId, number := restaurantModel.GetLastWaitingListByRestaurantId(restaurantId, userId)

	c.IndentedJSON(http.StatusOK, gin.H{"waitingListId": waitingListId, "number": number})
}

type WaitingListRequest struct {
	UserId string
}
