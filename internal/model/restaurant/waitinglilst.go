package restaurant

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	schema "stone.com/sushigo/internal/schema"
	model "stone.com/sushigo/internal/schema/model"
)

func GetLastWaitingListByRestaurantId(restaurantId string, userId string) (string, string) { // *model.WaitingList
	db := schema.GetDBInstance()
	dateString := getDateString()
	fmt.Printf("restaurantId %v", restaurantId)
	fmt.Printf("userId %v", userId)
	var waitinglist model.WaitingList
	err := db.QueryTable("waiting_list").Filter("restaurant_id", restaurantId).Filter("date", dateString).One(&waitinglist)

	// var nextWaitingNumber = 0
	fmt.Printf("Returned waitinglist %p", waitinglist)
	if err == orm.ErrNoRows {
		// No result
		fmt.Printf("Not row found")
	}
	// if waitinglist {

	// }
	return "1", "2"
}

func getDateString() string {
	now := time.Now().Format("2006-01-02")
	fmt.Printf("Now: %s", now)
	return now
}
