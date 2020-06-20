package models

import (
	"luakit/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

type UcenterOpenid struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid        int    `json:"cuid"`         // ucenter id
	PlatformKey string `json:"platform_key"` // 平台key
	Type        string `json:"type"`         // 类型 wechat ali app
	Openid      string `json:"openid"`
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterOpenid))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddUcenterOpenid(m *UcenterOpenid) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据cuid 和platformkey 获取openid
func GetUcenterOpenid(platform_key string, cuid int64) (v *UcenterOpenid, err error) {
	o := orm.NewOrm()
	v = &UcenterOpenid{}

	if err = o.QueryTable(new(UcenterOpenid)).Filter("PlatformKey", platform_key).Filter("Cuid", cuid).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
