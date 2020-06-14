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

type UcenterAccounts struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid        int     `orm:"cuid"`         // ucenter uid
	PlatformKey string  `orm:"platform_key"` // 牵扯的平台
	Type        int     `orm:"type"`         // 0直接收益 1分润收益 2师徒收益 10现金红包收益 50充值收益 100提现 101购买商品
	Level       int     `orm:"level"`        // 收益等级 比如 0一级分享收益 1二级分享收益
	Content     string  `orm:"content"`      // 详细内容
	Describe    string  `orm:"describe"`     // 描述  主要是显示这里
	ProjectId   int     `orm:"project_id"`   // 项目id
	OrderId     int     `orm:"order_id"`     // 订单id
	OrderNo     int     `orm:"order_no"`     // 订单编号
	Price       float64 `orm:"price"`        // 金额
	IsDz        int     `orm:"is_dz"`        // 是否到账 1是
	SourceCuid  int     `orm:"source_cuid"`  // 来源用户。比如是谁分享产生的给你费用
	ProjectName string  `orm:"project_name"` // 项目名称
	Title       string  `orm:"title"`        // 标题
	AccountNo   string  `orm:"account_no"`   // 订单号
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterAccounts))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddUcenterAccounts(m *UcenterAccounts) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPostById retrieves Post by Id. Returns error if
// Id doesn't exist
func GetAccountsById(id int64) (v *UcenterAccounts, err error) {
	o := orm.NewOrm()
	v = &UcenterAccounts{}

	if err = o.QueryTable(new(UcenterAccounts)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllUcenterAccounts(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UcenterAccounts))
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

	var l []UcenterAccounts
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

// UpdatePost updates Post by Id and returns error if
// the record to be updated doesn't exist
func UpdateUcenterAccountsById(m *UcenterAccounts) (err error) {
	o := orm.NewOrm()
	v := Message{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePost deletes Post by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUcenterAccounts(id int64) (err error) {
	o := orm.NewOrm()
	v := Message{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Message{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
