package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"log"
	"luakit/utils"
	"strconv"
	"time"
)

type UcenterOrdersLog struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid        int     `json:"cuid"`
	PlatformKey string  `json:"platform_key"`
	OrderId     int     `json:"order_id"`
	OrderNo     string  `json:"order_no"`
	Type        int     `json:"type"`         // 0平台收益 1商家收益(工人) 2分润收入 3师徒收入 10合伙人收入
	ProjectType int     `json:"project_type"` // 项目类型 0上门服务
	ProjectId   int     `json:"project_id"`   // 对应项目id us_id
	Describe    string  `json:"describe"`     // 描述
	Price       float64 `json:"price"`
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterOrdersLog))
}

const (
	UCENTER_ORDERS_LOG_TYPE_PLATFORM = 0  // 平台
	UCENTER_ORDERS_LOG_TYPE_WORK     = 1  // 工人
	UCENTER_ORDERS_LOG_TYPE_SHARE    = 2  // 分润
	UCENTER_ORDERS_LOG_TYPE_ST       = 3  // 师徒
	UCENTER_ORDERS_LOG_TYPE_PARTNER  = 10 // 合伙人
)

// 新增订单日志
func AddUcenterOrdersLog(m *UcenterOrdersLog) (id int64, err error) {
	m.Flag = 1
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 根据订单编号获取总共金额
func SumUcenterOrdersLogOrOrderId(order_id int) (price float64, err error) {

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("select SUM(price) from vhake_ucenter_orders_log where order_id = ? AND flag =1", order_id).Values(&maps)
	if err != nil {
		log.Println("errle", err)
		return 0, err
	}
	if num > 0 {
		if maps[0]["SUM(price)"] != nil {
			price, err = strconv.ParseFloat(maps[0]["SUM(price)"].(string), 64)
			if err != nil {
				return 0, err
			}
			return price, err
		} else {
			return 0, errors.New("查询失败1")
		}

	} else {
		return 0, errors.New("查询失败")
	}

}
