package common

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"log"
	"luakit/models"
	"luakit/utils"
	"strconv"
)

type Accounts struct {
	OrderNo             string                      `json:"order_no"`            //订单编号
	OrderId             int                         `json:"order_id"`            //订单编号
	SmUsersServiceId    int64                       `json:"sm_users_service_id"` // 上门端用户服务订单id
	SmUsersService      models.SmUsersService       `json:"sm_users_service"`    //上门服务订单
	SmUsersServiceWorks []models.SmUsersServiceWork `json:"sm_users_service_works"`
	SmService           models.SmService            `json:"sm_service"`   // 上门服务
	PlatformKey         string                      `json:"platform_key"` //平台key
	Commission          models.UcenterCommission    `json:"commission"`
	Orders              models.UcenterOrders        `json:"orders"`
	Users               models.UcenterUsers         `json:"users"`
}

// 初始化订单和用户
func (this *Accounts) InitOrdersNoAndUser() (err error) {
	// 查询订单是否存在
	order, err := models.GetUcenterOrdersOrOrderNo(this.OrderNo)
	if err != nil {
		return errors.New("订单不存在")
	}
	this.Orders = *order
	this.OrderNo = order.OrderNo
	this.OrderId = int(order.Id)
	// 查询对应用户是否存在
	users, err := models.GetUcenterUsersById(int64(order.Cuid))
	if err != nil {
		return errors.New("用户不存在")
	}
	this.Users = *users
	return nil
}

// 初始化订单id和用户
func (this *Accounts) InitOrdersAndUser() (err error) {
	// 查询订单是否存在
	order, err := models.GetUcenterOrdersOrOrderId(this.OrderId)
	if err != nil {
		return errors.New("订单不存在")
	}
	this.Orders = *order
	this.OrderNo = order.OrderNo
	this.OrderId = int(order.Id)

	if this.Orders.Status < 2 {
		return errors.New("订单状态不满足")
	}
	if this.Orders.IsPay != 1 {
		return errors.New("订单未支付")
	}
	if this.Orders.IsAccounts != 0 {
		return errors.New("订单已结算,请勿重复结算")
	}

	// 查询对应用户是否存在
	users, err := models.GetUcenterUsersById(int64(order.Cuid))
	if err != nil {
		return errors.New("用户不存在")
	}
	this.Users = *users
	return nil
}

// 根据SmUsersService 的id 初始化
func (this *Accounts) InitSmUsersService() (err error) {
	us, err := models.GetSmUsersServiceById(this.SmUsersServiceId)
	if err != nil {
		return errors.New("服务订单不存在")
	}

	this.SmUsersService = *us
	//查询状态
	if this.SmUsersService.Status < 5 {
		return errors.New("当前订单状态不可结算")
	}
	this.OrderId = this.SmUsersService.OrderId

	//初始化服务
	service, err := models.GetSmServiceById(int64(us.ServiceId))
	if err != nil {
		return err
	}
	this.SmService = *service

	//初始化订单
	if err := this.InitOrdersAndUser(); err != nil {
		return err
	}

	if this.SmUsersService.WorkCuid != 0 {
		//初始化工人列表
		err := this.initWorks()
		if err != nil {
			return errors.New("工人初始化失败！")
		}
	}
	return nil
}

// 初始化工人列表
func (this *Accounts) initWorks() (err error) {
	row, list, err := models.GetSmUsersServiceWorkListByUsId(this.SmUsersServiceId)
	if err != nil {
		return errors.New("工匠查询失败")
	}
	if row > 0 {
		this.SmUsersServiceWorks = *list
		return nil
	} else {
		return errors.New("查询不到工人")
	}
}

