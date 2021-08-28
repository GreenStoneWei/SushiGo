package waitinglist

import (
	"time"

	"github.com/astaxie/beego/orm"
	schema "stone.com/sushigo/internal/schema"
	model "stone.com/sushigo/internal/schema/model"
)

func PatchWaitingListById(waitingListId int, reqBody PatchWaitingListRequest) {
	validatePatchReq(reqBody)

	db := schema.GetDBInstance()

	var waitinglist model.WaitingList
	err := db.QueryTable("waiting_list").Filter("Id", waitingListId).One(&waitinglist)
	if err == orm.ErrNoRows {
		panic("Waiting List Not Found")
	}

	if !waitinglist.CancelAt.IsZero() {
		panic("The waiting list is canceled already")
	}

	if reqBody.CheckinAt != "" {
		if !waitinglist.FinishAt.IsZero() {
			panic("Cannot checkin after finish")
		}
		waitinglist.CheckInAt = time.Now()
	}

	if reqBody.FinishAt != "" {
		if waitinglist.CheckInAt.IsZero() {
			panic("Cannot finish if didn't checkin")
		}
		waitinglist.FinishAt = time.Now()
	}

	if reqBody.CancelAt != "" {
		waitinglist.CancelAt = time.Now()
	}

	db.Update(&waitinglist)
}

func validatePatchReq(reqBody PatchWaitingListRequest) {
	reqTimestamps := []string{reqBody.CancelAt, reqBody.CheckinAt, reqBody.FinishAt}

	notNilKeys := 0
	for i := range reqTimestamps {
		if len(reqTimestamps[i]) > 1 {
			notNilKeys++
		}
	}

	if notNilKeys > 1 {
		panic("err")
	}
}

type PatchWaitingListRequest struct {
	CancelAt  string `json:"checkinAt,omitempty"`
	CheckinAt string `json:"cancelAt,omitempty"`
	FinishAt  string `json:"finishAt,omitempty"`
}
