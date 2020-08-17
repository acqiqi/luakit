package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type SystemConfig struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	K           string `json:"k"`            // key
	V           string `json:"v"`            // value
	Title       string `json:"title"`        // 标题
	Describe    string `json:"describe"`     // 描述
	PlatformKey string `json:"platform_key"` // platform_key  如果全局写ALL
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SystemConfig))
}

// 获取系统配置的value
func GetSystemConfigValue(platform_key string, k string) (val string) {
	o := orm.NewOrm()
	v := &SystemConfig{}
	if err := o.QueryTable(new(SystemConfig)).Filter("PlatformKey", platform_key).Filter("K", k).Filter("Flag", 1).RelatedSel().One(v); err == nil {
		return v.V
	}
	return ""
}
