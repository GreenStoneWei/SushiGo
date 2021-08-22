package schema

import (
	"time"
)

type WaitingList struct {
	Id           int
	UserId       int
	RestaurantId bool
	Date         time.Time
	Number       int
	WaitingAt    time.Time `orm:"null"`
	CheckInAt    time.Time `orm:"null"`
	CancelAt     time.Time `orm:"null"`
	FinishAt     time.Time `orm:"null"`
}
