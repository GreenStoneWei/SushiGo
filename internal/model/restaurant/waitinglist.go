package restaurant

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	schema "stone.com/sushigo/internal/schema"
	model "stone.com/sushigo/internal/schema/model"
)

func GetLastWaitingListByRestaurantId(restaurantId int, userId int) (int64, int) { // *model.WaitingList
	db := schema.GetDBInstance()
	dateString := getDateString()

	var waitinglist model.WaitingList
	err := db.QueryTable("waiting_list").Filter("restaurant_id", restaurantId).Filter("date", dateString).OrderBy("-Id").One(&waitinglist) // orderBy -col means DESC

	var nextWaitingNumber int
	var waitingListId int64

	if err == orm.ErrNoRows {
		fmt.Println("Not row found")
		waitinglist := model.WaitingList{
			UserId:       userId,
			RestaurantId: restaurantId,
			Date:         dateString,
			Number:       1,
			WaitingAt:    time.Now(),
		}
		id, err := db.Insert(&waitinglist)
		if err != nil {
			panic(err)
		}
		waitingListId = id
	} else {
		waitinglist := model.WaitingList{
			UserId:       userId,
			RestaurantId: restaurantId,
			Date:         dateString,
			Number:       waitinglist.Number + 1,
			WaitingAt:    time.Now(),
		}
		id, err := db.Insert(&waitinglist)
		if err != nil {
			panic(err)
		}
		waitingListId = id
	}
	return waitingListId, nextWaitingNumber
}

func getDateString() string {
	now := time.Now().Format("2006-01-02")
	return now
}
