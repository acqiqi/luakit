package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type SmUsers struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid            int     `json:"cuid"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
	IsSetArea       int     `json:"is_set_area"`
	AreaName        string  `json:"area_name"`          // 当前区域的名称
	AreaLevel       int     `json:"area_level"`         // 0省1市 2区 3小区
	AddreaaCheckId  int     `json:"addreaa_check_id"`   // 对应的id city 或者housingid
	IsNewUser       int     `json:"is_new_user"`        // 是否是当前端新用户
	IsGjAuth        int     `json:"is_gj_auth"`         // 是否工匠认证 0否 1认证中 2已认证 -1驳回
	GjIdcardTop     string  `json:"gj_idcard_top"`      // 身份证正面
	GjIdcardBom     string  `json:"gj_idcard_bom"`      // 身份证背面
	GjIdcardId      string  `json:"gj_idcard_id"`       // 身份证号码
	GjIdcardEndTime string  `json:"gj_idcard_end_time"` // 到期时间
	GjIdcardAddress string  `json:"gj_idcard_address"`  // 身份证地址
	GjName          string  `json:"gj_name"`            // 认证姓名
	GjAvatar        string  `json:"gj_avatar"`          // 工匠头像
	GjProvinceId    int     `json:"gj_province_id"`     // 省id
	GjCityId        int     `json:"gj_city_id"`         // 市id
	GjAreaId        int     `json:"gj_area_id"`         // 区id
	GjAreaLevel     int     `json:"gj_area_level"`      // 区域等级 0省 1市 2区
	GjMobile        string  `json:"gj_mobile"`          // 工匠认证手机号
	GjDesc          string  `json:"gj_desc"`            // 工匠介绍
	GjSfCuid        int     `json:"gj_sf_cuid"`         // 师傅id
	IsSf            int     `json:"is_sf"`              // 是否师父
	GjErrMsg        string  `json:"gj_err_msg"`         // 认证失败文案
	SfTime          string  `json:"sf_time"`            // 出师日期
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SmUsers))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddSmUsers(m *SmUsers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
