package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type MarketingCardMoney struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Name         string  `json:"name"`          // 名称
	Desc         string  `json:"desc"`          // 描述
	Logo         string  `json:"logo"`          // Logo
	Background   string  `json:"background"`    // 背景
	Price        float64 `json:"price"`         // 面值
	Status       int     `json:"status"`        // 状态 0正常 1使用 -1删除
	Cuid         int     `json:"cuid"`          // 用户id
	CardNo       string  `json:"card_no"`       // 卡号
	CardPassword string  `json:"card_password"` // 卡号密码
	PartnerId    int     `json:"partner_id"`
	EndTime      int     `json:"end_time"` // 到期时间 0永久
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(MarketingCardMoney))
}

// 根据合伙人id 获取全部卡券
func GetMarketingCardMoneyAllList(partner_id int64) (row int64, v *[]MarketingCardMoney, err error) {
	o := orm.NewOrm()
	v = &[]MarketingCardMoney{}
	if row, err = o.QueryTable(new(MarketingCardMoney)).Filter("PartnerId", partner_id).Filter("Status", 0).Filter("flag", 1).RelatedSel().All(v); err == nil {
		return row, v, nil
	}
	return 0, nil, err
}
