package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

// 用户服务工人表
type SmUsersServiceWork struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	//User *UcenterUsers `orm:"rel(fk);column(cuid)" json:"user"`
	Cuid int `json:"cuid"`

	UsersServiceId int     `json:"users_service_id"` // 用户服务表id
	Price          float64 `json:"price"`            // 分得金额。默认等分。
	IsEnd          int     `json:"is_end"`           // 是否结束，一般跟订单结算
	ServiceId      int     `json:"service_id"`       // 服务id
	OrderId        int     `json:"order_id"`         // 订单id
	//Order        *UcenterOrders     `orm:"rel(fk)" json:"order_id"`         // 订单id

	FrPrice float64 `json:"fr_price"` // 最终从这个用户分润出去的金额
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SmUsersServiceWork))
}

// 根据用户服务id 获取主线工人列表
func GetSmUsersServiceWorkListByUsId(us_id int64) (row int64, v *[]SmUsersServiceWork, err error) {
	o := orm.NewOrm()
	v = &[]SmUsersServiceWork{}
	if row, err = o.QueryTable(new(SmUsersServiceWork)).Filter("UsersServiceId", us_id).RelatedSel().All(v); err == nil {
		return row, v, nil
	}
	return row, nil, err
}

// UpdatePost updates Post by Id and returns error if
// the record to be updated doesn't exist
func UpdateSmUsersServiceWorkById(m *SmUsersServiceWork) (err error) {
	o := orm.NewOrm()
	v := SmUsersServiceWork{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
