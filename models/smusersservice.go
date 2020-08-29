package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

// 用户服务表
type SmUsersService struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	ServiceId        int     `json:"service_id"`      // 服务id
	ServiceSkuId     int     `json:"service_sku_id"`  // sku id
	ServiceAreaId    int     `json:"service_area_id"` // 区域id
	ApptId           int     `json:"appt_id"`
	CartId           int     `json:"cart_id"` // 购物车对应id
	ServiceLogo      string  `json:"service_logo"`
	ServiceTitle     string  `json:"service_title"` // 标题
	ApptTitle        string  `json:"appt_title"`
	CostPrice        float64 `json:"cost_price"`       // 原价单价
	UnitPrice        float64 `json:"unit_price"`       // 单价
	TotalCostPrice   float64 `json:"total_cost_price"` // 全部原价
	TotalPrice       float64 `json:"total_price"`      // 全部价格
	OrderId          int     `json:"order_id"`
	Cuid             int     `json:"cuid"`
	Status           int     `json:"status"`      // 0 未付款 1付款(拼团属于还没有拼成)  5已确认（拼团拼成）  8已接单 10已服务 待处理
	WorkCuid         int     `json:"work_cuid"`   // 分配工人uid
	NotAppt          int     `json:"not_appt"`    // 是否暂时不选择服务区间 0否 1是
	Price            float64 `json:"price"`       // 实际付款价格
	Num              int     `json:"num"`         // 数量
	GjPrice          float64 `json:"gj_price"`    // 工匠金额
	IsGjPrice        int     `json:"is_gj_price"` // 是否固定工匠金额
	SkuName          string  `json:"sku_name"`    // sku 名称
	AddressId        int     `json:"address_id"`
	ServiceNo        string  `json:"service_no"`
	IsComment        int     `json:"is_comment"`      // 是否评论
	IsWorkComment    int     `json:"is_work_comment"` // 是否商家评论
	Qianzi           string  `json:"qianzi"`
	IsGp             int     `json:"is_gp"`               // 是否团购
	UsersServiceGpId int     `json:"users_service_gp_id"` // 用户团购id
	ServiceGpId      int     `json:"service_gp_id"`       // 团购id
	IsSeckill        int     `json:"is_seckill"`          // 是否秒杀
	ServiceSeckillId int     `json:"service_seckill_id"`
	Describe         string  `json:"describe"`
	Pics             string  `json:"pics"`
	IsNewWork        int     `json:"is_new_work"` // 多工人订单-后结算
	IsJiesuan        int     `json:"is_jiesuan"`
	TotalFrPrice     float64 `json:"total_fr_price"`   // 总分润出去的金额
	TotalWorkPrice   float64 `json:"total_work_price"` // 总工人金额
	AdminDesc        string  `json:"admin_desc"`

	UseShareLv1 int `orm:"use_share_lv1"` // 强制一级分享
	UseShareLv2 int `orm:"use_share_lv2"` // 强制二级分享
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SmUsersService))
}

// 用户服务订单
func GetSmUsersServiceById(id int64) (v *SmUsersService, err error) {
	o := orm.NewOrm()
	v = &SmUsersService{}
	if err = o.QueryTable(new(SmUsersService)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 查询没有结束的订单
func CheckSmUsersServiceNotEnd(order_id int64) (row int64, v *[]SmUsersService, err error) {
	o := orm.NewOrm()
	v = &[]SmUsersService{}
	if row, err = o.QueryTable(new(SmUsersService)).Filter("OrderId", order_id).Filter("IsJiesuan", 0).Filter("flag", 1).RelatedSel().All(v); err == nil {
		return row, v, nil
	}
	return 0, nil, err
}

// 更新用户服务订单
func UpdateSmUsersServiceById(m *SmUsersService) (err error) {
	o := orm.NewOrm()
	v := SmUsersService{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 设置服务是结算状态
func SetSmUsersServiceJiesuan(id int64) error {
	us, err := GetSmUsersServiceById(id)
	if err != nil {
		return err
	}
	us.IsJiesuan = 1
	if err := UpdateSmUsersServiceById(us); err != nil {
		return err
	}
	return nil
}
