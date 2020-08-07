package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type UcenterAccounts struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid        int     `json:"cuid"`         // ucenter uid
	PlatformKey string  `json:"platform_key"` // 牵扯的平台
	Type        int     `json:"type"`         // 0直接收益 1分润收益 2师徒收益 10现金红包收益 50充值收益 100提现 101购买商品
	Level       int     `json:"level"`        // 收益等级 比如 0一级分享收益 1二级分享收益
	Content     string  `json:"content"`      // 详细内容
	Describe    string  `json:"describe"`     // 描述  主要是显示这里
	ProjectId   int     `json:"project_id"`   // 项目id
	OrderId     int     `json:"order_id"`     // 订单id
	OrderNo     string  `json:"order_no"`     // 订单编号
	Price       float64 `json:"price"`        // 金额
	IsDz        int     `json:"is_dz"`        // 是否到账 1是
	SourceCuid  int     `json:"source_cuid"`  // 来源用户。比如是谁分享产生的给你费用
	ProjectName string  `json:"project_name"` // 项目名称
	Title       string  `json:"title"`        // 标题
	AccountNo   string  `json:"account_no"`   // 订单号
	IsOld       int     `json:"is_old"`       //是否老订单
}

const (
	ACCOUNTS_ZJSY   = 0  //直接收益
	ACCOUNTS_SHARE  = 1  //分享收益
	ACCOUNTS_ST     = 2  //师徒收益
	ACCOUNTS_PACKET = 10 //现金红包收益
	ACCOUNTS_CZ     = 50 //充值收益
)

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterAccounts))
}

// 新增收支信息
func AddUcenterAccounts(m *UcenterAccounts) (id int64, err error) {
	m.Flag = 1
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
