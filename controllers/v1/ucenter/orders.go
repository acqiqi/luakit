package ucenter

import (
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
)

type OrdersController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *OrdersController) Prepare() {
	log.Println("Prepare")
	platform := apibase.Auth2PlatformBase(this.Controller)
	this.Platform = *platform
}

// 获取订单详情
func (this *OrdersController) GetOrdersInfo() {
	c := struct {
		OrderNo string `json:"order_no"`
	}{}
	utils.GetPostJson(this.Controller, &c)

	order, err := models.GetUcenterOrdersOrOrderNo(c.OrderNo)
	log.Println("hehe")
	if err != nil {
		utils.ApiErr(this.Controller, "订单不存在")
	}

	utils.ApiOk(this.Controller, "获取成功", order)
}

// 创建订单
func (this *OrdersController) CreateOrders() {
	c := struct {
		Cuid        int64   `json:"cuid"`         // ucenter id
		CouponId    int     `json:"coupon_id"`    // 优惠券id
		CouponPrice float64 `json:"coupon_price"` // 优惠券抵扣金额
		CostPrice   float64 `json:"cost_price"`   // 原价
		UnitPrice   float64 `json:"unit_price"`   // 单价
		Price       float64 `json:"price"`        // 现价 支付总价
		VipPriceD   float64 `json:"vip_price_d"`  // vip折扣多少钱
		GoodsNum    float64 `json:"goods_num"`    // 商品总数量
		Describe    string  `json:"describe"`     // 描述
		OrderType   int     `json:"order_type"`   // 所属类型 一般0默认商品
		ProjectId   int     `json:"project_id"`   //对应项目id
	}{}
	utils.GetPostJson(this.Controller, &c)

	user, err := models.GetUcenterUsersById(c.Cuid)
	if err != nil {
		utils.ApiErr(this.Controller, "用户不存在")
	}

	if c.CostPrice <= 0 {
		utils.ApiErr(this.Controller, "请设置原价")
	}
	if c.UnitPrice <= 0 {
		utils.ApiErr(this.Controller, "请设置单价")
	}
	if c.Price <= 0 {
		utils.ApiErr(this.Controller, "请设置价格")
	}
	model := models.UcenterOrders{
		Id:          0,
		Flag:        1,
		OrderNo:     uuid.NewV4().String(), // 订单号
		PlatformKey: this.Platform.PlatformKey,
		Cuid:        int(user.Id),
		CouponKey:   "",
		CouponId:    0,
		CouponPrice: 0,
		CostPrice:   c.CostPrice,
		UnitPrice:   c.UnitPrice,
		Price:       c.Price,
		GoodsNum:    c.GoodsNum,
		PayType:     0,
		Status:      0, // 已下单
		OrderType:   c.OrderType,
		Describe:    c.Describe,
		ProjectId:   c.ProjectId,
		VipPriceD:   c.VipPriceD,
		Pics:        "",
	}

	_, err = models.AddUcenterOrders(&model)
	if err != nil {
		utils.ApiErr(this.Controller, err.Error())
	}

	utils.ApiOk(this.Controller, "创建成功", model)
}
