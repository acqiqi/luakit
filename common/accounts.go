package common

import "luakit/models"

type Accounts struct {
	OrderNo     string                   `json:"order_no"`
	PlatformKey string                   `json:"platform_key"`
	Commission  models.UcenterCommission `json:"commission"`
}

func (this *Accounts) InitData() {
	//获取分润信息
	//var err error
	//this.Commission, err = models.GetUcenterCommissionById(5)
}
