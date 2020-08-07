package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

// 红包表
type MarketingPacket struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Cuid        int     `json:"cuid"`
	Price       float64 `json:"price"`
	Title       string  `json:"title"`
	Describe    string  `json:"describe"` // 描述
	Status      int     `json:"status"`   // 0正常 1使用 -1过期或者禁用
	Type        string  `json:"type"`     // 来源渠道
	PlatformKey string  `json:"platform_key"`
	PacketNo    string  `json:"packet_no"`
	SCuid       int     `json:"s_cuid"` // 来源用户
}

const (
	PACKET_TYPE_WORK        = "PACKET_TYPE_WORK"        //工人收入
	PACKET_TYPE_WORK_ST1    = "PACKET_TYPE_WORK_ST1"    //工人师徒1
	PACKET_TYPE_WORK_ST2    = "PACKET_TYPE_WORK_ST2"    //工人师徒2
	PACKET_TYPE_SHARE1      = "PACKET_TYPE_SHARE1"      //分享收益1
	PACKET_TYPE_SHARE2      = "PACKET_TYPE_SHARE2"      //分享收益2
	PACKET_TYPE_SHARE_REG   = "PACKET_TYPE_SHARE_REG"   //注册分享
	PACKET_TYPE_PW          = "PACKET_TYPE_PW"          //大转盘
	PACKET_TYPE_KILL_PACKET = "PACKET_TYPE_KILL_PACKET" //现金红包抽奖
)

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(MarketingPacket))
}

// 生成用户红包
func GenerUserPacket(m *MarketingPacket) (id int64, err error) {
	m.Status = 0
	m.Flag = 1
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