// 结算服务订单
func (this *Accounts) AccountsSmServiceOrders() error {
	log.Println("查看订单")
	log.Println(this.Orders)
	log.Println("查看服务订单")
	log.Println(this.SmUsersService)
	log.Println("查看工人列表")
	log.Println(this.SmUsersServiceWorks)
	log.Println(len(this.SmUsersServiceWorks))
	log.Println("查看服务")
	log.Println(this.SmService)

	//查询订单状态是否是已结束
	if this.SmUsersService.Status != 10 {
		return errors.New("当前订单未结束")
	}
	//判断是否结算
	if this.SmUsersService.IsJiesuan != 0 {
		return errors.New("已经结算，请勿重复结算")
	}

	// 处理工人
	log.Println("工人信息", len(this.SmUsersServiceWorks))
	if len(this.SmUsersServiceWorks) == 0 {
		return errors.New("安排工人数据有误")
	}
	work_total_price := 0.00
	// 循环处理工人数据
	for _, v := range this.SmUsersServiceWorks {
		// 结算单个工人金额
		work_item_total_price, err := this.accountsWork(v)
		if err != nil {
			return err
		}
		work_total_price = work_total_price + work_item_total_price
		log.Println(work_item_total_price)
	}

	this.Orders.TotalUsePrice = this.Orders.TotalUsePrice + work_total_price
	// 更新订单的占用金额
	if err := models.UpdateUcenterOrdersById(&this.Orders); err != nil {
		return errors.New("订单修改占用金额失败")
	}
	log.Println("全部工人金额", work_total_price)

	//操作服务订单已结算
	models.SetSmUsersServiceJiesuan(this.SmUsersService.Id)
	// 查询当前服务订单有多少个未结算, 用于全部结算之后结算总订单
	row, _, err := models.CheckSmUsersServiceNotEnd(int64(this.OrderId))
	if err != nil {
		return err
	}
	log.Println("未服务订单数量:", row)
	// 检测未结算订单 小于等于1 结算总订单
	if row < 1 {
		//直接可处理分润结束总订单
		this.AccountsOrders()
	}
	return nil

	//this.SmUsersService
}

