package models

import (
	"luakit/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type UcenterAddress struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	AdrName        string  `orm:"adr_name"`      // 选择地址名称
	AdrLatitude    float64 `orm:"adr_latitude"`  // 维度
	AdrLongitude   float64 `orm:"adr_longitude"` // 经度
	Address        string  `orm:"address"`       // 详细门牌号
	Cuid           int     `orm:"cuid"`
	IsDefault      int     `orm:"is_default"`       // 是否默认
	Name           string  `orm:"name"`             // 姓名
	Mobile         string  `orm:"mobile"`           // 手机号
	AreaLevel      int     `orm:"area_level"`       // 0省1市 2区 3小区
	AddreaaCheckId int     `orm:"addreaa_check_id"` // 对应的id city 或者housingid
	PlatformKey    string  `orm:"platform_key"`     // 平台key
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterAddress))
}

func GetUcenterAddressOrId(id int64, cuid int64) (v *UcenterAddress, err error) {
	o := orm.NewOrm()
	v = &UcenterAddress{}

	if err = o.QueryTable(new(UcenterAddress)).Filter("Id", id).Filter("Cuid", cuid).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
