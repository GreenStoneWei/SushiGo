package restaurant

import (
	schema "stone.com/sushigo/internal/schema"
	model "stone.com/sushigo/internal/schema/model"
)

func GetAllRestaurants() []*model.Restaurant {
	db := schema.GetDBInstance()
	var restaurants []*model.Restaurant
	db.QueryTable("restaurant").All(&restaurants, "Id", "WaitingLimit", "IsWaitlineOpen")
	return restaurants
}

func GetRestaurantById(restaurantId int) model.Restaurant {
	db := schema.GetDBInstance()
	var restaurant model.Restaurant
	db.QueryTable("restaurant").Filter("Id", restaurantId).One(&restaurant)
	return restaurant
}
