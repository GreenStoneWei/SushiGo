package schema

type Restaurant struct {
	Id             int
	WaitingLimit   int
	IsWaitlineOpen bool
	CheckInNumber  int
	WaitingNumber  int
}
