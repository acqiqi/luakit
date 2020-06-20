package models

import (
	"fmt"
	"luakit/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type MarketingSeckillPacket struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Type   int    `json:"type"`   // 0红包(不需要付费)
	Title  string `json:"title"`  // 标题
	Status int    `json:"status"` // 0未开始 1开始 2结束
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(MarketingSeckillPacket))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddMarketingSeckillPacket(m *MarketingSeckillPacket) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//根据状态直接获取已开始的红包
func GetMarketingSeckillPacketOrBegin() (v *MarketingSeckillPacket, err error) {
	o := orm.NewOrm()
	v = &MarketingSeckillPacket{}

	if err = o.QueryTable(new(MarketingSeckillPacket)).Filter("Status", 1).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdatePost updates Post by Id and returns error if
// the record to be updated doesn't exist
func UpdateMarketingSeckillPacketById(m *MarketingSeckillPacket) (err error) {
	o := orm.NewOrm()
	v := MarketingSeckillPacket{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
