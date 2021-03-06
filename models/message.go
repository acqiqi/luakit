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

type Message struct {
	Id              int64     `json:"id"`
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt       time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag            int       `orm:"default(1)" json:"flag"` //-1删除
	Cuid            int64     `json:"cuid"`
	MessageKey      string    `json:"message_key"`
	Title           string    `json:"title"`        // 标题
	Desc            string    `json:"desc"`         // 描述
	Content         string    `json:"content"`      // 内容
	MessageType     int       `json:"message_type"` // 消息类型
	PathType        string    `json:"path_type"`    // 链接类型
	PathId          string    `json:"path_id"`      // 链接id
	IsFormId        int       `json:"is_form_id"`   // 是否消息模板
	SmallTplId      string    `json:"small_tpl_id"` // 消息模板id
	SmallTplContent string    `json:"small_tpl_content"`
	SmallTplOpenid  string    `json:"small_tpl_openid"`
	IsSms           int       `json:"is_sms"` // 是否发送短信
	Mobile          string    `json:"mobile"`
	SmsContent      string    `json:"sms_content"` // 短信内容
	IsEmail         int       `json:"is_email"`    // 是否发邮件
	Email           string    `json:"email"`
	EmailTitle      string    `json:"email_title"`
	EmailContent    string    `json:"email_content"`
	MsgTplId        int64     `json:"msg_tpl_id"`   // message tpl id
	PlatformKey     string    `json:"platform_key"` // 平台key
	PushData        string    `json:"push_data"`
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(Message))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddMessage(m *Message) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPostById retrieves Post by Id. Returns error if
// Id doesn't exist
func GetMessageById(id int64) (v *Message, err error) {
	o := orm.NewOrm()
	v = &Message{}

	if err = o.QueryTable(new(Message)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 根据messagekey 获取消息模板
func GetMessageByMessageKey(message_key string) (v *Message, err error) {
	o := orm.NewOrm()
	v = &Message{}

	if err = o.QueryTable(new(Message)).Filter("MessageKey", message_key).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllMessage(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Message))
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

	var l []Message
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
func UpdateMessageById(m *Message) (err error) {
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
func DeleteMessage(id int64) (err error) {
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
