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

type SmUsersComment struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	ServiceId      int    `json:"service_id"`       // 服务id
	ApptId         int    `json:"appt_id"`          // 服务区间
	ServiceSkuId   int    `json:"service_sku_id"`   // 服务sku
	OrderId        int    `json:"order_id"`         // 订单id
	UsersServiceId int    `json:"users_service_id"` // 用户服务表id
	ServiceAreaId  int    `json:"service_area_id"`  // 区域id
	Cuid           int    `json:"cuid"`
	Type           int    `json:"type"` // 0用户 1商家
	Content        string `json:"content"`
	Pics           string `json:"pics"`
	VideoUrl       string `json:"video_url"`
	Avatar         string `json:"avatar"`   // 备份头像
	Nickname       string `json:"nickname"` // 备份昵称
	Star           int    `json:"star"`     // 星 1非常不满意 2不满意 3一般 4满意 5 非常满意
	Tags           string `json:"tags"`
	IsOld          int    `json:"is_old"`
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SmUsersComment))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddSmUsersComment(m *SmUsersComment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPostById retrieves Post by Id. Returns error if
// Id doesn't exist
func GetSmUsersCommentById(id int64) (v *SmUsersComment, err error) {
	o := orm.NewOrm()
	v = &SmUsersComment{}

	if err = o.QueryTable(new(SmUsersComment)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 根据messagekey 获取消息模板
func GetSmUsersCommentByMessageKey(message_key string) (v *SmUsersComment, err error) {
	o := orm.NewOrm()
	v = &SmUsersComment{}

	if err = o.QueryTable(new(SmUsersComment)).Filter("MessageKey", message_key).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllSmUsersComment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SmUsersComment))
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

	var l []SmUsersComment
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
func UpdateSmUsersCommentById(m *SmUsersComment) (err error) {
	o := orm.NewOrm()
	v := SmUsersComment{Id: m.Id}
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
func DeleteSmUsersComment(id int64) (err error) {
	o := orm.NewOrm()
	v := SmUsersComment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SmUsersComment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
