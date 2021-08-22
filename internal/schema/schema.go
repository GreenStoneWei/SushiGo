package schema

import (
	model "stone.com/sushigo/internal/schema/model"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	// 設定預設資料庫
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=gogogo password=gogogo host=127.0.0.1 port=5434 dbname=sushigo sslmode=disable")
	// 註冊定義的 model
	orm.RegisterModel(new(model.User), new(model.Restaurant), new(model.WaitingList))
	// 建立 table
	orm.RunSyncdb("default", false, true)

}

func GetDBInstance() orm.Ormer {
	o := orm.NewOrm()
	return o
}