// 结算主订单工人
func (this *Accounts) accountsWork(work models.SmUsersServiceWork) (total_price float64, err error) {

	total_price = 0

	if work.Price > 0 {
		log.Println("工人钱包数据", work)
		// 查询师徒分润比
		st_all, st_lv1, st_lv2, work_scale, err := this.getStScale()
		if err != nil {
			return 0, err
		}
		//抽出分润金额
		st_all_price := utils.Decimal(float64(work.Price) * (float64(st_all) / 100))
		st_lv1_price := utils.Decimal(st_all_price * (float64(st_lv1) / 100))
		st_lv2_price := utils.Decimal(st_all_price * (float64(st_lv2) / 100))
		//查询工人抽成
		work_scale_price := utils.Decimal(float64(work.Price) * (float64(work_scale) / 100))
		log.Println("师徒收益总", st_all_price, st_lv1_price, st_lv2_price, work_scale_price)

		work_price := work.Price - st_all_price - work_scale_price //计算出工人实际获取金额

		//插入计算平台抽成 直接插入记录
		m := new(models.UcenterOrdersLog)
		m.PlatformKey = "DDSM_CLEINT"
		m.Price = work_scale_price
		m.Cuid = 0
		m.OrderId = int(this.Orders.Id)
		m.OrderNo = this.Orders.OrderNo
		m.ProjectId = int(this.SmUsersService.Id)
		m.Type = models.UCENTER_ORDERS_LOG_TYPE_PLATFORM
		m.Describe = "平台工人抽成"
		models.AddUcenterOrdersLog(m)
		total_price = total_price + work_scale_price //插入给工人的钱

		// 查询工人用户
		work_user, err := models.GetUcenterUsersById(int64(work.Cuid))
		if err != nil {
			return 0, errors.New("工人用户查询失败")
		}
		//插入工人收益
		models.SetUcenterUsersOkMoney(work_user.Id, work_price)
		am := models.UcenterAccounts{
			Cuid:        int(work_user.Id),
			PlatformKey: "DDSM_MANAGER",
			Type:        models.ACCOUNTS_ZJSY,
			Level:       0,
			Content:     "订单直接收益" + strconv.FormatFloat(utils.Decimal(work_price), 'f', 2, 64),
			Describe:    "订单直接收益" + strconv.FormatFloat(utils.Decimal(work_price), 'f', 2, 64),
			ProjectId:   0,
			OrderId:     int(this.Orders.Id),
			OrderNo:     this.Orders.OrderNo,
			Price:       work_price,
			IsDz:        1,
			SourceCuid:  0,
			ProjectName: this.SmService.Title,
			Title:       "红包到账",
			AccountNo:   uuid.NewV4().String(),
		}
		models.AddUcenterAccounts(&am)

		m = new(models.UcenterOrdersLog)
		m.PlatformKey = "DDSM_CLEINT"
		m.Price = work_price
		m.Cuid = work.Cuid
		m.OrderId = int(this.Orders.Id)
		m.OrderNo = this.Orders.OrderNo
		m.ProjectId = int(this.SmUsersService.Id)
		m.Type = models.UCENTER_ORDERS_LOG_TYPE_SHARE
		m.Describe = "直接收益"
		models.AddUcenterOrdersLog(m)
		total_price = total_price + work_price //插入给工人的钱
		//发布消息
		msg_utils := new(MessageCenterUtils)
		msg_utils.Price = m.Price
		msg_utils.Users = *work_user
		msg_utils.SmService = this.SmService
		msg_utils.SmUsersService = this.SmUsersService
		msg_utils.Orders = this.Orders

		msg_utils.PlatformKey = "DDSM_CLIENT"
		msg_utils.MessageKey = "SY_DZ"
		msg_utils.PushMessage()

		//插入一级分润收益
		//查询是否有师傅
		check_lv1, err := models.CheckStUser(int64(work.Cuid))
		if err == nil {
			log.Println("查到师傅", check_lv1)

			check_lv1_user, err := models.GetUcenterUsersById(int64(check_lv1.Cuid))
			if err != nil {
				return 0, errors.New("一级师父查询有误")
			}

			//处理第一级师徒关系
			//生成红包
			upack := new(models.MarketingPacket)
			upack.PlatformKey = "DDSM_MANAGER"
			upack.Cuid = check_lv1.Cuid
			upack.SCuid = work.Cuid
			upack.Price = st_lv1_price
			upack.Title = "一级师徒收益"
			upack.Describe = "师徒收益到账" + strconv.FormatFloat(utils.Decimal(st_lv1_price), 'f', 2, 64)
			upack.Type = models.PACKET_TYPE_WORK_ST1
			upack.PacketNo = uuid.NewV4().String()
			models.GenerUserPacket(upack)
			// 订单日志
			m := new(models.UcenterOrdersLog)
			m.PlatformKey = "DDSM_CLEINT"
			m.Price = st_lv1_price
			m.Cuid = check_lv1.Cuid
			m.OrderId = int(this.Orders.Id)
			m.OrderNo = this.Orders.OrderNo
			m.ProjectId = int(this.SmUsersService.Id)
			m.Type = models.UCENTER_ORDERS_LOG_TYPE_ST
			m.Describe = "一级师徒收益"
			models.AddUcenterOrdersLog(m)
			total_price = total_price + st_lv1_price //插入给一级师父的钱
			//发布消息
			msg_utils := new(MessageCenterUtils)
			msg_utils.Price = m.Price
			msg_utils.Users = *check_lv1_user
			msg_utils.SmService = this.SmService
			msg_utils.SmUsersService = this.SmUsersService
			msg_utils.Orders = this.Orders

			msg_utils.PlatformKey = "DDSM_CLIENT"
			msg_utils.MessageKey = "ST_SY_DZ"
			msg_utils.PushMessage()

			//查询是否有二级师傅
			check_lv2, err := models.CheckStUser(int64(check_lv1.Cuid))
			if err == nil {
				log.Println("查到二级师傅", check_lv2)

				check_lv2_user, err := models.GetUcenterUsersById(int64(check_lv2.Cuid))
				if err != nil {
					return 0, errors.New("二级师父查询有误")
				}

				//处理第二级师徒关系
				//生成红包
				upack := new(models.MarketingPacket)
				upack.PlatformKey = "DDSM_MANAGER"
				upack.Cuid = check_lv2.Cuid
				upack.SCuid = check_lv1.Cuid
				upack.Price = st_lv2_price
				upack.Title = "二级师徒收益"
				upack.Describe = "师徒收益到账" + strconv.FormatFloat(utils.Decimal(st_lv1_price), 'f', 2, 64)
				upack.Type = models.PACKET_TYPE_WORK_ST2
				upack.PacketNo = uuid.NewV4().String()
				models.GenerUserPacket(upack)
				// 订单日志
				m := new(models.UcenterOrdersLog)
				m.PlatformKey = "DDSM_CLEINT"
				m.Price = st_lv2_price
				m.Cuid = check_lv2.Cuid
				m.OrderId = int(this.Orders.Id)
				m.OrderNo = this.Orders.OrderNo
				m.ProjectId = int(this.SmUsersService.Id)
				m.Type = models.UCENTER_ORDERS_LOG_TYPE_ST
				m.Describe = "二级师徒收益"
				models.AddUcenterOrdersLog(m)
				total_price = total_price + st_lv2_price //插入给二级级师父的钱

				//发布消息
				msg_utils := new(MessageCenterUtils)
				msg_utils.Price = m.Price
				msg_utils.Users = *check_lv2_user
				msg_utils.SmService = this.SmService
				msg_utils.SmUsersService = this.SmUsersService
				msg_utils.Orders = this.Orders

				msg_utils.PlatformKey = "DDSM_CLIENT"
				msg_utils.MessageKey = "ST_SY_DZ"
				msg_utils.PushMessage()
			}
		}

		//work.Price = 20
		//models.UpdateSmUsersServiceWorkById(&work)

		//msg := new(MessageCenterUtils)
		//msg.Orders = this.Orders
		//msg.Users = this.Users
		//msg.Price = 12.333
		//msg.Id = 3
		//msg.PushMessage()
		log.Println("师徒分润比例", st_all, st_lv1, st_lv2)
	} else {
		return 0, errors.New("工人订单金额不能小于等于0")
	}
	return total_price, nil
}

