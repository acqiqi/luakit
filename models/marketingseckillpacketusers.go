package models

import (
	"luakit/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type MarketingSeckillPacketUsers struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid                int    `json:"cuid"`
	PlatformKey         string `json:"platform_key"`
	SeckillPacketLinkId int    `json:"seckill_packet_link_id"`
	SeckillPacketId     int    `json:"seckill_packet_id"`
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(MarketingSeckillPacketUsers))
}

//根据cuid  seckillpacket id 获取
func GetMarketingSeckillPacketUsersOrUser(cuid int64, seckill_packet_id int64) (v *MarketingSeckillPacketUsers, err error) {
	o := orm.NewOrm()
	v = &MarketingSeckillPacketUsers{}
	if err = o.QueryTable(new(MarketingSeckillPacketUsers)).Filter("Cuid", cuid).Filter("SeckillPacketId", seckill_packet_id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func AddMarketingSeckillPacketUsers(m *MarketingSeckillPacketUsers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
