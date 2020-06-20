package models

import (
	"errors"
	"fmt"
	"luakit/utils"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UcenterPlatform struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	PlatformName       string `json:"platform_name"`     // 平台名称
	PlatformUsername   string `json:"platform_username"` // 平台账号
	PlatformPassword   string `json:"platform_password"` // 平台密码
	PlatformAppType    int    `json:"platform_app_type"` // 平台类型 0 api 1 web 5 wechat 6 alipay
	PlatformKey        string `json:"platform_key"`      // 平台key 完全标识
	Status             int    `json:"status"`            // 状态 1运行 0暂停维护 -1禁用
	Ak                 string `json:"ak"`
	Sk                 string `json:"sk"`
	PayName            string `json:"pay_name"`
	PayAk              string `json:"pay_ak"`
	PaySk              string `json:"pay_sk"`
	PayNotifyUrl       string `json:"pay_notify_url"` // 支付回调
	PayNotifyFunc      string `json:"pay_notify_func"`
	MessageCallbackUrl string `json:"message_callback_url"` //消息中心回调地址
	PlatformSecret     string `json:"platform_secret"`
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterPlatform))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddPlatform(m *UcenterPlatform) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPostById retrieves Post by Id. Returns error if
// Id doesn't exist
func GetPlatformById(id int64) (v *UcenterPlatform, err error) {
	o := orm.NewOrm()
	v = &UcenterPlatform{}

	if err = o.QueryTable(new(UcenterPlatform)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetPlatformByKey(platform_key string) (v *UcenterPlatform, err error) {
	o := orm.NewOrm()
	v = &UcenterPlatform{}

	if err = o.QueryTable(new(UcenterPlatform)).Filter("PlatformKey", platform_key).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllPlatform(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UcenterPlatform))
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

	var l []UcenterPlatform
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
func UpdatePlatformById(m *UcenterPlatform) (err error) {
	o := orm.NewOrm()
	v := UcenterPlatform{Id: m.Id}
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
func DeletePlatform(id int64) (err error) {
	o := orm.NewOrm()
	v := UcenterPlatform{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UcenterPlatform{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
