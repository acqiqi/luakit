package ucenter

import (
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
	"strconv"
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

// 获取订单列表
func (this *OrdersController) GetOrdersList() {
	bindData := struct {
		utils.PageDataStruct
		Cuid       int64 `json:"cuid"`
		IsPlatform int   `json:"is_platform"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)

	if bindData.Page == 0 {
		bindData.Page = 1
	}
	if bindData.Limit == 0 {
		bindData.Limit = 20
	}

	_, err := models.GetUcenterUsersById(bindData.Cuid)
	if err != nil {
		utils.ApiErr(this.Controller, "当前用户不存在")
	}

	maps := make(map[string]string)
	maps["cuid"] = strconv.FormatInt(bindData.Cuid, 10)
	if bindData.IsPlatform == 1 {
		maps["platform_key"] = this.Platform.PlatformKey
	}
	maps["flag"] = "1"

	list, err := models.GetAllUcenterOrders(maps, []string{}, []string{"id"}, []string{"desc"}, bindData.Limit*(bindData.Page-1), bindData.Limit)
	if err != nil {
		log.Println(err.Error())
		utils.ApiErr(this.Controller, "获取失败")
	}
	utils.ApiOk(this.Controller, "获取成功", utils.PageDataStruct{
		Page:  bindData.Page,
		Limit: bindData.Limit,
		Lists: list,
	})
}

// 获取订单详情
func (this *OrdersController) GetOrdersInfo() {
	bindData := struct {
		OrderNo string `json:"order_no"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)

	order, err := models.GetUcenterOrdersOrOrderNo(bindData.OrderNo)
	log.Println("hehe")
	if err != nil {
		utils.ApiErr(this.Controller, "订单不存在")
	}

	utils.ApiOk(this.Controller, "获取成功", order)
}

// 创建订单
func (this *OrdersController) CreateOrders() {
	bindData := struct {
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
	utils.GetPostJson(this.Controller, &bindData)

	user, err := models.GetUcenterUsersById(bindData.Cuid)
	if err != nil {
		utils.ApiErr(this.Controller, "用户不存在")
	}

	if bindData.CostPrice <= 0 {
		utils.ApiErr(this.Controller, "请设置原价")
	}
	if bindData.UnitPrice <= 0 {
		utils.ApiErr(this.Controller, "请设置单价")
	}
	if bindData.Price <= 0 {
		utils.ApiErr(this.Controller, "请设置价格")
	}

	//查询优惠券
	if bindData.CouponId > 0 {
		_, err := models.GetMarketingCouponById(int64(bindData.CouponId))
		if err != nil {
			utils.ApiErr(this.Controller, "优惠券不存在")
		}
	}

	model := models.UcenterOrders{
		Flag:        1,
		OrderNo:     uuid.NewV4().String(), // 订单号
		PlatformKey: this.Platform.PlatformKey,
		Cuid:        int(user.Id),
		CouponKey:   "",
		CouponId:    bindData.CouponId,
		CouponPrice: 0,
		CostPrice:   bindData.CostPrice,
		UnitPrice:   bindData.UnitPrice,
		Price:       bindData.Price,
		GoodsNum:    bindData.GoodsNum,
		PayType:     0,
		Status:      0, // 已下单
		OrderType:   bindData.OrderType,
		Describe:    bindData.Describe,
		ProjectId:   bindData.ProjectId,
		VipPriceD:   bindData.VipPriceD,
		Pics:        "",
	}

	_, err = models.AddUcenterOrders(&model)
	if err != nil {
		utils.ApiErr(this.Controller, err.Error())
	}

	utils.ApiOk(this.Controller, "创建成功", model)
}

//线上回调支付成功通知
func (this *OrdersController) NotifyOnlinePay() {
	bindData := struct {
		OrderNo     string `json:"order_no"`
		PayNo       string `json:"pay_no"`
		PayPlatform string `json:"pay_platform"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)

	order, err := models.GetUcenterOrdersOrOrderNo(bindData.OrderNo)
	if err != nil {
		utils.ApiErr(this.Controller, "订单不存在")
	}
	if order.PlatformKey != this.Platform.PlatformKey {
		utils.ApiErr(this.Controller, "当前订单不属于本平台，不可操作")
	}
	if order.Status != 0 {
		utils.ApiErr(this.Controller, "当前状态不可修改")
	}

	order.Status = 2
	order.IsPay = 1
	order.PayType = 0
	order.PayNo = bindData.PayNo
	order.PayPlatform = bindData.PayPlatform

	if err := models.UpdateUcenterOrdersById(order); err != nil {
		utils.ApiErr(this.Controller, "操作失败"+err.Error())
	}
	price_str := utils.FormatFloat(order.Price, 2)
	if order.PlatformKey == "DDSM_SHOP" {
		utils.SendAdminSmsLuosimao("叮当匠物客户下单了。金额：" + price_str)
	}

	mp := models.UcenterPriceLog{
		Flag:        1,
		Cuid:        order.Cuid,
		Price:       order.Price,
		PlatformKey: order.PlatformKey,
	}
	models.AddUcenterPriceLog(&mp)

	utils.ApiOk(this.Controller, "操作成功", order)
}

// 删除订单
func (this *OrdersController) CloseOrders() {
	bindData := struct {
		OrderNo string `json:"order_no"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)

	order, err := models.GetUcenterOrdersOrOrderNo(bindData.OrderNo)
	if err != nil {
		utils.ApiErr(this.Controller, "订单不存在")
	}
	if order.PlatformKey != this.Platform.PlatformKey {
		utils.ApiErr(this.Controller, "当前订单不属于本平台，不可操作")
	}
	if order.Status != 0 {
		utils.ApiErr(this.Controller, "当前状态不可修改")
	}
	err = models.DeleteUcenterOrders(order.Id)
	if err != nil {
		utils.ApiErr(this.Controller, "删除失败"+err.Error())
	}
}