// 获取师徒分润比例
func (this *Accounts) getStScale() (st_scale_all int, st_scale_lv1 int, st_scale_lv2 int, work_scale int, err error) {
	// 查询师徒总分润比
	st_scale_all_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_ST_SCALE_ALL")
	st_scale_all, err = strconv.Atoi(st_scale_all_str)
	if err != nil || st_scale_all <= 0 {
		return 0, 0, 0, 0, errors.New("师徒分润信息有误")
	}
	// 查询师徒一级
	st_scale_lv1_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_ST_SCALE_LV1")
	st_scale_lv1, err = strconv.Atoi(st_scale_lv1_str)
	if err != nil || st_scale_lv1 <= 0 {
		return 0, 0, 0, 0, errors.New("师徒分润一级信息有误")
	}
	// 查询师徒二级
	st_scale_lv2_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_ST_SCALE_LV2")
	st_scale_lv2, err = strconv.Atoi(st_scale_lv2_str)
	if err != nil || st_scale_lv2 <= 0 {
		return 0, 0, 0, 0, errors.New("师徒分润二级信息有误")
	}
	//查询工人抽成
	work_scale_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_PLATFORM_WORK_SCALE")
	work_scale, err = strconv.Atoi(work_scale_str)
	if err != nil || work_scale <= 0 {
		return 0, 0, 0, 0, errors.New("工人抽成信息有误")
	}
	// 返回 全部 第一级 第二级 工人抽成
	return st_scale_all, st_scale_lv1, st_scale_lv2, work_scale, nil
}

// 结算总订单
func (this *Accounts) AccountsOrders() error {
	//已经处理过状态无需处理

	//查询订单属于什么类型
	switch this.Orders.OrderType {
	case 0:
		// 普通服务订单 暂时这么处理 应该需要PlatformKey !!!!!!
		err := this.accountsOrdersOrSmService()
		if err != nil {
			return err
		}
		break
	case 1:
		// 普通服务订单 暂时这么处理 应该需要PlatformKey !!!!!!
		err := this.accountsOrdersOrSmService()
		if err != nil {
			return err
		}
		break
	case 2:
		err := this.accountsOrdersOrSmService()
		if err != nil {
			return err
		}
		break
	default:
		return errors.New("当前订单类型不可结算")
		break
	}
	return nil
}

