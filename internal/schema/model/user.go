package schema

type User struct {
	Id          int
	FullName    string `orm:"size(100)"`
	PhoneNumber string `orm:"size(100);null"`
}
