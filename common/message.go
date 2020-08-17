package common

import (
	"log"
	"luakit/models"
	"luakit/utils"
	"strconv"
	"time"
)

type MessageCenterUtils struct {
	Id             int                   `json:"id"`
	AdminMsg       string                `json:"admin_msg"`
	Price          float64               `json:"price"`            // 金额
	Desc           string                `json:"desc"`             //描述
	Users          models.UcenterUsers   `json:"ucenter_users"`    //用户信息
	ToUsers        models.UcenterUsers   `json:"to_ucenter_users"` //发送给用户信息
	SmUsers        models.SmUsers        `json:"sm_users"`         // 上门用户
	ToSmUsers      models.SmUsers        `json:"to_sm_users"`      //发送给上门用户
	SmService      models.SmService      `json:"sm_service"`       //上门服务
	SmUsersService models.SmUsersService `json:"sm_users_service"` //上门用户服务
	Orders         models.UcenterOrders  `json:"orders"`           //订单

	PlatformKey string `json:"platform_key"` //平台key
	MessageKey  string `json:"message_key"`  //消息key

}

// 发送消息
func (this *MessageCenterUtils) PushMessage() error {
	maps := make(map[string]string)
	// to_uc
	maps["to_uc_nickname"] = this.ToUsers.Nickname
	maps["to_uc_avatar"] = this.ToUsers.Avatar
	maps["to_uc_mobile"] = this.ToUsers.Mobile
	maps["to_uc_id"] = strconv.Itoa(int(this.ToUsers.Id))
	// orders
	maps["orders_id"] = strconv.Itoa(int(this.Orders.Id))
	maps["orders_no"] = this.Orders.OrderNo
	maps["orders_pay_no"] = this.Orders.PayNo
	// sm_users
	maps["sm_user_gj_name"] = this.SmUsers.GjName
	maps["sm_user_gj_idcard_id"] = this.SmUsers.GjIdcardId
	maps["sm_user_gj_idcard_address"] = this.SmUsers.GjIdcardAddress
	maps["sm_user_gj_err_msg"] = this.SmUsers.GjErrMsg
	// admin mobile
	maps["admin_mobile_1"] = "18666623878"
	maps["admin_mobile_2"] = "18688832458"
	maps["admin_mobile_3"] = "13682400039"
	maps["admin_msg"] = ""

	// opt
	maps["id"] = strconv.Itoa(this.Id)
	now := time.Now()                                    // 当前 datetime 时间
	maps["set_time"] = now.Format("2006-01-02 15:04:05") // 把当前 datetime 时间转换成时间字符串
	price_str := strconv.FormatFloat(utils.Decimal(this.Price), 'f', 2, 64)
	maps["price"] = price_str

	// users
	maps["users_id"] = strconv.Itoa(int(this.Users.Id))
	maps["users_nickname"] = this.Users.Nickname
	maps["users_mobile"] = this.Users.Avatar
	maps["users_avatar"] = this.Users.Mobile
	log.Println("打印消息", maps)

	push_data := struct {
		PushData    map[string]string `json:"push_data"`
		PlatformKey string            `json:"platform_key"`
		Cuid        int64             `json:"cuid"`
		MessageKey  string            `json:"message_key"`
	}{
		PushData:    maps,
		PlatformKey: this.PlatformKey,
		Cuid:        this.Users.Id,
		MessageKey:  this.MessageKey,
	}
	base_url := utils.ServiceConfig.String("base_url") + "v1/message/api/publishmessagequeue"

	err := utils.HttpPostJsonNotCallback(base_url, push_data, this.PlatformKey)
	if err != nil {
		return err
	}
	return nil
}
