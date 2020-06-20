package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type UcenterCommission struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	PlatformKey   string `json:"platform_key"`    // 平台key
	AllShareScale int    `json:"all_share_scale"` // 分享最大比例
	AllOptScale   int    `json:"all_opt_scale"`   // 其他收益最大比例
	ShareLv1Scale int    `json:"share_lv1_scale"` // 分享一级
	ShareLv2Scale int    `json:"share_lv2_scale"` // 分享二级
	OptLv1Scale   int    `json:"opt_lv1_scale"`   // 其他收益一级
	OptLv2Scale   int    `json:"opt_lv2_scale"`   // 其他收益二级
	Status        int    `json:"status"`          // 0停用 1启用
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterCommission))
}

// 根据订单编号获取订单
func GetUcenterCommissionById(id int64) (v *UcenterCommission, err error) {
	o := orm.NewOrm()
	v = &UcenterCommission{}
	if err = o.QueryTable(new(UcenterCommission)).Filter("Id", id).Filter("Flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
