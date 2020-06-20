package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"reflect"
	"strings"
	"time"
)

type MarketingCoupon struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	CouponTplId   int     `json:"coupon_tpl_id"`   // 模板id
	CouponQueueId int     `json:"coupon_queue_id"` // 批量发送id 如果有
	SourceType    int     `json:"source_type"`     // 0后台发送 1新用户注册 2事件触发
	Cuid          int     `json:"cuid"`
	Logo          string  `json:"logo"`
	Title         string  `json:"title"`        // 优惠券名称
	PlatformKey   string  `json:"platform_key"` // 平台key
	Price         float64 `json:"price"`        // 优惠金额/最大优惠金额
	FullPrice     float64 `json:"full_price"`   // 满金额条件
	Type          int     `json:"type"`         // 0满减 1全局折扣 2满折
	Zkb           int     `json:"zkb"`          // 折扣比 0 100
	ProjectId     int     `json:"project_id"`   // 对应项目id （上门端是skuid 或者serviceid）
	ProjectType   int     `json:"project_type"` // 对应项目类型。每个平台不同  （上门端0 sku 1 service）
	EndTime       int     `json:"end_time"`     // 到期时间。被转换的时间戳
	Describe      string  `json:"describe"`     // 描述
	IsUse         int     `json:"is_use"`       // 是否使用
	OrderId       int     `json:"order_id"`     // 使用后订单id
	CouponKey     string  `json:"coupon_key"`   // 优惠券唯一key
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(MarketingCoupon))
}

// 根据订单编号获取订单
func GetMarketingCouponById(id int64) (v *MarketingCoupon, err error) {
	o := orm.NewOrm()
	v = &MarketingCoupon{}
	if err = o.QueryTable(new(MarketingCoupon)).Filter("Id", id).Filter("Flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllMarketingConpon(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	query["flag"] = "1"
	o := orm.NewOrm()
	qs := o.QueryTable(new(MarketingCoupon))
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

	var l []MarketingCoupon
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