// 上门服务订单类型订单结算
func (this *Accounts) accountsOrdersOrSmService() error {
	total_price := 0.00

	// 查询订单分润比
	share_all, share_lv1, share_lv2, share_scale, err := this.getShareScale()
	if err != nil {
		return err
	}

	//把当前订单金额抽出来
	all_order_price := this.Orders.Price
	// ！！！ 查询是否有关联订单
	link_count, link_orders, _ := models.GetLinkOrders(this.Orders.Id)
	if link_count > 0 {
		log.Println(link_count, link_orders)
		for _, v := range *link_orders {
			//注入金额进来
			all_order_price = all_order_price + v.Price
			//修改订单状态
			v.Status = 9
			models.UpdateUcenterOrdersById(&v)
		}
	}

	//抽出分润金额
	share_all_price := utils.Decimal(float64(all_order_price) * (float64(share_all) / 100))
	share_lv1_price := utils.Decimal(share_all_price * (float64(share_lv1) / 100))
	share_lv2_price := utils.Decimal(share_all_price * (float64(share_lv2) / 100))
	//查询平台抽成
	platform_scale_price := utils.Decimal(float64(all_order_price) * (float64(share_scale) / 100))
	log.Println("分润总收益", share_all_price, share_lv1_price, share_lv2_price, platform_scale_price)

	in_price := all_order_price - share_all_price - platform_scale_price //计算出工人实际获取金额
	log.Println("剩余订单金额", in_price)

	//订单日志
	m := new(models.UcenterOrdersLog)
	m.PlatformKey = "DDSM_CLEINT"
	m.Price = platform_scale_price
	m.Cuid = 0
	m.OrderId = int(this.Orders.Id)
	m.OrderNo = this.Orders.OrderNo
	m.ProjectId = int(this.SmUsersService.Id)
	m.Type = models.UCENTER_ORDERS_LOG_TYPE_PLATFORM
	m.Describe = "平台订单抽成"
	models.AddUcenterOrdersLog(m)
	total_price = total_price + platform_scale_price //插入给一级分享的钱

	//！！！先查询是否有一级use
	if this.Orders.UseShareLv1 > 0 {
		use_share_user_lv1, err := models.GetUcenterUsersById(int64(this.Orders.UseShareLv1))
		if err == nil {
			log.Println("强制一级分享用户", use_share_user_lv1)

			//处理强制一级分享收益
			//生成红包
			upack := new(models.MarketingPacket)
			upack.PlatformKey = "DDSM_CLIENT"
			upack.Cuid = int(use_share_user_lv1.Id)
			upack.SCuid = int(this.Users.Id)
			upack.Price = share_lv1_price
			upack.Title = "一级分享收益"
			upack.Describe = "收益到账" + strconv.FormatFloat(utils.Decimal(share_lv1_price), 'f', 2, 64)
			upack.Type = models.PACKET_TYPE_SHARE1
			upack.PacketNo = uuid.NewV4().String()
			models.GenerUserPacket(upack)
			// 订单日志
			m := new(models.UcenterOrdersLog)
			m.PlatformKey = "DDSM_CLEINT"
			m.Price = share_lv1_price
			m.Cuid = int(use_share_user_lv1.Id)
			m.OrderId = int(this.Orders.Id)
			m.OrderNo = this.Orders.OrderNo
			m.ProjectId = int(this.SmUsersService.Id)
			m.Type = models.UCENTER_ORDERS_LOG_TYPE_SHARE
			m.Describe = "一级分享收益"
			models.AddUcenterOrdersLog(m)
			total_price = total_price + share_lv1_price //插入给一级分享的钱
			//发布消息
			msg_utils := new(MessageCenterUtils)
			msg_utils.Price = m.Price
			msg_utils.Users = *use_share_user_lv1
			msg_utils.SmService = this.SmService
			msg_utils.SmUsersService = this.SmUsersService
			msg_utils.Orders = this.Orders
			msg_utils.ToUsers = this.Users

			msg_utils.PlatformKey = "DDSM_CLIENT"
			msg_utils.MessageKey = "SHARE_SY_DZ"
			msg_utils.PushMessage()
		}

		use_share_user_lv2, err := models.GetUcenterUsersById(int64(this.Orders.UseShareLv2))
		if err == nil {
			log.Println("强制二级分享用户", use_share_user_lv2)

			//处理第强制二级分享收益
			//生成红包
			upack := new(models.MarketingPacket)
			upack.PlatformKey = "DDSM_CLIENT"
			upack.Cuid = int(use_share_user_lv2.Id)
			upack.SCuid = int(use_share_user_lv2.Id)
			upack.Price = share_lv2_price
			upack.Title = "二级分享收益"
			upack.Describe = "收益到账" + strconv.FormatFloat(utils.Decimal(share_lv2_price), 'f', 2, 64)
			upack.Type = models.PACKET_TYPE_SHARE2
			upack.PacketNo = uuid.NewV4().String()
			models.GenerUserPacket(upack)
			// 订单日志
			m := new(models.UcenterOrdersLog)
			m.PlatformKey = "DDSM_CLEINT"
			m.Price = share_lv2_price
			m.Cuid = int(use_share_user_lv2.Id)
			m.OrderId = int(this.Orders.Id)
			m.OrderNo = this.Orders.OrderNo
			m.ProjectId = int(this.SmUsersService.Id)
			m.Type = models.UCENTER_ORDERS_LOG_TYPE_SHARE
			m.Describe = "二级分享收益"
			models.AddUcenterOrdersLog(m)
			total_price = total_price + share_lv2_price //插入给二级分享的钱
			//发布消息
			msg_utils := new(MessageCenterUtils)
			msg_utils.Price = m.Price
			msg_utils.Users = *use_share_user_lv2
			msg_utils.SmService = this.SmService
			msg_utils.SmUsersService = this.SmUsersService
			msg_utils.Orders = this.Orders
			msg_utils.ToUsers = this.Users

			msg_utils.PlatformKey = "DDSM_CLIENT"
			msg_utils.MessageKey = "SHARE_SY_DZ"
			msg_utils.PushMessage()
		}
	} else {
		//如果没有强制use 就按照默认的一二级分享
		//查询当前用户是否有一级分润
		share_user_lv1, err := models.CheckShareUser(this.Users.Id)
		if err == nil {
			log.Println("一级分享用户", share_user_lv1)

			//处理第一级分享收益
			//生成红包
			upack := new(models.MarketingPacket)
			upack.PlatformKey = "DDSM_CLIENT"
			upack.Cuid = int(share_user_lv1.Id)
			upack.SCuid = int(this.Users.Id)
			upack.Price = share_lv1_price
			upack.Title = "一级分享收益"
			upack.Describe = "收益到账" + strconv.FormatFloat(utils.Decimal(share_lv1_price), 'f', 2, 64)
			upack.Type = models.PACKET_TYPE_SHARE1
			upack.PacketNo = uuid.NewV4().String()
			models.GenerUserPacket(upack)
			// 订单日志
			m := new(models.UcenterOrdersLog)
			m.PlatformKey = "DDSM_CLEINT"
			m.Price = share_lv1_price
			m.Cuid = int(share_user_lv1.Id)
			m.OrderId = int(this.Orders.Id)
			m.OrderNo = this.Orders.OrderNo
			m.ProjectId = int(this.SmUsersService.Id)
			m.Type = models.UCENTER_ORDERS_LOG_TYPE_SHARE
			m.Describe = "一级分享收益"
			models.AddUcenterOrdersLog(m)
			total_price = total_price + share_lv1_price //插入给一级分享的钱
			//发布消息
			msg_utils := new(MessageCenterUtils)
			msg_utils.Price = m.Price
			msg_utils.Users = *share_user_lv1
			msg_utils.SmService = this.SmService
			msg_utils.SmUsersService = this.SmUsersService
			msg_utils.Orders = this.Orders
			msg_utils.ToUsers = this.Users

			msg_utils.PlatformKey = "DDSM_CLIENT"
			msg_utils.MessageKey = "SHARE_SY_DZ"
			msg_utils.PushMessage()

			//查询当前用户是否有二级分润
			share_user_lv2, err := models.CheckShareUser(share_user_lv1.Id)
			if err == nil {
				log.Println("二级分享用户", share_user_lv2)

				//处理第二级分享收益
				//生成红包
				upack := new(models.MarketingPacket)
				upack.PlatformKey = "DDSM_CLIENT"
				upack.Cuid = int(share_user_lv2.Id)
				upack.SCuid = int(share_user_lv1.Id)
				upack.Price = share_lv2_price
				upack.Title = "二级分享收益"
				upack.Describe = "收益到账" + strconv.FormatFloat(utils.Decimal(share_lv2_price), 'f', 2, 64)
				upack.Type = models.PACKET_TYPE_SHARE2
				upack.PacketNo = uuid.NewV4().String()
				models.GenerUserPacket(upack)
				// 订单日志
				m := new(models.UcenterOrdersLog)
				m.PlatformKey = "DDSM_CLEINT"
				m.Price = share_lv2_price
				m.Cuid = int(share_user_lv2.Id)
				m.OrderId = int(this.Orders.Id)
				m.OrderNo = this.Orders.OrderNo
				m.ProjectId = int(this.SmUsersService.Id)
				m.Type = models.UCENTER_ORDERS_LOG_TYPE_SHARE
				m.Describe = "二级分享收益"
				models.AddUcenterOrdersLog(m)
				total_price = total_price + share_lv2_price //插入给二级分享的钱
				//发布消息
				msg_utils := new(MessageCenterUtils)
				msg_utils.Price = m.Price
				msg_utils.Users = *share_user_lv2
				msg_utils.SmService = this.SmService
				msg_utils.SmUsersService = this.SmUsersService
				msg_utils.Orders = this.Orders
				msg_utils.ToUsers = *share_user_lv1

				msg_utils.PlatformKey = "DDSM_CLIENT"
				msg_utils.MessageKey = "SHARE_SY_DZ"
				msg_utils.PushMessage()
			}
		}
	}

	//查询结算金额
	all_use_price, err := models.SumUcenterOrdersLogOrOrderId(int(this.Orders.Id))
	if err != nil {
		return errors.New("结算金额查询失败")
	}
	log.Println("所有金额", all_use_price)

	platform_price := all_order_price - all_use_price

	//查询是否有合伙人
	partner, err := models.GetPartnerUsersById(int64(this.Orders.PartnerId))
	if err == nil {
		log.Println("有合伙人", partner)

		////直接分钱给合伙人
		m := new(models.UcenterOrdersLog)
		m.PlatformKey = "DDSM_CLEINT"
		m.Price = platform_price
		m.Cuid = int(partner.Id) // 合伙人的id
		m.OrderId = int(this.Orders.Id)
		m.OrderNo = this.Orders.OrderNo
		m.ProjectId = int(this.SmUsersService.Id)
		m.Type = models.UCENTER_ORDERS_LOG_TYPE_PARTNER
		m.Describe = "合伙人直接收益"
		models.AddUcenterOrdersLog(m)
		models.SetPartnerUsersOkMoney(partner.Id, platform_price)
	} else {

		// 平台直接收取
		m := new(models.UcenterOrdersLog)
		m.PlatformKey = "DDSM_CLEINT"
		m.Price = platform_price
		m.Cuid = 0
		m.OrderId = int(this.Orders.Id)
		m.OrderNo = this.Orders.OrderNo
		m.ProjectId = int(this.SmUsersService.Id)
		m.Type = models.UCENTER_ORDERS_LOG_TYPE_PLATFORM
		m.Describe = "平台最终收益"
		models.AddUcenterOrdersLog(m)
		log.Println("没有合伙人")
	}

	//设置结算
	models.SetUcenterOrdersAccount(this.Orders.Id, all_use_price)
	return nil

}

