package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type SmService struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Title           string  `json:"title"`              // 标题
	Logo            string  `json:"logo"`               // 列表logo
	Icon            string  `json:"icon"`               // 首页icon
	Banner          string  `json:"banner"`             // banner列表 json
	CatsId          int     `json:"cats_id"`            // 分类id
	Describe        string  `json:"describe"`           // 描述
	Content         string  `json:"content"`            // 富文本内容
	IsVideo         int     `json:"is_video"`           // 是否显示视频banner
	VideoUrl        string  `json:"video_url"`          // 视频url
	ShowNum         int     `json:"show_num"`           // 显示数量
	PayNum          int     `json:"pay_num"`            // 销售数量
	CollectNum      int     `json:"collect_num"`        // 收藏数量
	ShareNum        int     `json:"share_num"`          // 分享数量
	MinPrice        float64 `json:"min_price"`          // 最小售价（用于限定最低单品购买以及显示）
	MaxPrice        float64 `json:"max_price"`          // 最大售价，只用于显示
	MinCostPrice    float64 `json:"min_cost_price"`     // 最小销售原价
	MaxCostPrice    float64 `json:"max_cost_price"`     // 最高售价原价
	ProjectType     int     `json:"project_type"`       // 0正常 1团购
	IsShow          int     `json:"is_show"`            // 是否显示
	Status          int     `json:"status"`             // 0停售 1正常销售  10预售
	IsDelete        int     `json:"is_delete"`          // 是否远程删除
	IsTopic         int     `json:"is_topic"`           // 是否推荐
	IsNew           int     `json:"is_new"`             // 是否new
	Sort            int     `json:"sort"`               // 排序
	AreaType        int     `json:"area_type"`          // 限定区域范围等级 0全国 1省 2市 3区 4小区
	ProvinceId      int     `json:"province_id"`        // 省id
	CityId          int     `json:"city_id"`            // 市id
	AreaId          int     `json:"area_id"`            // 区id
	HousingId       int     `json:"housing_id"`         // 小区id
	IsVipPrice      int     `json:"is_vip_price"`       // 是否有vip折扣
	IsClientUseAppt int     `json:"is_client_use_appt"` // 判断是否需要用户端选择服务区间
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SmService))
}

// 根据id 获取服务详情
func GetSmServiceById(id int64) (v *SmService, err error) {
	o := orm.NewOrm()
	v = &SmService{}
	if err = o.QueryTable(new(SmService)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
