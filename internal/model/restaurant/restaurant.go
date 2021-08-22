package restaurant

import (
	"fmt"

	schema "stone.com/sushigo/internal/schema"
	model "stone.com/sushigo/internal/schema/model"
)

func GetAllRestaurants() []*model.Restaurant {
	db := schema.GetDBInstance()
	var restaurants []*model.Restaurant
	// fmt.Printf("Returned Rows Num: %s, %s", test)
	row, err := db.QueryTable("restaurant").All(&restaurants, "Id", "WaitingLimit", "IsWaitlineOpen")
	fmt.Printf("Returned Rows Num: %s, %s", row, err)
	return restaurants
}
