package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type UcenterPriceLog struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid        int     `json:"cuid"`
	Price       float64 `json:"price"`        // 金额
	PlatformKey string  `json:"platform_key"` // 平台key
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterPriceLog))
}

func AddUcenterPriceLog(m *UcenterPriceLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