// 获取订单分润比例
func (this *Accounts) getShareScale() (share_scale_all int, share_scale_lv1 int, share_scale_lv2 int, platform_scale int, err error) {
	// 查询师徒总分润比
	share_scale_all_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_SHARE_SCALE_ALL")
	share_scale_all, err = strconv.Atoi(share_scale_all_str)
	if err != nil || share_scale_all <= 0 {
		return 0, 0, 0, 0, errors.New("分享收益信息有误")
	}

	// 查询师徒一级
	share_scale_lv1_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_SHARE_SCALE_LV1")
	share_scale_lv1, err = strconv.Atoi(share_scale_lv1_str)
	if err != nil || share_scale_lv1 <= 0 {
		return 0, 0, 0, 0, errors.New("分享收益一级信息有误")
	}

	// 查询师徒二级
	share_scale_lv2_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_SHARE_SCALE_LV2")
	share_scale_lv2, err = strconv.Atoi(share_scale_lv2_str)
	if err != nil || share_scale_lv2 <= 0 {
		return 0, 0, 0, 0, errors.New("分享收益二级信息有误")
	}
	//查询工人抽成
	platform_scale_str := models.GetSystemConfigValue("DDSM_CLIENT", "SM_PLATFORM_SCALE")
	platform_scale, err = strconv.Atoi(platform_scale_str)
	if err != nil || platform_scale <= 0 {
		return 0, 0, 0, 0, errors.New("工人抽成信息有误")
	}
	// 返回 全部 第一级 第二级 工人抽成
	return share_scale_all, share_scale_lv1, share_scale_lv2, platform_scale, nil
}
