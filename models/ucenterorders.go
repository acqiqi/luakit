package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"reflect"
	"strings"
	"time"
)

type UcenterOrders struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	OrderNo       string  `json:"order_no"`       // 订单编号
	PlatformKey   string  `json:"platform_key"`   // 平台key
	Cuid          int     `json:"cuid"`           // ucenter id
	CouponKey     string  `json:"coupon_key"`     // 优惠券key
	CouponId      int     `json:"coupon_id"`      // 优惠券id
	CouponPrice   float64 `json:"coupon_price"`   // 优惠券抵扣金额
	CostPrice     float64 `json:"cost_price"`     // 原价
	UnitPrice     float64 `json:"unit_price"`     // 单价
	Price         float64 `json:"price"`          // 现价 支付总价
	GoodsNum      float64 `json:"goods_num"`      // 商品总数量
	PayType       int     `json:"pay_type"`       // 支付类型 0线上 1线下
	PayPlatform   string  `json:"pay_platform"`   // 支付平台 wechat alipay wallet ...
	Status        int     `json:"status"`         // 状态 0下单 2审核通过 7已接单 8已发货 9已结算或者已收货
	IsPay         int     `json:"is_pay"`         // 是否支付 0否 1线下提交 2已支付
	IsUComment    int     `json:"is_u_comment"`   // 是否用户评论 0否 1是
	IsMComment    int     `json:"is_m_comment"`   // 是否商家评论 0否1是
	PayTime       int     `json:"pay_time"`       // 支付时间
	ServicePrice  float64 `json:"service_price"`  // 服务费
	SharePrice    float64 `json:"share_price"`    // 分享出去多少钱
	PlatformPrice float64 `json:"platform_price"` // 平台受益
	ShareLv1      float64 `json:"share_lv1"`      // 一级分享
	ShareLv2      float64 `json:"share_lv2"`      // 二级分享
	OptLv1        float64 `json:"opt_lv1"`        // 一级其他分享
	OptLv2        float64 `json:"opt_lv2"`        // 二级其他分享
	ShareLv1Cuid  int     `json:"share_lv1_cuid"` // 一级分享用户id
	ShareLv2Cuid  int     `json:"share_lv2_cuid"`
	IsAccounts    int     `json:"is_accounts"` // 是否结算
	OptLv1Cuid    int     `json:"opt_lv1_cuid"`
	OptLv2Cuid    int     `json:"opt_lv2_cuid"`
	OrderType     int     `json:"order_type"` // 订单类型0常规购买订单  1团购 10VIP
	PayNo         string  `json:"pay_no"`     // 线上付款订单
	Describe      string  `json:"describe"`   // 描述
	ProjectId     int     `json:"project_id"`
	VipPriceD     float64 `json:"vip_price_d"` // vip折扣多少钱
	Pics          string  `json:"pics"`
	PartnerPrice  float64 `json:"partner_price"`   //合伙人金额
	PartnerId     int     `json:"partner_id"`      //合伙人id
	TotalUsePrice float64 `json:"total_use_price"` //全部占用金额
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterOrders))
}

// 新增用户订单
func AddUcenterOrders(m *UcenterOrders) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据订单编号获取订单
func GetUcenterOrdersOrOrderNo(order_no string) (v *UcenterOrders, err error) {
	o := orm.NewOrm()
	v = &UcenterOrders{}
	if err = o.QueryTable(new(UcenterOrders)).Filter("OrderNo", order_no).Filter("Flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUcenterOrdersOrOrderId(id int) (v *UcenterOrders, err error) {
	o := orm.NewOrm()
	v = &UcenterOrders{}
	if err = o.QueryTable(new(UcenterOrders)).Filter("id", id).Filter("Flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllUcenterOrders(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	query["flag"] = "1"
	o := orm.NewOrm()
	qs := o.QueryTable(new(UcenterOrders))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []UcenterOrders
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// 根据id 修改订单
func UpdateUcenterOrdersById(m *UcenterOrders) (err error) {
	o := orm.NewOrm()
	v := UcenterOrders{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 删除订单
func DeleteUcenterOrders(id int64) (err error) {
	o := orm.NewOrm()
	v := UcenterOrders{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UcenterOrders{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// 设置服务是结算状态
func SetUcenterOrdersAccount(id int64, total_use_price float64) error {
	order, err := GetUcenterOrdersOrOrderId(int(id))
	if err != nil {
		return err
	}
	order.TotalUsePrice = total_use_price
	order.IsAccounts = 1
	if err := UpdateUcenterOrdersById(order); err != nil {
		return err
	}
	return nil
}
