package restaurant

import (
	"strconv"

	"github.com/astaxie/beego/orm"
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
	err := db.QueryTable("restaurant").Filter("Id", restaurantId).One(&restaurant)
	if err == orm.ErrNoRows {
		panic("restaurantId" + strconv.Itoa(restaurantId) + " Not Found")
	}
	return restaurant
}

func PatchRestaurantById(restaurantId int, reqBody PatchRestaurantRequest) {
	db := schema.GetDBInstance()

	var restaurant model.Restaurant
	err := db.QueryTable("restaurant").Filter("Id", restaurantId).One(&restaurant)
	if err == orm.ErrNoRows {
		panic("restaurantId" + strconv.Itoa(restaurantId) + " Not Found")
	}

	restaurant.IsWaitlineOpen = reqBody.IsWaitlineOpen
	restaurant.WaitingLimit = reqBody.WaitingLimit

	db.Update(&restaurant)
}

type PatchRestaurantRequest struct {
	IsWaitlineOpen bool `json:"isWaitlineOpen"`
	WaitingLimit   int  `json:"waitingLimit"`
}
