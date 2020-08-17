package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type PartnerUsers struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Username string  `json:"username"`  // 账号
	Password string  `json:"password"`  // 密码
	Mobile   string  `json:"mobile"`    // 手机号
	Nickname string  `json:"nickname"`  // 昵称
	Email    string  `json:"email"`     // 邮箱
	Gender   string  `json:"gender"`    // 性别
	Status   int     `json:"status"`    // 状态 0停用 1启用
	RoleType int     `json:"role_type"` // 0无权限 1小区合伙人 2区县级合伙人 3城市合伙人 4省合伙人
	Describe string  `json:"describe"`
	Money    float64 `json:"money"` // 余额
	Avatar   string  `json:"avatar"`
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(PartnerUsers))
}

// 用户服务订单
func GetPartnerUsersById(id int64) (v *PartnerUsers, err error) {
	o := orm.NewOrm()
	v = &PartnerUsers{}
	if err = o.QueryTable(new(PartnerUsers)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func UpdatePartnerUsersById(m *PartnerUsers) (err error) {
	o := orm.NewOrm()
	v := PartnerUsers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 插入金额
func SetPartnerUsersOkMoney(cuid int64, price float64) error {
	user, err := GetPartnerUsersById(cuid)
	if err != nil {
		return err
	}
	user.Money = user.Money + price
	if err := UpdatePartnerUsersById(user); err != nil {
		return err
	}
	return nil
}
