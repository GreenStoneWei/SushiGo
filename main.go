package main

import (
	"github.com/gin-gonic/gin"
	restaurantController "stone.com/sushigo/internal/controller/restaurant"

	// _ 讀取 init
	_ "stone.com/sushigo/internal/schema"
)

func main() {
	router := gin.Default()

	router.GET("/restaurants", restaurantController.GetAllRestaurants)
	router.GET("/restaurants/:id", restaurantController.GetRestaurantById)
	router.PATCH("/restaurants/:id", restaurantController.PatchRestaurantById)
	router.POST("/restaurants/:id/waitinglist", restaurantController.PostWaitingList)
	router.PATCH("/waitinglist/:id", restaurantController.PatchWaitingListById)
	router.Run("localhost:9595")
}
