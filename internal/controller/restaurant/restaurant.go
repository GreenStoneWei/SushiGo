package restaurant

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	restaurantModel "stone.com/sushigo/internal/model/restaurant"
)

func GetAllRestaurants(c *gin.Context) {
	result := restaurantModel.GetAllRestaurants()
	c.IndentedJSON(http.StatusOK, result)
}

func PostWaitingList(c *gin.Context) {
	restaurantId := c.Param("id")
	fmt.Println("restaurantId %v", restaurantId)
	var requestBody struct{ userId string }
	err := c.BindJSON(&requestBody)
	if err != nil {
		return
	}
	userId := requestBody.userId
	fmt.Println("requestBody %v", requestBody)
	fmt.Println("userId %v", userId)
	waitingListId, number := restaurantModel.GetLastWaitingListByRestaurantId(restaurantId, userId)

	c.IndentedJSON(http.StatusOK, gin.H{waitingListId: waitingListId, number: number})
}
