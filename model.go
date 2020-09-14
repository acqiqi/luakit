package main

type Ucenter_optshare struct {
	Id           int    `orm:"id"`
	Cuid         int    `orm:"cuid"`           // 用户id
	PlatformKey  string `orm:"platform_key"`   // 平台key
	ShareLv1Cuid int    `orm:"share_lv1_cuid"` // 上级
	ShareLv2Cuid int    `orm:"share_lv2_cuid"` // 上级二级
	Level        int    `orm:"level"`          // 分享等级 0一级 1二级
	Status       int    `orm:"status"`         // 0不产生分润 1产生分润
	Flag         int    `orm:"flag"`           // 删除标识
	CreatedAt    string `orm:"created_at"`
	UpdatedAt    string `orm:"updated_at"`
}

type Admin_message struct {
	Id         int    `orm:"id"`
	UidReceive int    `orm:"uid_receive"` // 接收消息的用户id
	UidSend    int    `orm:"uid_send"`    // 发送消息的用户id
	Type       string `orm:"type"`        // 消息分类
	Content    string `orm:"content"`     // 消息内容
	Status     int    `orm:"status"`      // 状态
	CreateTime int    `orm:"create_time"` // 创建时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	ReadTime   int    `orm:"read_time"`   // 阅读时间
}

type Marketing_pw_orders struct {
	Id          int    `orm:"id"`
	Cuid        int    `orm:"cuid"`
	PlatformKey string `orm:"platform_key"`
	PwId        int    `orm:"pw_id"`
	PwPrizeId   int    `orm:"pw_prize_id"`
	ShareNum    int    `orm:"share_num"` // 需要分享多少人
	UseNum      int    `orm:"use_num"`   // 已经分享多少人
	Status      int    `orm:"status"`    // 0创建 1成功 -1未达成
	Flag        int    `orm:"flag"`      // 删除标识
	UpdatedAt   string `orm:"updated_at"`
	CreatedAt   string `orm:"created_at"`
}

type Marketing_seckill_packet_users struct {
	Id                  int    `orm:"id"`
	Cuid                int    `orm:"cuid"`
	PlatformKey         string `orm:"platform_key"`
	SeckillPacketLinkId int    `orm:"seckill_packet_link_id"`
	SeckillPacketId     int    `orm:"seckill_packet_id"`
	Flag                int    `orm:"flag"` // 删除标识
	CreatedAt           string `orm:"created_at"`
	UpdatedAt           string `orm:"updated_at"`
}

type Sm_show_comment struct {
	Id        int    `orm:"id"`
	Cuid      int    `orm:"cuid"`
	Nickname  string `orm:"nickname"`
	Avatar    string `orm:"avatar"`
	ShowId    int    `orm:"show_id"`
	Content   string `orm:"content"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Ucenter_action_log struct {
	Id        int    `orm:"id"`
	Cuid      int    `orm:"cuid"`
	Action    string `orm:"action"`  // 行为
	Content   string `orm:"content"` // 内容
	Title     string `orm:"title"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Admin_icon_list struct {
	Id     int    `orm:"id"`
	IconId int    `orm:"icon_id"` // 所属图标id
	Title  string `orm:"title"`   // 图标标题
	Class  string `orm:"class"`   // 图标类名
	Code   string `orm:"code"`    // 图标关键词
}

type Dispatch_log struct {
	Id         int    `orm:"id"`
	DispatchId int    `orm:"dispatch_id"`
	Cuid       int    `orm:"cuid"`
	Content    string `orm:"content"`
	Pics       string `orm:"pics"`
	Avatar     string `orm:"avatar"`
	Nickname   string `orm:"nickname"`
	Flag       int    `orm:"flag"` // 删除标识
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
}

type Marketing_packet struct {
	Id          int     `orm:"id"`
	Cuid        int     `orm:"cuid"`
	Price       float64 `orm:"price"`
	Title       string  `orm:"title"`
	Describe    string  `orm:"describe"` // 描述
	Status      int     `orm:"status"`   // 0正常 1使用 -1过期或者禁用
	Type        string  `orm:"type"`     // 来源渠道
	Flag        int     `orm:"flag"`     // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	PlatformKey string  `orm:"platform_key"`
	PacketNo    string  `orm:"packet_no"`
	SCuid       int     `orm:"s_cuid"` // 来源用户
	EndTime     int     `orm:"end_time"`
}

type Partner_tixian struct {
	Id          int     `orm:"id"`
	Price       float64 `orm:"price"`        // 提现金额
	FirstMoney  float64 `orm:"first_money"`  // 提现前金额
	Status      int     `orm:"status"`       // 状态0提现中 1成功 -1 失败
	TixianFlag  string  `orm:"tixian_flag"`  // 空或者bank 银行 wechat 微信钱包 alipay 支付宝钱包
	BankName    string  `orm:"bank_name"`    // 银行卡银行
	BankId      string  `orm:"bank_id"`      // 银行卡号
	BankUser    string  `orm:"bank_user"`    // 银行卡姓名
	BankAddress string  `orm:"bank_address"` // 开户行详细地址
	Flag        int     `orm:"flag"`         // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	TixianNo    string  `orm:"tixian_no"`
	PartnerId   int     `orm:"partner_id"`
}

type Ucenter_orders struct {
	Id             int     `orm:"id"`
	OrderNo        string  `orm:"order_no"`     // 订单编号
	PlatformKey    string  `orm:"platform_key"` // 平台key
	Cuid           int     `orm:"cuid"`         // ucenter id
	CouponKey      string  `orm:"coupon_key"`   // 优惠券key
	CouponId       int     `orm:"coupon_id"`    // 优惠券id
	CouponPrice    float64 `orm:"coupon_price"` // 优惠券抵扣金额
	CostPrice      float64 `orm:"cost_price"`   // 原价
	UnitPrice      float64 `orm:"unit_price"`   // 单价
	Price          float64 `orm:"price"`        // 现价 支付总价
	GoodsNum       float64 `orm:"goods_num"`    // 商品总数量
	PayType        int     `orm:"pay_type"`     // 支付类型 0线上 1线下
	PayPlatform    string  `orm:"pay_platform"` // 支付平台 wechat alipay wallet ...
	Status         int     `orm:"status"`       // 状态 0下单 1线下付款提交 2审核通过 7已接单 8已发货 9已结算或者已收货
	IsPay          int     `orm:"is_pay"`       // 是否支付 0否 1线下提交 2已支付
	IsUComment     int     `orm:"is_u_comment"` // 是否用户评论 0否 1是
	IsMComment     int     `orm:"is_m_comment"` // 是否商家评论 0否1是
	Flag           int     `orm:"flag"`         // 删除标识
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	PayTime        int     `orm:"pay_time"`       // 支付时间
	ServicePrice   float64 `orm:"service_price"`  // 服务费
	SharePrice     float64 `orm:"share_price"`    // 分享出去多少钱
	PlatformPrice  float64 `orm:"platform_price"` // 平台受益
	ShareLv1       float64 `orm:"share_lv1"`      // 一级分享
	ShareLv2       float64 `orm:"share_lv2"`      // 二级分享
	OptLv1         float64 `orm:"opt_lv1"`        // 一级其他分享
	OptLv2         float64 `orm:"opt_lv2"`        // 二级其他分享
	ShareLv1Cuid   int     `orm:"share_lv1_cuid"` // 一级分享用户id
	ShareLv2Cuid   int     `orm:"share_lv2_cuid"`
	OptLv1Cuid     int     `orm:"opt_lv1_cuid"`
	OptLv2Cuid     int     `orm:"opt_lv2_cuid"`
	IsAccounts     int     `orm:"is_accounts"` // 是否结算
	OrderType      int     `orm:"order_type"`  // 订单类型0常规购买订单  1团购 10VIP 20快捷支付
	PayNo          string  `orm:"pay_no"`      // 线上付款订单
	Describe       string  `orm:"describe"`    // 描述
	ProjectId      int     `orm:"project_id"`
	VipPriceD      float64 `orm:"vip_price_d"` // vip折扣多少钱
	Pics           string  `orm:"pics"`
	PartnerPrice   float64 `orm:"partner_price"`    // 合伙人金额
	PartnerId      int     `orm:"partner_id"`       // 合伙人id
	TotalUsePrice  float64 `orm:"total_use_price"`  // 全部占用金额
	UserPayDesc    string  `orm:"user_pay_desc"`    // 用户线下支付备注
	MainOrderId    int     `orm:"main_order_id"`    // 关联订单id
	TotalLinkPrice float64 `orm:"total_link_price"` // 所有关联订单金额
	UseShareLv2    int     `orm:"use_share_lv2"`    // 强制二级分享
	UseShareLv1    int     `orm:"use_share_lv1"`    // 强制一级分享
}

type Marketing_score struct {
	Id          int    `orm:"id"`
	Title       string `orm:"title"`        // 获取积分标题
	PlatformKey string `orm:"platform_key"` // 平台key
	Score       int    `orm:"score"`        // 获取积分
	Flag        int    `orm:"flag"`         // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	Type        int    `orm:"type"` // 0获取 1支出
	Cuid        int    `orm:"cuid"`
	ScoreRuleId int    `orm:"score_rule_id"` // 规则id
	OldScore    int    `orm:"old_score"`     // 获取之前的积分
	ProjectId   int    `orm:"project_id"`    // 积分使用的项目id
	ProjectType int    `orm:"project_type"`  // 积分使用类型
}

type Sm_service_gp_sku struct {
	Id            int    `orm:"id"`
	ServiceId     int    `orm:"service_id"`
	ServiceGpId   int    `orm:"service_gp_id"`
	ServiceAreaId int    `orm:"service_area_id"` // 区域id
	ServiceSkuId  int    `orm:"service_sku_id"`  // sku id
	Flag          int    `orm:"flag"`            // 删除标识
	CreatedAt     string `orm:"created_at"`
	UpdatedAt     string `orm:"updated_at"`
}

type Ucenter_tixian struct {
	Id          int     `orm:"id"`
	PlatformKey string  `orm:"platform_key"` // 平台key
	Cuid        int     `orm:"cuid"`         // ucenter id
	Price       float64 `orm:"price"`        // 提现金额
	FirstMoney  float64 `orm:"first_money"`  // 提现前金额
	Status      int     `orm:"status"`       // 状态0提现中 1成功 -1 失败
	TixianFlag  string  `orm:"tixian_flag"`  // 空或者bank 银行 wechat 微信钱包 alipay 支付宝钱包
	BankName    string  `orm:"bank_name"`    // 银行卡银行
	BankId      string  `orm:"bank_id"`      // 银行卡号
	BankUser    string  `orm:"bank_user"`    // 银行卡姓名
	BankAddress string  `orm:"bank_address"` // 开户行详细地址
	Flag        int     `orm:"flag"`         // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	TixianNo    string  `orm:"tixian_no"`
}

type Helper_cats struct {
	Id          int    `orm:"id"`
	PlatformKey string `orm:"platform_key"`
	CatName     string `orm:"cat_name"` // 分类名称
	CatLogo     string `orm:"cat_logo"` // 分类logo
	CatDesc     string `orm:"cat_desc"` // 分类描述
	IsShow      int    `orm:"is_show"`  // 是否显示
	Sort        int    `orm:"sort"`     // 排序
	Flag        int    `orm:"flag"`     // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
}

type Hksmart_specialty struct {
	Id        int    `orm:"id"`
	Name      string `orm:"name"`   // 专业名
	Desc      string `orm:"desc"`   // 描述
	TagNo     string `orm:"tag_no"` // 专业编号
	Flag      int    `orm:"flag"`
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Ucenter_collect struct {
	Id          int    `orm:"id"`
	ProjectType string `orm:"project_type"` // 项目类型
	PlatformKey string `orm:"platform_key"` // 平台key
	Cuid        int    `orm:"cuid"`         // 用户id
	ProjectId   int    `orm:"project_id"`   // 项目id
	Flag        int    `orm:"flag"`         // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
}

type System_pages struct {
	Id          int    `orm:"id"`
	Title       string `orm:"title"`    // 标题
	Describe    string `orm:"describe"` // 描述
	Content     string `orm:"content"`  // 富文本
	PlatformKey string `orm:"platform_key"`
	IsShow      int    `orm:"is_show"`
	Flag        int    `orm:"flag"` // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	Sort        int    `orm:"sort"` // 排序
	Logo        string `orm:"logo"`
	PageKey     string `orm:"page_key"` // 页面key
}

type Sm_quck_pay_orders struct {
	Id         int     `orm:"id"`
	Title      string  `orm:"title"`    // 标题
	Describe   string  `orm:"describe"` // 描述
	Logo       string  `orm:"logo"`
	Background string  `orm:"background"` // 背景图
	Price      float64 `orm:"price"`      // 支付金额
	Images     string  `orm:"images"`
	Flag       int     `orm:"flag"` // 删除标识
	CreatedAt  string  `orm:"created_at"`
	UpdatedAt  string  `orm:"updated_at"`
	IsShow     int     `orm:"is_show"`
	Cuid       int     `orm:"cuid"`
	QuckPayId  int     `orm:"quck_pay_id"`
	OrderId    int     `orm:"order_id"`
	IsPay      int     `orm:"is_pay"` // 0待支付 1支付
	Status     int     `orm:"status"` // 0待支付 1支付
}

type Sm_service_seckill struct {
	Id        int     `orm:"id"`
	ServiceId int     `orm:"service_id"` // 服务id
	Title     string  `orm:"title"`      // 秒杀标题
	Content   string  `orm:"content"`    // 富文本
	Price     float64 `orm:"price"`      // 价格
	CostPrice float64 `orm:"cost_price"` // 原价
	BeginTime string  `orm:"begin_time"` // 开始时间
	EndTime   string  `orm:"end_time"`   // 到期时间
	LimitType int     `orm:"limit_type"` // 拼团区域限制 0不限制sku相等下单  1限制sku相等下单
	Status    int     `orm:"status"`     // 0未开始  1已开始 10已结束
	Logo      string  `orm:"logo"`       // 秒杀Logo
	Banner    string  `orm:"banner"`     // 秒杀 Banner
	Num       int     `orm:"num"`        // 秒杀数量
	UseNum    int     `orm:"use_num"`    // 已秒数量
	IsTopic   int     `orm:"is_topic"`
	Flag      int     `orm:"flag"` // 删除标识
	CreatedAt string  `orm:"created_at"`
	UpdatedAt string  `orm:"updated_at"`
	SkuName   string  `orm:"sku_name"` // sku name
	IsShow    int     `orm:"is_show"`  // 是否显示
	Sort      int     `orm:"sort"`
}

type Sm_users struct {
	Id              int     `orm:"id"`
	Cuid            int     `orm:"cuid"`
	Longitude       float64 `orm:"longitude"`
	Latitude        float64 `orm:"latitude"`
	IsSetArea       int     `orm:"is_set_area"`
	AreaName        string  `orm:"area_name"`        // 当前区域的名称
	AreaLevel       int     `orm:"area_level"`       // 0省1市 2区 3小区
	AddreaaCheckId  int     `orm:"addreaa_check_id"` // 对应的id city 或者housingid
	Flag            int     `orm:"flag"`             // 删除标识
	CreatedAt       string  `orm:"created_at"`
	UpdatedAt       string  `orm:"updated_at"`
	IsNewUser       int     `orm:"is_new_user"`        // 是否是当前端新用户
	IsGjAuth        int     `orm:"is_gj_auth"`         // 是否工匠认证 0否 1认证中 2已认证 -1驳回
	GjIdcardTop     string  `orm:"gj_idcard_top"`      // 身份证正面
	GjIdcardBom     string  `orm:"gj_idcard_bom"`      // 身份证背面
	GjIdcardId      string  `orm:"gj_idcard_id"`       // 身份证号码
	GjIdcardEndTime string  `orm:"gj_idcard_end_time"` // 到期时间
	GjIdcardAddress string  `orm:"gj_idcard_address"`  // 身份证地址
	GjName          string  `orm:"gj_name"`            // 认证姓名
	GjAvatar        string  `orm:"gj_avatar"`          // 工匠头像
	GjProvinceId    int     `orm:"gj_province_id"`     // 省id
	GjCityId        int     `orm:"gj_city_id"`         // 市id
	GjAreaId        int     `orm:"gj_area_id"`         // 区id
	GjAreaLevel     int     `orm:"gj_area_level"`      // 区域等级 0省 1市 2区
	GjMobile        string  `orm:"gj_mobile"`          // 工匠认证手机号
	GjDesc          string  `orm:"gj_desc"`            // 工匠介绍
	GjSfCuid        int     `orm:"gj_sf_cuid"`         // 师傅id
	IsSf            int     `orm:"is_sf"`              // 是否师父
	GjErrMsg        string  `orm:"gj_err_msg"`         // 认证失败文案
	SfTime          string  `orm:"sf_time"`            // 出师日期
}

type System_banner struct {
	Id          int    `orm:"id"`
	Title       string `orm:"title"`    // 标题
	Describe    string `orm:"describe"` // 描述
	Banner      string `orm:"banner"`
	GotoType    int    `orm:"goto_type"` // 0不转跳 1单页 2path
	Path        string `orm:"path"`      // 路径
	Param       string `orm:"param"`     // 页面参数 或者id
	PlatformKey string `orm:"platform_key"`
	IsShow      int    `orm:"is_show"`
	Sort        int    `orm:"sort"`
	Flag        int    `orm:"flag"` // 删除标识
	UpdatedAt   string `orm:"updated_at"`
	CreatedAt   string `orm:"created_at"`
	BannerKey   string `orm:"banner_key"`
	VideoUrl    string `orm:"video_url"`
	HousingId   int    `orm:"housing_id"` // 对应的id city 或者housingid
}

type Ucenter_commission struct {
	Id            int    `orm:"id"`
	PlatformKey   string `orm:"platform_key"`    // 平台key
	AllShareScale int    `orm:"all_share_scale"` // 分享最大比例
	AllOptScale   int    `orm:"all_opt_scale"`   // 其他收益最大比例
	ShareLv1Scale int    `orm:"share_lv1_scale"` // 分享一级
	ShareLv2Scale int    `orm:"share_lv2_scale"` // 分享二级
	OptLv1Scale   int    `orm:"opt_lv1_scale"`   // 其他收益一级
	OptLv2Scale   int    `orm:"opt_lv2_scale"`   // 其他收益二级
	Status        int    `orm:"status"`          // 0停用 1启用
	Flag          int    `orm:"flag"`            // 删除标识
	CreatedAt     string `orm:"created_at"`
	UpdatedAt     string `orm:"updated_at"`
}

type Admin_attachment struct {
	Id         int    `orm:"id"`
	Uid        int    `orm:"uid"`         // 用户id
	Name       string `orm:"name"`        // 文件名
	Module     string `orm:"module"`      // 模块名，由哪个模块上传的
	Path       string `orm:"path"`        // 文件路径
	Thumb      string `orm:"thumb"`       // 缩略图路径
	Url        string `orm:"url"`         // 文件链接
	Mime       string `orm:"mime"`        // 文件mime类型
	Ext        string `orm:"ext"`         // 文件类型
	Size       int    `orm:"size"`        // 文件大小
	Md5        string `orm:"md5"`         // 文件md5
	Sha1       string `orm:"sha1"`        // sha1 散列值
	Driver     string `orm:"driver"`      // 上传驱动
	Download   int    `orm:"download"`    // 下载次数
	CreateTime int    `orm:"create_time"` // 上传时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	Sort       int    `orm:"sort"`        // 排序
	Status     int    `orm:"status"`      // 状态
	Width      int    `orm:"width"`       // 图片宽度
	Height     int    `orm:"height"`      // 图片高度
}

type Admin_icon struct {
	Id         int    `orm:"id"`
	Name       string `orm:"name"`        // 图标名称
	Url        string `orm:"url"`         // 图标css地址
	Prefix     string `orm:"prefix"`      // 图标前缀
	FontFamily string `orm:"font_family"` // 字体名
	CreateTime int    `orm:"create_time"` // 创建时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	Status     int    `orm:"status"`      // 状态
}

type Hksmart_users struct {
	Id         int    `orm:"id"`
	Cuid       int    `orm:"cuid"`
	IsVip      int    `orm:"is_vip"`       // 是否是vip
	VipEndTime int    `orm:"vip_end_time"` // vip到期时间
	IsPro      int    `orm:"is_pro"`       // 是否专业版
	ProEndTime int    `orm:"pro_end_time"` // 专业版到期时间
	Flag       int    `orm:"flag"`         // 删除标识
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
}

type Ucenter_platform struct {
	Id                 int    `orm:"id"`
	PlatformName       string `orm:"platform_name"`     // 平台名称
	PlatformUsername   string `orm:"platform_username"` // 平台账号
	PlatformPassword   string `orm:"platform_password"` // 平台密码
	PlatformAppType    int    `orm:"platform_app_type"` // 平台类型 0 api 1 web 5 wechat 6 alipay
	PlatformKey        string `orm:"platform_key"`      // 平台key 完全标识
	Status             int    `orm:"status"`            // 状态 1运行 0暂停维护 -1禁用
	Flag               int    `orm:"flag"`              // 删除标识
	CreatedAt          string `orm:"created_at"`
	UpdatedAt          string `orm:"updated_at"`
	Ak                 string `orm:"ak"`
	Sk                 string `orm:"sk"`
	PayName            string `orm:"pay_name"`
	PayAk              string `orm:"pay_ak"`
	PaySk              string `orm:"pay_sk"`
	PayNotifyUrl       string `orm:"pay_notify_url"` // 支付回调
	PayNotifyFunc      string `orm:"pay_notify_func"`
	PlatformSk         string `orm:"platform_sk"`
	MessageCallbackUrl string `orm:"message_callback_url"` // 消息中心回调地址
	PlatformSecret     string `orm:"platform_secret"`      // 平台secret
	UcCallbackUrl      string `orm:"uc_callback_url"`      // 用户中心回调地址
}

type Xinxi_jubao struct {
	Id          int    `orm:"id"`
	Cuid        int    `orm:"cuid"`
	ProjectId   int    `orm:"project_id"`
	ProjectType int    `orm:"project_type"`
	Content     string `orm:"content"`
	Flag        int    `orm:"flag"` // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
}

type Admin_access struct {
	Module string `orm:"module"` // 模型名称
	Group  string `orm:"group"`  // 权限分组标识
	Uid    int    `orm:"uid"`    // 用户id
	Nid    string `orm:"nid"`    // 授权节点id
	Tag    string `orm:"tag"`    // 分组标签
}

type Admin_hook_plugin struct {
	Id         int    `orm:"id"`
	Hook       string `orm:"hook"`        // 钩子id
	Plugin     string `orm:"plugin"`      // 插件标识
	CreateTime int    `orm:"create_time"` // 添加时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	Sort       int    `orm:"sort"`        // 排序
	Status     int    `orm:"status"`      // 状态
}

type Ucenter_users struct {
	Id             int     `orm:"id"`
	Username       string  `orm:"username"`        // 账号
	Password       string  `orm:"password"`        // 密码
	Mobile         string  `orm:"mobile"`          // 手机号
	Nickname       string  `orm:"nickname"`        // 昵称
	Email          string  `orm:"email"`           // 邮箱
	Avatar         string  `orm:"avatar"`          // 头像
	Gender         string  `orm:"gender"`          // 性别
	Status         int     `orm:"status"`          // 状态 0停用 1启用
	RoleType       int     `orm:"role_type"`       // 0无权限 1小区合伙人 2区县级合伙人 3城市合伙人 4省合伙人
	Score          int     `orm:"score"`           // 积分
	Money          float64 `orm:"money"`           // 余额
	OkMoney        float64 `orm:"ok_money"`        // 可提现余额
	NoMoney        float64 `orm:"no_money"`        // 不可提现金额
	LastLoginIp    string  `orm:"last_login_ip"`   // 最后一次登录ip
	LastLoginTime  int     `orm:"last_login_time"` // 最后一次登录时间戳
	LastLongitude  float64 `orm:"last_longitude"`  // 最后一次经度
	LastLatitude   float64 `orm:"last_latitude"`   // 最后一次维度
	ShareOne       int     `orm:"share_one"`       // 一级分享
	ShareTwo       int     `orm:"share_two"`       // 二级分享
	StOne          int     `orm:"st_one"`          // 一级师徒
	StTwo          int     `orm:"st_two"`          // 二级师徒
	UserKey        string  `orm:"user_key"`        // 用户注册唯一key
	WechatUnionid  string  `orm:"wechat_unionid"`  // 微信相关unionid
	RegType        int     `orm:"reg_type"`        // 注册类型 0手机号验证码 1账号
	RegSource      string  `orm:"reg_source"`      // 注册来源 例如 手机 微信 小程序
	Flag           int     `orm:"flag"`            // 删除标识
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	RegPlatformKey string  `orm:"reg_platform_key"` // 从哪个平台注册的
	IsVip          int     `orm:"is_vip"`           // 是否是vip
	BindUserinfo   int     `orm:"bind_userinfo"`    // 是否绑定用户信息
	VipEndTime     string  `orm:"vip_end_time"`     // vip到期时间
	IsPayPassword  int     `orm:"is_pay_password"`  // 是否填写支付密码
	PayPassword    string  `orm:"pay_password"`     // 支付密码
	OldManagerCuid int     `orm:"old_manager_cuid"`
	IsWork         int     `orm:"is_work"`         // 是否工人0否 1认证 2通过 -1拒绝
	IsAuth         int     `orm:"is_auth"`         // 是否实名认证 0 否 1审核 2通过 -1拒绝
	IdcardId       string  `orm:"idcard_id"`       // 身份证号
	IdcardBom      string  `orm:"idcard_bom"`      // 身份证背面
	IdcardTop      string  `orm:"idcard_top"`      // 身份证正面
	IdcardRen      string  `orm:"idcard_ren"`      // 手持证件
	IdcardAddress  string  `orm:"idcard_address"`  // 身份证地址
	IdcardNation   string  `orm:"idcard_nation"`   // 身份民族
	IdcardGender   string  `orm:"idcard_gender"`   // 性别
	IdcardProvince string  `orm:"idcard_province"` // 省
	IdcardCity     string  `orm:"idcard_city"`     // 市
	IdcardArea     string  `orm:"idcard_area"`     // 区
}

type Partner_users struct {
	Id        int     `orm:"id"`
	Username  string  `orm:"username"`  // 账号
	Password  string  `orm:"password"`  // 密码
	Mobile    string  `orm:"mobile"`    // 手机号
	Nickname  string  `orm:"nickname"`  // 昵称
	Email     string  `orm:"email"`     // 邮箱
	Gender    string  `orm:"gender"`    // 性别
	Status    int     `orm:"status"`    // 状态 0停用 1启用
	RoleType  int     `orm:"role_type"` // 0无权限 1小区合伙人 2区县级合伙人 3城市合伙人 4省合伙人
	Flag      int     `orm:"flag"`      // 删除标识
	CreatedAt string  `orm:"created_at"`
	UpdatedAt string  `orm:"updated_at"`
	Describe  string  `orm:"describe"`
	Money     float64 `orm:"money"` // 余额
	Avatar    string  `orm:"avatar"`
	About     string  `orm:"about"`
}

type Sm_service_appt_tpl struct {
	Id        int    `orm:"id"`
	TplName   string `orm:"tpl_name"` // 模板名称
	Describe  string `orm:"describe"` // 描述
	Flag      int    `orm:"flag"`     // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	PartnerId int    `orm:"partner_id"` // 合伙人id
}

type Admin_packet struct {
	Id         int    `orm:"id"`
	Name       string `orm:"name"`       // 数据包名
	Title      string `orm:"title"`      // 数据包标题
	Author     string `orm:"author"`     // 作者
	AuthorUrl  string `orm:"author_url"` // 作者url
	Version    string `orm:"version"`
	Tables     string `orm:"tables"`      // 数据表名
	CreateTime int    `orm:"create_time"` // 创建时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	Status     int    `orm:"status"`      // 状态
}

type Helper struct {
	Id          int    `orm:"id"`
	CatId       int    `orm:"cat_id"`
	Title       string `orm:"title"`   // 标题
	Content     string `orm:"content"` // 富文本
	Logo        string `orm:"logo"`    // logo
	Sort        int    `orm:"sort"`    // 排序
	Flag        int    `orm:"flag"`    // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	PlatformKey string `orm:"platform_key"`
	IsTopic     int    `orm:"is_topic"`
	IsShow      int    `orm:"is_show"` // 是否显示
}

type Marketing_pw struct {
	Id          int     `orm:"id"`
	PwName      string  `orm:"pw_name"`  // 转盘名称
	Describe    string  `orm:"describe"` // 描述
	PlatformKey string  `orm:"platform_key"`
	MinPrice    float64 `orm:"min_price"`  // 最低达成金额
	Background  string  `orm:"background"` // 背景图
	ShareNum    int     `orm:"share_num"`  // 需要分享人数量
	Circle      int     `orm:"circle"`     // 周期 天
	BeginTime   string  `orm:"begin_time"` // 开始时间
	IsStart     int     `orm:"is_start"`   // 是否开始
	Flag        int     `orm:"flag"`       // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	RuleContent string  `orm:"rule_content"` // 规则描述富文本
}

type Marketing_card_tpl_service struct {
	Id            int     `orm:"id"`
	CardTplId     int     `orm:"card_tpl_id"`
	ServiceId     int     `orm:"service_id"`     // 服务id
	ServiceSkuId  int     `orm:"service_sku_id"` // sku id
	CostPrice     float64 `orm:"cost_price"`     // 原价单价
	UnitPrice     float64 `orm:"unit_price"`     // 单价
	Num           int     `orm:"num"`            // 数量
	Flag          int     `orm:"flag"`           // 删除标识
	CreatedAt     string  `orm:"created_at"`
	UpdatedAt     string  `orm:"updated_at"`
	ServiceTitle  string  `orm:"service_title"` // 标题
	SkuName       string  `orm:"sku_name"`      // sku 名称
	ServiceAreaId int     `orm:"service_area_id"`
}

type Sm_quotation_tpl struct {
	Id        int    `orm:"id"`
	TplName   string `orm:"tpl_name"` // 模板名称
	Describe  string `orm:"describe"` // 备注
	TplTags   string `orm:"tpl_tags"`
	TplKey    string `orm:"tpl_key"` // 模板编号
	TotalNum  int    `orm:"total_num"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Sm_show_zan struct {
	Id        int    `orm:"id"`
	ShowId    int    `orm:"show_id"`
	Cuid      int    `orm:"cuid"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Xinxi_news_cats struct {
	Id        int    `orm:"id"`
	Name      string `orm:"name"`
	Logo      string `orm:"logo"`
	Pid       int    `orm:"pid"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	Sort      int    `orm:"sort"`
	IsShow    int    `orm:"is_show"`
}

type Admin_user struct {
	Id            int     `orm:"id"`
	Username      string  `orm:"username"`        // 用户名
	Nickname      string  `orm:"nickname"`        // 昵称
	Password      string  `orm:"password"`        // 密码
	Email         string  `orm:"email"`           // 邮箱地址
	EmailBind     int     `orm:"email_bind"`      // 是否绑定邮箱地址
	Mobile        string  `orm:"mobile"`          // 手机号码
	MobileBind    int     `orm:"mobile_bind"`     // 是否绑定手机号码
	Avatar        int     `orm:"avatar"`          // 头像
	Money         float64 `orm:"money"`           // 余额
	Score         int     `orm:"score"`           // 积分
	Role          int     `orm:"role"`            // 角色ID
	Group         int     `orm:"group"`           // 部门id
	SignupIp      int     `orm:"signup_ip"`       // 注册ip
	CreateTime    int     `orm:"create_time"`     // 创建时间
	UpdateTime    int     `orm:"update_time"`     // 更新时间
	LastLoginTime int     `orm:"last_login_time"` // 最后一次登录时间
	LastLoginIp   int     `orm:"last_login_ip"`   // 登录ip
	Sort          int     `orm:"sort"`            // 排序
	Status        int     `orm:"status"`          // 状态：0禁用，1启用
}

type Message_queue struct {
	Id              int    `orm:"id"`
	Cuid            int    `orm:"cuid"`
	MessageKey      string `orm:"message_key"`
	Title           string `orm:"title"`        // 标题
	Desc            string `orm:"desc"`         // 描述
	Content         string `orm:"content"`      // 内容
	MessageType     int    `orm:"message_type"` // 消息类型
	PathType        string `orm:"path_type"`    // 链接类型
	PathId          string `orm:"path_id"`      // 链接id
	IsFormId        int    `orm:"is_form_id"`   // 是否消息模板
	SmallTplId      string `orm:"small_tpl_id"` // 消息模板id
	SmallTplContent string `orm:"small_tpl_content"`
	SmallTplOpenid  string `orm:"small_tpl_openid"`
	IsSms           int    `orm:"is_sms"` // 是否发送短信
	Mobile          string `orm:"mobile"`
	SmsContent      string `orm:"sms_content"` // 短信内容
	IsEmail         int    `orm:"is_email"`    // 是否发邮件
	Email           string `orm:"email"`
	EmailTitle      string `orm:"email_title"`
	EmailContent    string `orm:"email_content"`
	IsSend          int    `orm:"is_send"`    // 是否发送
	MsgTplId        int    `orm:"msg_tpl_id"` // message tpl id
	CreatedAt       string `orm:"created_at"`
	UpdatedAt       string `orm:"updated_at"`
	DeletedAt       string `orm:"deleted_at"`
	Flag            int    `orm:"flag"`   // -1删除
	IsPop           int    `orm:"is_pop"` // 1已经出列
	PushData        string `orm:"push_data"`
	PlatformKey     string `orm:"platform_key"`
	IsUcId          int    `orm:"is_uc_id"`
	SmallTplPath    string `orm:"small_tpl_path"`
}

type Sm_service struct {
	Id              int     `orm:"id"`
	Title           string  `orm:"title"`          // 标题
	Logo            string  `orm:"logo"`           // 列表logo
	Icon            string  `orm:"icon"`           // 首页icon
	Banner          string  `orm:"banner"`         // banner列表 json
	CatsId          int     `orm:"cats_id"`        // 分类id
	Describe        string  `orm:"describe"`       // 描述
	Content         string  `orm:"content"`        // 富文本内容
	IsVideo         int     `orm:"is_video"`       // 是否显示视频banner
	VideoUrl        string  `orm:"video_url"`      // 视频url
	ShowNum         int     `orm:"show_num"`       // 显示数量
	PayNum          int     `orm:"pay_num"`        // 销售数量
	CollectNum      int     `orm:"collect_num"`    // 收藏数量
	ShareNum        int     `orm:"share_num"`      // 分享数量
	MinPrice        float64 `orm:"min_price"`      // 最小售价（用于限定最低单品购买以及显示）
	MaxPrice        float64 `orm:"max_price"`      // 最大售价，只用于显示
	MinCostPrice    float64 `orm:"min_cost_price"` // 最小销售原价
	MaxCostPrice    float64 `orm:"max_cost_price"` // 最高售价原价
	ProjectType     int     `orm:"project_type"`   // 0正常 1团购
	IsShow          int     `orm:"is_show"`        // 是否显示
	Status          int     `orm:"status"`         // 0停售 1正常销售  10预售
	Flag            int     `orm:"flag"`           // 删除标识
	CreatedAt       string  `orm:"created_at"`
	UpdatedAt       string  `orm:"updated_at"`
	IsDelete        int     `orm:"is_delete"`          // 是否远程删除
	IsTopic         int     `orm:"is_topic"`           // 是否推荐
	IsNew           int     `orm:"is_new"`             // 是否new
	Sort            int     `orm:"sort"`               // 排序
	AreaType        int     `orm:"area_type"`          // 限定区域范围等级 0全国 1省 2市 3区 4小区
	ProvinceId      int     `orm:"province_id"`        // 省id
	CityId          int     `orm:"city_id"`            // 市id
	AreaId          int     `orm:"area_id"`            // 区id
	HousingId       int     `orm:"housing_id"`         // 小区id
	IsVipPrice      int     `orm:"is_vip_price"`       // 是否有vip折扣
	IsClientUseAppt int     `orm:"is_client_use_appt"` // 判断是否需要用户端选择服务区间
}

type Hksmart_users_database struct {
	Id        int    `orm:"id"`
	Cuid      int    `orm:"cuid"`
	DId       int    `orm:"d_id"`       // 库id
	DOrderId  int    `orm:"d_order_id"` // 库购买订单号
	Status    int    `orm:"status"`     // 0未付款 1已支付 -1到期
	EndTime   int    `orm:"end_time"`   // 到期时间
	PayTime   int    `orm:"pay_time"`   // 支付时间
	Flag      int    `orm:"flag"`       // 删除标识
	UpdatedAt string `orm:"updated_at"`
	CreatedAt string `orm:"created_at"`
	PayType   int    `orm:"pay_type"` // 支付类型0月卡 1买断
	UdNo      string `orm:"ud_no"`
	BeginTime int    `orm:"begin_time"` // 激活时间
}

type Message struct {
	Id              int    `orm:"id"`
	Cuid            int    `orm:"cuid"`
	MessageKey      string `orm:"message_key"`
	Title           string `orm:"title"`        // 标题
	Desc            string `orm:"desc"`         // 描述
	Content         string `orm:"content"`      // 内容
	MessageType     int    `orm:"message_type"` // 消息类型
	PathType        string `orm:"path_type"`    // 链接类型
	PathId          string `orm:"path_id"`      // 链接id
	IsFormId        int    `orm:"is_form_id"`   // 是否消息模板
	SmallTplId      string `orm:"small_tpl_id"` // 消息模板id
	SmallTplContent string `orm:"small_tpl_content"`
	SmallTplOpenid  string `orm:"small_tpl_openid"`
	IsSms           int    `orm:"is_sms"` // 是否发送短信
	Mobile          string `orm:"mobile"`
	SmsContent      string `orm:"sms_content"` // 短信内容
	IsEmail         int    `orm:"is_email"`    // 是否发邮件
	Email           string `orm:"email"`
	EmailTitle      string `orm:"email_title"`
	EmailContent    string `orm:"email_content"`
	MsgTplId        int    `orm:"msg_tpl_id"`   // message tpl id
	PlatformKey     string `orm:"platform_key"` // 平台key
	PushData        string `orm:"push_data"`
	Flag            int    `orm:"flag"` // -1删除
	CreatedAt       string `orm:"created_at"`
	UpdatedAt       string `orm:"updated_at"`
	IsRead          int    `orm:"is_read"` // 是否已读
}

type Sm_service_gp struct {
	Id        int     `orm:"id"`
	ServiceId int     `orm:"service_id"`  // 服务id
	Title     string  `orm:"title"`       // 团购标题
	Content   string  `orm:"content"`     // 富文本
	AllNum    int     `orm:"all_num"`     // 满足多少成大团
	ShareNum  int     `orm:"share_num"`   // 满足多少成小团
	Type      int     `orm:"type"`        // 0满大团 1满小团  2满大小团 5阶梯团 10不管团多少都tm成
	Price     float64 `orm:"price"`       // 价格
	CostPrice float64 `orm:"cost_price"`  // 原价
	BeginTime string  `orm:"begin_time"`  // 开始时间
	EndTime   string  `orm:"end_time"`    // 到期时间
	LimitType int     `orm:"limit_type"`  // 拼团区域限制 0不限制sku相等下单  1限制sku相等下单
	UseMaxNum int     `orm:"use_max_num"` // 当前拼团一个用户最大可拼次数
	Status    int     `orm:"status"`      // 0未开始  1已开始 10已结束
	Flag      int     `orm:"flag"`        // 删除标识
	CreatedAt string  `orm:"created_at"`
	UpdatedAt string  `orm:"updated_at"`
	IsShow    int     `orm:"is_show"`     // 是否显示
	Logo      string  `orm:"logo"`        // 团购Logo
	Banner    string  `orm:"banner"`      // 团购Banner
	UseAllNum int     `orm:"use_all_num"` // 大团购数量
	IsTopic   int     `orm:"is_topic"`
	SkuName   string  `orm:"sku_name"` // sku name
	IsVip     int     `orm:"is_vip"`   // 是否需要会员购买
	Sort      int     `orm:"sort"`     // 排序
}

type Sm_quotation_link struct {
	Id          int     `orm:"id"`
	QuotationId int     `orm:"quotation_id"` // 报价id
	CatId       int     `orm:"cat_id"`       // 分类id
	CatName     string  `orm:"cat_name"`     // 分类名称
	Describe    string  `orm:"describe"`     // 描述
	TotalPrice  float64 `orm:"total_price"`
	TotalNum    int     `orm:"total_num"`  // 总数
	TableJson   string  `orm:"table_json"` // 数据集合
	LinkNo      string  `orm:"link_no"`    // 编号
	Flag        int     `orm:"flag"`       // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
}

type Admin_action struct {
	Id         int    `orm:"id"`
	Module     string `orm:"module"`      // 所属模块名
	Name       string `orm:"name"`        // 行为唯一标识
	Title      string `orm:"title"`       // 行为标题
	Remark     string `orm:"remark"`      // 行为描述
	Rule       string `orm:"rule"`        // 行为规则
	Log        string `orm:"log"`         // 日志规则
	Status     int    `orm:"status"`      // 状态
	CreateTime int    `orm:"create_time"` // 创建时间
	UpdateTime int    `orm:"update_time"` // 更新时间
}

type Hksmart_users_orders struct {
	Id        int     `orm:"id"`
	Cuid      int     `orm:"cuid"`
	OrderId   int     `orm:"order_id"`  // 订单id
	DId       int     `orm:"d_id"`      // database id
	Status    int     `orm:"status"`    // 0未付款 1付款 2使用 -1删除
	UseUdId   int     `orm:"use_ud_id"` // 使用之后关联的users database id
	OrderNo   string  `orm:"order_no"`  // 子订单id
	Flag      int     `orm:"flag"`      // 删除标识
	UpdatedAt string  `orm:"updated_at"`
	CreatedAt string  `orm:"created_at"`
	UnitPrice float64 `orm:"unit_price"` // 单价
	Price     float64 `orm:"price"`      // 总价
	Num       int     `orm:"num"`        // 多少月
	PayTime   int     `orm:"pay_time"`
}

type Sm_quotation_key struct {
	Id        int     `orm:"id"`
	Content   string  `orm:"content"`  // 工作内容
	Unit      string  `orm:"unit"`     // 计量单位
	Zl        float64 `orm:"zl"`       // 主料基础价
	Fl        float64 `orm:"fl"`       // 辅料基础价
	Work      float64 `orm:"work"`     // 工人基础价
	CatId     int     `orm:"cat_id"`   // 大类id
	Describe  string  `orm:"describe"` // 描述
	Flag      int     `orm:"flag"`     // 删除标识
	CreatedAt string  `orm:"created_at"`
	UpdatedAt string  `orm:"updated_at"`
}

type Hksmart_city struct {
	Id        int    `orm:"id"`
	Name      string `orm:"name"`   // 省市名
	Desc      string `orm:"desc"`   // 描述
	TagNo     string `orm:"tag_no"` // 标识编号
	Level     int    `orm:"level"`  // 0省 1市级
	Pid       int    `orm:"pid"`    // 0省 1省父id
	Flag      int    `orm:"flag"`
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Marketing_coupon_queue struct {
	Id          int     `orm:"id"`
	CouponTplId int     `orm:"coupon_tpl_id"`
	Logo        string  `orm:"logo"`
	Title       string  `orm:"title"`        // 优惠券名称
	PlatformKey string  `orm:"platform_key"` // 平台key
	Price       float64 `orm:"price"`        // 优惠金额/最大优惠金额
	FullPrice   float64 `orm:"full_price"`   // 满金额条件
	Type        int     `orm:"type"`         // 0满减 1全局折扣 2满折
	Zkb         int     `orm:"zkb"`          // 折扣比 0 100
	ProjectId   int     `orm:"project_id"`   // 对应项目id （上门端是skuid 或者serviceid）
	ProjectType int     `orm:"project_type"` // 对应项目类型。每个平台不同  （上门端0 sku 1 service）
	EndTime     int     `orm:"end_time"`     // 到期时间。被转换的时间戳
	Describe    string  `orm:"describe"`     // 描述
	Num         int     `orm:"num"`          // 最大发放数量
	QueueType   int     `orm:"queue_type"`   // 0用户列表 1省 2市 3区 10全国
	IsSend      int     `orm:"is_send"`      // 是否发放成功
	ProvinceId  int     `orm:"province_id"`  // 省id
	CityId      int     `orm:"city_id"`      // 市id
	AreaId      int     `orm:"area_id"`      // 区id
	Cuids       string  `orm:"cuids"`        // 用户数组，最大不可超过2万人
	Flag        int     `orm:"flag"`         // 删除标识
	UpdatedAt   string  `orm:"updated_at"`
	CreatedAt   string  `orm:"created_at"`
	QueueName   string  `orm:"queue_name"` // 队列名称
	UseNum      int     `orm:"use_num"`    // 已经发放数量
	IsRun       int     `orm:"is_run"`     // 是否在队列运行中
	ErrMsg      string  `orm:"err_msg"`
}

type Xinxi_jobs struct {
	Id                int     `orm:"id"`
	Pics              string  `orm:"pics"`
	Mobile            string  `orm:"mobile"`
	Title             string  `orm:"title"`
	Name              string  `orm:"name"`
	WorkType          string  `orm:"work_type"`
	LocationAddress   string  `orm:"location_address"`
	LocationArea      string  `orm:"location_area"`
	LocationLongitude float64 `orm:"location_longitude"`
	LocationLatitude  float64 `orm:"location_latitude"`
	LocationProvince  string  `orm:"location_province"`
	LocationCity      string  `orm:"location_city"`
	Content           string  `orm:"content"`
	Flag              int     `orm:"flag"` // 删除标识
	CreatedAt         string  `orm:"created_at"`
	UpdatedAt         string  `orm:"updated_at"`
	Cuid              int     `orm:"cuid"`
	IsVerify          int     `orm:"is_verify"`
}

type Log struct {
	Id        int    `orm:"id"`
	Content   string `orm:"content"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Sm_quotation_cats struct {
	Id          int    `orm:"id"`
	CatName     string `orm:"cat_name"`     // 分类名称
	DefaultSort int    `orm:"default_sort"` // 默认排序
	Describe    string `orm:"describe"`     // 描述
	Flag        int    `orm:"flag"`         // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
}

type Sm_service_sku struct {
	Id            int     `orm:"id"`
	ServiceId     int     `orm:"service_id"` // 服务id
	SkuName       string  `orm:"sku_name"`   // sku名称
	Describe      string  `orm:"describe"`   // 描述
	SkuLogo       string  `orm:"sku_logo"`   // logo 小图标
	Price         float64 `orm:"price"`      // 售价
	CostPrice     float64 `orm:"cost_price"` // 原价
	SpNum         int     `orm:"sp_num"`     // 单个人可以购买促销产品数量
	SpType        int     `orm:"sp_type"`    // 促销类型 0非促销 1新用户首单限定 2当前sku限定 3当前商品限定  10捆绑销售
	Sort          int     `orm:"sort"`       // 排序
	Flag          int     `orm:"flag"`       // 删除标识
	CreatedAt     string  `orm:"created_at"`
	UpdatedAt     string  `orm:"updated_at"`
	Stock         int     `orm:"stock"`           // 库存
	ApptTplId     int     `orm:"appt_tpl_id"`     // 预生成模板id
	ServiceAreaId int     `orm:"service_area_id"` // 区域id
	GjPrice       float64 `orm:"gj_price"`        // 工匠金额
	IsGjPrice     int     `orm:"is_gj_price"`     // 是否固定工匠金额
	SkuTagId      int     `orm:"sku_tag_id"`      // SKU标识
	IsShow        int     `orm:"is_show"`
}

type Partner_bank struct {
	Id          int    `orm:"id"`
	Uname       string `orm:"uname"`        // 姓名
	BankId      string `orm:"bank_id"`      // 银行卡号
	BankName    string `orm:"bank_name"`    // 银行名称
	BankAddress string `orm:"bank_address"` // 开户行
	Mobile      string `orm:"mobile"`       // 预留手机号
	IsDefault   int    `orm:"is_default"`   // 是否默认
	Cuid        int    `orm:"cuid"`
	Flag        int    `orm:"flag"` // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	PartnerId   int    `orm:"partner_id"`
}

type Ucenter_bank struct {
	Id          int    `orm:"id"`
	Uname       string `orm:"uname"`        // 姓名
	BankId      string `orm:"bank_id"`      // 银行卡号
	BankName    string `orm:"bank_name"`    // 银行名称
	BankAddress string `orm:"bank_address"` // 开户行
	Mobile      string `orm:"mobile"`       // 预留手机号
	IsDefault   int    `orm:"is_default"`   // 是否默认
	Cuid        int    `orm:"cuid"`
	Flag        int    `orm:"flag"` // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
}

type Ucenter_orders_log struct {
	Id          int     `orm:"id"`
	Cuid        int     `orm:"cuid"`
	PlatformKey string  `orm:"platform_key"`
	OrderId     int     `orm:"order_id"`
	OrderNo     string  `orm:"order_no"`
	Type        int     `orm:"type"`         // 0平台收益 1商家收益(工人) 2分润收入 3师徒收入 10合伙人收入
	ProjectType int     `orm:"project_type"` // 项目类型 0上门服务
	ProjectId   int     `orm:"project_id"`   // 对应项目id us_id
	Describe    string  `orm:"describe"`     // 描述
	Flag        int     `orm:"flag"`         // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	Price       float64 `orm:"price"` // 对应金额
}

type Dispatch struct {
	Id                  int     `orm:"id"`
	PartnerId           int     `orm:"partner_id"`    // 合伙人id
	Name                string  `orm:"name"`          // 项目名称
	AdrName             string  `orm:"adr_name"`      // lbs地址名称
	AdrLatitude         float64 `orm:"adr_latitude"`  // 纬度
	AdrLongitude        float64 `orm:"adr_longitude"` // 经度
	Address             string  `orm:"address"`       // 详细地址
	ZxType              string  `orm:"zx_type"`       // 装修类型
	ProductArea         string  `orm:"product_area"`  // 项目面积
	BeginTime           string  `orm:"begin_time"`
	EndTime             string  `orm:"end_time"`
	QuotedType          int     `orm:"quoted_type"` // 报价方式 0自主 1商家
	Price               float64 `orm:"price"`       // 自主报价金额
	IsSjs               int     `orm:"is_sjs"`      // 是否需要设计师
	Describe            string  `orm:"describe"`
	Pics                string  `orm:"pics"`
	Status              int     `orm:"status"` // -1审核没过 0待审核 1过审 5确定接单 9完工
	Cuid                int     `orm:"cuid"`   // 最终接单时 工人id
	DispatchNo          string  `orm:"dispatch_no"`
	OrderNo             string  `orm:"order_no"` // 订单id 预留
	Flag                int     `orm:"flag"`     // -1删除
	CreatedAt           string  `orm:"created_at"`
	UpdatedAt           string  `orm:"updated_at"`
	PartnerPrice        float64 `orm:"partner_price"`         // 合伙人实际扣除金额
	IsBack              int     `orm:"is_back"`               // 是否退钱，如果1说明已经把款项退回给合伙人
	UsersServiceId      int     `orm:"users_service_id"`      // 用户服务id
	EndEnrollTime       string  `orm:"end_enroll_time"`       // 报名截止时间
	CheckEndTime        int     `orm:"check_end_time"`        // 工人提交end信息的时间
	IsAdminComment      int     `orm:"is_admin_comment"`      // 是否后台评价
	AdminCommentStar    int     `orm:"admin_comment_star"`    // 评分
	AdminCommentPics    string  `orm:"admin_comment_pics"`    // 评价图片
	AdminCommentContent string  `orm:"admin_comment_content"` // 评价内容
	AdminCommentTime    int     `orm:"admin_comment_time"`
}

type Dispatch_enroll struct {
	Id         int     `orm:"id"`
	DispatchId int     `orm:"dispatch_id"`
	Cuid       int     `orm:"cuid"`
	Price      float64 `orm:"price"`
	Describe   string  `orm:"describe"`
	Pics       string  `orm:"pics"`
	Status     int     `orm:"status"` // 状态0正常 1已接单
	Flag       int     `orm:"flag"`   // -1删除
	CreatedAt  string  `orm:"created_at"`
	UpdatedAt  string  `orm:"updated_at"`
	IsSystem   int     `orm:"is_system"`  // 是否系统排单
	PartnerId  int     `orm:"partner_id"` // 合伙人id
}

type Marketing_pw_prize struct {
	Id         int     `orm:"id"`
	PwId       int     `orm:"pw_id"`       // 转盘id
	PrizeTitle string  `orm:"prize_title"` // 奖品名称
	PrizeType  int     `orm:"prize_type"`  // 上门端0红包 1优惠券 2服务 10冇
	ProjectId  int     `orm:"project_id"`  // 项目id
	ProjectTag string  `orm:"project_tag"` // 比如skutag  优惠券id
	Price      float64 `orm:"price"`       // 实际金额 - 红包类直接使用这个
	CostPrice  float64 `orm:"cost_price"`  // 原价
	Icon       string  `orm:"icon"`        // ICON
	Scale      int     `orm:"scale"`       // 获奖比例
	Flag       int     `orm:"flag"`
	CreatedAt  string  `orm:"created_at"`
	UpdatedAt  string  `orm:"updated_at"`
	Status     int     `orm:"status"`   // 0未开启 1开启
	Describe   string  `orm:"describe"` // 中奖描述
	Sort       int     `orm:"sort"`     // 排序
}

type Marketing_coupon_tpl struct {
	Id          int     `orm:"id"`
	Logo        string  `orm:"logo"`
	Title       string  `orm:"title"`        // 优惠券名称
	PlatformKey string  `orm:"platform_key"` // 平台key
	Price       float64 `orm:"price"`        // 优惠金额/最大优惠金额
	FullPrice   float64 `orm:"full_price"`   // 满金额条件
	Type        int     `orm:"type"`         // 0满减 1全局折扣 2满折
	Zkb         int     `orm:"zkb"`          // 折扣比 0 100
	ProjectId   int     `orm:"project_id"`   // 对应项目id （上门端是skuid 或者serviceid）
	ProjectType int     `orm:"project_type"` // 对应项目类型。每个平台不同  （上门端0 sku 1 service）
	EndTime     int     `orm:"end_time"`     // 到期时间 天
	Flag        int     `orm:"flag"`         // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	Describe    string  `orm:"describe"` // 描述
	TplKey      string  `orm:"tpl_key"`  // 模板KEY
}

type Sm_service_log struct {
	Id             int    `orm:"id"`
	ServiceId      int    `orm:"service_id"`
	ServiceSkuId   int    `orm:"service_sku_id"`   // sku id
	ServiceAreaId  int    `orm:"service_area_id"`  // 区域id
	UsersServiceId int    `orm:"users_service_id"` // 用户服务表id
	Cuid           int    `orm:"cuid"`
	Content        string `orm:"content"`
	Pics           string `orm:"pics"`
	Avatar         string `orm:"avatar"`   // 备份头像
	Nickname       string `orm:"nickname"` // 备份昵称
	Flag           int    `orm:"flag"`     // 删除标识
	CreatedAt      string `orm:"created_at"`
	UpdatedAt      string `orm:"updated_at"`
}

type Ucenter_users_work struct {
	Id                int     `orm:"id"`
	Cuid              int     `orm:"cuid"`
	Name              string  `orm:"name"`     // 姓名
	CardId            string  `orm:"card_id"`  // 身份证号
	CardTop           string  `orm:"card_top"` // 身份证正面
	CardRen           string  `orm:"card_ren"` // 手持
	Address           string  `orm:"address"`  // 身份证地址
	Nation            string  `orm:"nation"`   // 民族
	Gender            string  `orm:"gender"`   // 性别
	Province          string  `orm:"province"` // 省
	City              string  `orm:"city"`     // 市
	Area              string  `orm:"area"`     // 区
	IsAuth            int     `orm:"is_auth"`
	Flag              int     `orm:"flag"` // 删除标识
	CreatedAt         string  `orm:"created_at"`
	UpdatedAt         string  `orm:"updated_at"`
	WorkType          string  `orm:"work_type"`         // 工人类型
	LocationAddress   string  `orm:"location_address"`  // 定位地址
	LocationProvince  string  `orm:"location_province"` // 定位省
	LocationCity      string  `orm:"location_city"`     // 定位市
	LocationArea      string  `orm:"location_area"`     // 定位区
	LocationLongitude float64 `orm:"location_longitude"`
	LocationLatitude  float64 `orm:"location_latitude"`
	OfficeMobile      string  `orm:"office_mobile"`
	Miaoshu           string  `orm:"miaoshu"`
	GongzuoTime       string  `orm:"gongzuo_time"`
	IsGongren         int     `orm:"is_gongren"`
	Birth             string  `orm:"birth"` // 生日
	IsGongzuo         int     `orm:"is_gongzuo"`
	GongrenType       string  `orm:"gongren_type"`
	Home              string  `orm:"home"`
}

type Admin_role struct {
	Id            int    `orm:"id"`             // 角色id
	Pid           int    `orm:"pid"`            // 上级角色
	Name          string `orm:"name"`           // 角色名称
	Description   string `orm:"description"`    // 角色描述
	MenuAuth      string `orm:"menu_auth"`      // 菜单权限
	Sort          int    `orm:"sort"`           // 排序
	CreateTime    int    `orm:"create_time"`    // 创建时间
	UpdateTime    int    `orm:"update_time"`    // 更新时间
	Status        int    `orm:"status"`         // 状态
	Access        int    `orm:"access"`         // 是否可登录后台
	DefaultModule int    `orm:"default_module"` // 默认访问模块
}

type Admin_users struct {
	Id        int    `orm:"id"`
	Username  string `orm:"username"` // 账号
	Password  string `orm:"password"` // 密码
	Mobile    string `orm:"mobile"`   // 手机号
	Nickname  string `orm:"nickname"` // 昵称
	Email     string `orm:"email"`    // 邮箱
	Gender    string `orm:"gender"`   // 性别
	Status    int    `orm:"status"`   // 状态 0停用 1启用
	Flag      int    `orm:"flag"`     // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Face_device struct {
	Id            int     `orm:"id"`
	Name          string  `orm:"name"`           // 设备名字
	Describe      string  `orm:"describe"`       // 设备描述
	Address       string  `orm:"address"`        // 详细地址
	Mac           string  `orm:"mac"`            // 设备mac
	Sign          string  `orm:"sign"`           // 设备sign
	Cmd           int     `orm:"cmd"`            // 荔枝默认14
	RelayOpenVol  int     `orm:"relay_open_vol"` // 继电器极性
	ChkFacePose   int     `orm:"chk_face_pose"`  // 正脸判断使能
	FaceThreshold int     `orm:"face_threshold"` // 不传则表示不修改。识别阈值，范围1-100。建议80
	LedBright     float64 `orm:"led_bright"`     // 不传则表示不修改。补光灯亮度调节，如0.1， 0.01
	RelayOpenS    int     `orm:"relay_open_s"`   // 不传则表示不修改。继电器开光延时，可选1，2，3等
	Flag          int     `orm:"flag"`           // 删除标识
	CreatedAt     string  `orm:"created_at"`
	UpdatedAt     string  `orm:"updated_at"`
	IsDelete      int     `orm:"is_delete"`    // 是否远程删除
	NetworkType   int     `orm:"network_type"` // 0 wifi 1网口 2 4G
	NetworkData   string  `orm:"network_data"` // 联网配置参数存储
}

type Message_tpl struct {
	Id              int    `orm:"id"`
	MessageKey      string `orm:"message_key"` // 唯一标识
	Title           string `orm:"title"`
	Desc            string `orm:"desc"`
	Content         string `orm:"content"`
	AppType         int    `orm:"app_type"` // 0用户端 1商家端
	IsMsg           int    `orm:"is_msg"`   // 是否发送消息
	MessageType     string `orm:"message_type"`
	PathType        string `orm:"path_type"`         // 路径类型
	PathId          string `orm:"path_id"`           // 路径id 或者路径
	IsFormId        int    `orm:"is_form_id"`        // 是否使用小程序模板id推送
	SmallTplId      string `orm:"small_tpl_id"`      // 小程序模板id
	SmallTplContent string `orm:"small_tpl_content"` // 小程序模板内容 json
	IsSms           int    `orm:"is_sms"`            // 是否发送短信
	SmsContent      string `orm:"sms_content"`       // 短信内容
	IsEmail         int    `orm:"is_email"`          // 是否发送短信
	EmailTitle      string `orm:"email_title"`
	EmailContent    string `orm:"email_content"`
	CreatedAt       string `orm:"created_at"`
	UpdatedAt       string `orm:"updated_at"`
	DeletedAt       string `orm:"deleted_at"`
	Flag            int    `orm:"flag"`     // -1删除
	IsUcId          int    `orm:"is_uc_id"` // 是否使用用户平台
	PlatformKey     string `orm:"platform_key"`
	SmallTplPath    string `orm:"small_tpl_path"`
}

type System_log struct {
	Id        int    `orm:"id"`
	Content   string `orm:"content"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Admin_menu struct {
	Id         int    `orm:"id"`
	Pid        int    `orm:"pid"`         // 上级菜单id
	Module     string `orm:"module"`      // 模块名称
	Title      string `orm:"title"`       // 菜单标题
	Icon       string `orm:"icon"`        // 菜单图标
	UrlType    string `orm:"url_type"`    // 链接类型（link：外链，module：模块）
	UrlValue   string `orm:"url_value"`   // 链接地址
	UrlTarget  string `orm:"url_target"`  // 链接打开方式：_blank,_self
	OnlineHide int    `orm:"online_hide"` // 网站上线后是否隐藏
	CreateTime int    `orm:"create_time"` // 创建时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	Sort       int    `orm:"sort"`        // 排序
	SystemMenu int    `orm:"system_menu"` // 是否为系统菜单，系统菜单不可删除
	Status     int    `orm:"status"`      // 状态
	Params     string `orm:"params"`      // 参数
}

type Admin_log struct {
	Id         int    `orm:"id"`          // 主键
	ActionId   int    `orm:"action_id"`   // 行为id
	UserId     int    `orm:"user_id"`     // 执行用户id
	ActionIp   int    `orm:"action_ip"`   // 执行行为者ip
	Model      string `orm:"model"`       // 触发行为的表
	RecordId   int    `orm:"record_id"`   // 触发行为的数据id
	Remark     string `orm:"remark"`      // 日志备注
	Status     int    `orm:"status"`      // 状态
	CreateTime int    `orm:"create_time"` // 执行行为的时间
}

type Sm_service_sku_tags struct {
	Id        int    `orm:"id"`
	TagName   string `orm:"tag_name"` // 标签名称
	Describe  string `orm:"describe"` // 描述
	IsShow    int    `orm:"is_show"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	PartnerId int    `orm:"partner_id"` // 合伙人id
}

type Ucenter_price_log struct {
	Id          int     `orm:"id"`
	Cuid        int     `orm:"cuid"`
	Price       float64 `orm:"price"`        // 金额
	PlatformKey string  `orm:"platform_key"` // 平台key
	Flag        int     `orm:"flag"`         // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
}

type Sm_quck_pay struct {
	Id             int     `orm:"id"`
	Title          string  `orm:"title"`    // 标题
	Describe       string  `orm:"describe"` // 描述
	Logo           string  `orm:"logo"`
	Background     string  `orm:"background"` // 背景图
	Price          float64 `orm:"price"`      // 支付金额
	Images         string  `orm:"images"`
	Flag           int     `orm:"flag"` // 删除标识
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	IsShow         int     `orm:"is_show"`
	OrderId        int     `orm:"order_id"`
	IsPay          int     `orm:"is_pay"` // 0待支付 1支付
	QuckPayNo      string  `orm:"quck_pay_no"`
	PayCuid        int     `orm:"pay_cuid"`         // 支付用户id
	UsersServiceId int     `orm:"users_service_id"` // 用户服务表id
	UsersServiceNo string  `orm:"users_service_no"` // 用户订单编号
	PartnerId      int     `orm:"partner_id"`       // 合伙人id
	PayType        int     `orm:"pay_type"`         // 支付类型 0线上 1线下
	IsBindUs       int     `orm:"is_bind_us"`       // 是否绑定订单
	LinkOrderId    int     `orm:"link_order_id"`    // 关联总订单id
}

type System_config struct {
	Id          int    `orm:"id"`
	K           string `orm:"k"`            // key
	V           string `orm:"v"`            // value
	Title       string `orm:"title"`        // 标题
	Describe    string `orm:"describe"`     // 描述
	PlatformKey string `orm:"platform_key"` // platform_key  如果全局写ALL
	Flag        int    `orm:"flag"`         // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
}

type Sm_service_cats struct {
	Id          int    `orm:"id"`
	CatName     string `orm:"cat_name"`      // 分类名称
	CatLogo     string `orm:"cat_logo"`      // 分类logo
	CatHomeLogo string `orm:"cat_home_logo"` // 用于首页显示的logo
	CatTags     string `orm:"cat_tags"`      // 分类标签
	Sort        int    `orm:"sort"`          // 排序 倒序
	IsShow      int    `orm:"is_show"`
	Flag        int    `orm:"flag"` // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	Pid         int    `orm:"pid"`
	Level       int    `orm:"level"`    // 层级
	IsTopic     int    `orm:"is_topic"` // 是否首页置顶
	BgColor     string `orm:"bg_color"`
}

type Sm_users_comment struct {
	Id             int    `orm:"id"`
	ServiceId      int    `orm:"service_id"`       // 服务id
	ApptId         int    `orm:"appt_id"`          // 服务区间
	ServiceSkuId   int    `orm:"service_sku_id"`   // 服务sku
	OrderId        int    `orm:"order_id"`         // 订单id
	UsersServiceId int    `orm:"users_service_id"` // 用户服务表id
	ServiceAreaId  int    `orm:"service_area_id"`  // 区域id
	Cuid           int    `orm:"cuid"`
	Type           int    `orm:"type"` // 0用户 1商家
	Content        string `orm:"content"`
	Pics           string `orm:"pics"`
	VideoUrl       string `orm:"video_url"`
	Flag           int    `orm:"flag"` // 删除标识
	CreatedAt      string `orm:"created_at"`
	UpdatedAt      string `orm:"updated_at"`
	Avatar         string `orm:"avatar"`   // 备份头像
	Nickname       string `orm:"nickname"` // 备份昵称
	Star           int    `orm:"star"`     // 星 1非常不满意 2不满意 3一般 4满意 5 非常满意
	Tags           string `orm:"tags"`
	IsOld          int    `orm:"is_old"` // 老数据
}

type Ucenter_openid struct {
	Id          int    `orm:"id"`
	Cuid        int    `orm:"cuid"`         // ucenter id
	PlatformKey string `orm:"platform_key"` // 平台key
	Type        string `orm:"type"`         // 类型 wechat ali app
	Flag        int    `orm:"flag"`         // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	Openid      string `orm:"openid"`
}

type Marketing_coupon struct {
	Id            int     `orm:"id"`
	CouponTplId   int     `orm:"coupon_tpl_id"`   // 模板id
	CouponQueueId int     `orm:"coupon_queue_id"` // 批量发送id 如果有
	SourceType    int     `orm:"source_type"`     // 0后台发送 1新用户注册 2事件触发
	Cuid          int     `orm:"cuid"`
	Logo          string  `orm:"logo"`
	Title         string  `orm:"title"`        // 优惠券名称
	PlatformKey   string  `orm:"platform_key"` // 平台key
	Price         float64 `orm:"price"`        // 优惠金额/最大优惠金额
	FullPrice     float64 `orm:"full_price"`   // 满金额条件
	Type          int     `orm:"type"`         // 0满减 1全局折扣 2满折
	Zkb           int     `orm:"zkb"`          // 折扣比 0 100
	ProjectId     int     `orm:"project_id"`   // 对应项目id （上门端是skuid 或者serviceid）
	ProjectType   int     `orm:"project_type"` // 对应项目类型。每个平台不同  （上门端0 sku 1 service）
	EndTime       int     `orm:"end_time"`     // 到期时间。被转换的时间戳
	Describe      string  `orm:"describe"`     // 描述
	IsUse         int     `orm:"is_use"`       // 是否使用
	OrderId       int     `orm:"order_id"`     // 使用后订单id
	Flag          int     `orm:"flag"`         // 删除标识
	UpdatedAt     string  `orm:"updated_at"`
	CreatedAt     string  `orm:"created_at"`
	CouponKey     string  `orm:"coupon_key"` // 优惠券唯一key
}

type Sm_service_appt struct {
	Id            int    `orm:"id"`
	ApptItemTplId int    `orm:"appt_item_tpl_id"` // 对应服务时间名
	BeginTime     string `orm:"begin_time"`       // 开始时间
	EndTime       string `orm:"end_time"`         // 结束时间
	Title         string `orm:"title"`            // 用于前台展示
	Sort          int    `orm:"sort"`             // 排序
	IsShow        int    `orm:"is_show"`          // 是否显示
	Num           int    `orm:"num"`              // 可预约次数
	Status        int    `orm:"status"`           // 0待预约 1已约满  9暂停预约 -1禁止预约
	UseNum        int    `orm:"use_num"`          // 已预约人数
	AreaType      int    `orm:"area_type"`        // 限定区域范围等级 0全国 1省 2市 3区 4小区
	ProvinceId    int    `orm:"province_id"`      // 省id
	CityId        int    `orm:"city_id"`          // 市id
	AreaId        int    `orm:"area_id"`          // 区id
	HousingId     int    `orm:"housing_id"`       // 小区id
	CreatedAt     string `orm:"created_at"`
	ServiceId     int    `orm:"service_id"`     // 服务id
	ServiceSkuId  int    `orm:"service_sku_id"` // sku id
	UpdatedAt     string `orm:"updated_at"`
	Flag          int    `orm:"flag"`        // 删除标识
	ApptDate      string `orm:"appt_date"`   // 预约日期
	ApptTplId     int    `orm:"appt_tpl_id"` // 预生成模板id
	ServiceAreaId int    `orm:"service_area_id"`
}

type Sm_service_appt_item_tpl struct {
	Id        int    `orm:"id"`
	ApptTplId int    `orm:"appt_tpl_id"` // 模板tplid
	BeginTime string `orm:"begin_time"`  // 开始时间
	EndTime   string `orm:"end_time"`    // 结束时间
	Title     string `orm:"title"`       // 用于前台展示
	Sort      int    `orm:"sort"`        // 排序
	IsShow    int    `orm:"is_show"`     // 是否显示
	Num       int    `orm:"num"`         // 可预约次数
	Flag      int    `orm:"flag"`        // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Sm_service_area struct {
	Id         int    `orm:"id"`
	Title      string `orm:"title"`
	ServiceId  int    `orm:"service_id"`  // 服务id
	AreaType   int    `orm:"area_type"`   // 0 全国 1全省 2市 3区 4小区
	ProvinceId int    `orm:"province_id"` // 省id
	CityId     int    `orm:"city_id"`     // 市id
	AreaId     int    `orm:"area_id"`     // 区id
	HousingId  int    `orm:"housing_id"`  // 小区id
	Flag       int    `orm:"flag"`        // 删除标识
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
}

type Sm_users_cart struct {
	Id            int     `orm:"id"`
	ServiceId     int     `orm:"service_id"`     // 服务id
	ServiceSkuId  int     `orm:"service_sku_id"` // sku id
	ApptId        int     `orm:"appt_id"`        // 对应服务时间id
	Num           int     `orm:"num"`            // 购物车数量
	Cuid          int     `orm:"cuid"`           // 用户id
	UnitPrice     float64 `orm:"unit_price"`     // 售价
	CostPrice     float64 `orm:"cost_price"`     // 原价
	ServiceLogo   string  `orm:"service_logo"`
	ServiceTitle  string  `orm:"service_title"` // 标题
	ApptTitle     string  `orm:"appt_title"`
	Flag          int     `orm:"flag"` // 删除标识
	CreatedAt     string  `orm:"created_at"`
	UpdatedAt     string  `orm:"updated_at"`
	AddressId     string  `orm:"address_id"`      // 用户下单的地址id
	ServiceAreaId int     `orm:"service_area_id"` // 区域id
	NotAppt       int     `orm:"not_appt"`        // 是否暂时不选择服务区间 0否 1是
	GjPrice       float64 `orm:"gj_price"`        // 工匠金额
	IsGjPrice     int     `orm:"is_gj_price"`     // 是否固定工匠金额
}

type Admin_plugin struct {
	Id          int    `orm:"id"`
	Name        string `orm:"name"`        // 插件名称
	Title       string `orm:"title"`       // 插件标题
	Icon        string `orm:"icon"`        // 图标
	Description string `orm:"description"` // 插件描述
	Author      string `orm:"author"`      // 作者
	AuthorUrl   string `orm:"author_url"`  // 作者主页
	Config      string `orm:"config"`      // 配置信息
	Version     string `orm:"version"`     // 版本号
	Identifier  string `orm:"identifier"`  // 插件唯一标识符
	Admin       int    `orm:"admin"`       // 是否有后台管理
	CreateTime  int    `orm:"create_time"` // 安装时间
	UpdateTime  int    `orm:"update_time"` // 更新时间
	Sort        int    `orm:"sort"`        // 排序
	Status      int    `orm:"status"`      // 状态
}

type Marketing_seckill_packet struct {
	Id         int    `orm:"id"`
	Type       int    `orm:"type"`   // 0红包(不需要付费)
	Title      string `orm:"title"`  // 标题
	Status     int    `orm:"status"` // 0未开始 1开始 2结束
	Flag       int    `orm:"flag"`   // 删除标识
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
	Background string `orm:"background"` // 背景图
}

type Partner_accounts struct {
	Id         int     `orm:"id"`
	PartnerId  int     `orm:"partner_id"`  // 合伙人id
	OrderId    int     `orm:"order_id"`    // 订单id
	Price      float64 `orm:"price"`       // 到账金额
	AreaType   int     `orm:"area_type"`   // 0 全国 1全省 2市 3区 4小区
	HousingId  int     `orm:"housing_id"`  // 小区id
	ProvinceId int     `orm:"province_id"` // 省id
	CityId     int     `orm:"city_id"`     // 市id
	AreaId     int     `orm:"area_id"`     // 区id
	Flag       int     `orm:"flag"`        // 删除标识
	CreatedAt  string  `orm:"created_at"`
	UpdatedAt  string  `orm:"updated_at"`
}

type Sm_users_service_gp struct {
	Id             int    `orm:"id"`
	ServiceId      int    `orm:"service_id"`       // 服务id
	ServiceGpId    int    `orm:"service_gp_id"`    // 服务拼团id
	ServiceSkuId   int    `orm:"service_sku_id"`   // 服务sku
	OrderId        int    `orm:"order_id"`         // 订单id
	UsersServiceId int    `orm:"users_service_id"` // 用户服务表id
	Cuid           int    `orm:"cuid"`
	Status         int    `orm:"status"` // -1结束 0未支付 1支付 2完成
	Flag           int    `orm:"flag"`   // 删除标识
	CreatedAt      string `orm:"created_at"`
	UpdatedAt      string `orm:"updated_at"`
	Num            int    `orm:"num"`
}

type Ucenter_feedback struct {
	Id          int    `orm:"id"`
	Uname       string `orm:"uname"`
	Mobile      string `orm:"mobile"`
	Content     string `orm:"content"`
	Cuid        int    `orm:"cuid"`
	Flag        int    `orm:"flag"` // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	PlatformKey string `orm:"platform_key"`
	IsBlack     int    `orm:"is_black"` // 是否是黑名单提交
}

type Address_city struct {
	Id        int     `orm:"id"`
	Name      string  `orm:"name"`      // 名称
	Pid       int     `orm:"pid"`       // 父类
	Level     int     `orm:"level"`     // 0 省级 1城市 2区县级
	IsShow    int     `orm:"is_show"`   // 是否显示
	Flag      int     `orm:"flag"`      // 删除标识
	Longitude float64 `orm:"longitude"` // 经度
	Latitude  float64 `orm:"latitude"`  // 维度
	Scale     int     `orm:"scale"`     // 分润
	CreatedAt string  `orm:"created_at"`
	UpdatedAt string  `orm:"updated_at"`
}

type Sm_service_gp_level struct {
	Id          int     `orm:"id"`
	ServiceGpId int     `orm:"service_gp_id"` // 团购id
	Level       int     `orm:"level"`         // 等级
	Price       float64 `orm:"price"`         // 达成价格
	Flag        int     `orm:"flag"`          // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	Title       string  `orm:"title"` // 标题
	Num         int     `orm:"num"`   // 最低达成人数
}

type Sm_show struct {
	Id             int     `orm:"id"`
	HousingId      int     `orm:"housing_id"`  // 小区id
	ProvinceId     int     `orm:"province_id"` // 省
	CityId         int     `orm:"city_id"`     // 城市id
	AreaId         int     `orm:"area_id"`     // 区id
	Cuid           int     `orm:"cuid"`
	Avatar         string  `orm:"avatar"`           // tpl头像
	Nickname       string  `orm:"nickname"`         // tpl昵称
	AreaName       string  `orm:"area_name"`        // 对应区域名称
	IsVip          int     `orm:"is_vip"`           // 是否vip
	IsAdmin        int     `orm:"is_admin"`         // 是否后台发布
	Describe       string  `orm:"describe"`         // 描述内容
	Content        string  `orm:"content"`          // 富文本
	Pics           string  `orm:"pics"`             // 图片
	ServiceId      int     `orm:"service_id"`       // 服务id
	UsersServiceId int     `orm:"users_service_id"` // 用户服务id
	OrderId        int     `orm:"order_id"`         // 订单id
	ShareNum       int     `orm:"share_num"`        // 分享数
	ShowNum        int     `orm:"show_num"`         // 显示数
	CommentNum     int     `orm:"comment_num"`      // 评论数
	IsTopic        int     `orm:"is_topic"`         // 是否置顶
	IsNotice       int     `orm:"is_notice"`        // 是否公告
	Flag           int     `orm:"flag"`             // 删除标识
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	IsLbs          int     `orm:"is_lbs"` // 是否显示地址
	Longitude      float64 `orm:"longitude"`
	Latitude       float64 `orm:"latitude"`
	VideoUrl       string  `orm:"video_url"`
	IsAuth         int     `orm:"is_auth"`    // 是否认证通过
	ZanNum         int     `orm:"zan_num"`    // 赞数量
	PartnerId      int     `orm:"partner_id"` // 合伙人id
}

type Admin_config struct {
	Id         int    `orm:"id"`
	Name       string `orm:"name"`        // 名称
	Title      string `orm:"title"`       // 标题
	Group      string `orm:"group"`       // 配置分组
	Type       string `orm:"type"`        // 类型
	Value      string `orm:"value"`       // 配置值
	Options    string `orm:"options"`     // 配置项
	Tips       string `orm:"tips"`        // 配置提示
	AjaxUrl    string `orm:"ajax_url"`    // 联动下拉框ajax地址
	NextItems  string `orm:"next_items"`  // 联动下拉框的下级下拉框名，多个以逗号隔开
	Param      string `orm:"param"`       // 联动下拉框请求参数名
	Format     string `orm:"format"`      // 格式，用于格式文本
	Table      string `orm:"table"`       // 表名，只用于快速联动类型
	Level      int    `orm:"level"`       // 联动级别，只用于快速联动类型
	Key        string `orm:"key"`         // 键字段，只用于快速联动类型
	Option     string `orm:"option"`      // 值字段，只用于快速联动类型
	Pid        string `orm:"pid"`         // 父级id字段，只用于快速联动类型
	Ak         string `orm:"ak"`          // 百度地图appkey
	CreateTime int    `orm:"create_time"` // 创建时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	Sort       int    `orm:"sort"`        // 排序
	Status     int    `orm:"status"`      // 状态：0禁用，1启用
}

type Marketing_card struct {
	Id           int    `orm:"id"`
	PlatformKey  string `orm:"platform_key"`
	EndTime      int    `orm:"end_time"`    // 到期时间 时间戳
	CardLogo     string `orm:"card_logo"`   // 卡片的logo
	CardPic      string `orm:"card_pic"`    // 卡片图片
	Num          int    `orm:"num"`         // 服务数量
	UseNum       int    `orm:"use_num"`     // 服务使用数量
	AreaType     int    `orm:"area_type"`   // 区域类型 0 全国 1省 2 市 3区 4小区
	ProvinceId   int    `orm:"province_id"` // 省id
	CityId       int    `orm:"city_id"`     // 市id
	AreaId       int    `orm:"area_id"`     // 区id
	HousingId    int    `orm:"housing_id"`  // 小区id
	Flag         int    `orm:"flag"`        // 删除标识
	CreatedAt    string `orm:"created_at"`
	UpdatedAt    string `orm:"updated_at"`
	CardTplId    int    `orm:"card_tpl_id"`   // 模板id
	Cuid         int    `orm:"cuid"`          // 用户id
	IsUse        int    `orm:"is_use"`        // 是否使用
	CardNo       string `orm:"card_no"`       // 卡片号码
	CardPassword string `orm:"card_password"` // 卡片密码 预留
	Status       int    `orm:"status"`        // 0未使用 1使用 2用完 -1删除
	OrderNo      string `orm:"order_no"`      // 订单编号
	OrderId      int    `orm:"order_id"`      // 订单id
}

type Sm_quotation_tpl_link struct {
	Id        int    `orm:"id"`
	TplId     int    `orm:"tpl_id"`    // 模板id
	CatId     int    `orm:"cat_id"`    // 模板分类id
	CatName   string `orm:"cat_name"`  // 模板分类名称
	Describe  string `orm:"describe"`  // 备注
	TotalNum  int    `orm:"total_num"` // 项目数
	TableJson string `orm:"table_json"`
	LinkNo    string `orm:"link_no"` // 链接编号
	Flag      int    `orm:"flag"`    // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Ucenter_vip struct {
	Id        int     `orm:"id"`
	Title     string  `orm:"title"`      // 标题
	VipDay    int     `orm:"vip_day"`    // 时长
	CostPrice float64 `orm:"cost_price"` // 原价
	Price     float64 `orm:"price"`      // 现价 支付价格
	Describe  string  `orm:"describe"`   // 描述
	Sort      int     `orm:"sort"`       // 排序
	Flag      int     `orm:"flag"`       // 删除标识
	CreatedAt string  `orm:"created_at"`
	UpdatedAt string  `orm:"updated_at"`
}

type Address_housing struct {
	Id         int     `orm:"id"`
	Name       string  `orm:"name"`        // 小区名称
	ProvinceId int     `orm:"province_id"` // 省id
	CityId     int     `orm:"city_id"`     // 市id
	AreaId     int     `orm:"area_id"`     // 区id
	AddressId  int     `orm:"address_id"`  // 最后一级id
	Describe   string  `orm:"describe"`    // 描述
	IsShow     int     `orm:"is_show"`     // 是否显示
	Flag       int     `orm:"flag"`        // 删除标识
	Longitude  float64 `orm:"longitude"`   // 经度
	Latitude   float64 `orm:"latitude"`    // 维度
	CreatedAt  string  `orm:"created_at"`
	UpdatedAt  string  `orm:"updated_at"`
	Logo       string  `orm:"logo"`
	Points     string  `orm:"points"` // 区域坐标系
	Scale      int     `orm:"scale"`  // 分润
}

type Admin_module struct {
	Id           int    `orm:"id"`
	Name         string `orm:"name"`          // 模块名称（标识）
	Title        string `orm:"title"`         // 模块标题
	Icon         string `orm:"icon"`          // 图标
	Description  string `orm:"description"`   // 描述
	Author       string `orm:"author"`        // 作者
	AuthorUrl    string `orm:"author_url"`    // 作者主页
	Config       string `orm:"config"`        // 配置信息
	Access       string `orm:"access"`        // 授权配置
	Version      string `orm:"version"`       // 版本号
	Identifier   string `orm:"identifier"`    // 模块唯一标识符
	SystemModule int    `orm:"system_module"` // 是否为系统模块
	CreateTime   int    `orm:"create_time"`   // 创建时间
	UpdateTime   int    `orm:"update_time"`   // 更新时间
	Sort         int    `orm:"sort"`          // 排序
	Status       int    `orm:"status"`        // 状态
}

type Sm_users_service struct {
	Id               int     `orm:"id"`
	ServiceId        int     `orm:"service_id"`      // 服务id
	ServiceSkuId     int     `orm:"service_sku_id"`  // sku id
	ServiceAreaId    int     `orm:"service_area_id"` // 区域id
	ApptId           int     `orm:"appt_id"`
	CartId           int     `orm:"cart_id"` // 购物车对应id
	ServiceLogo      string  `orm:"service_logo"`
	ServiceTitle     string  `orm:"service_title"` // 标题
	ApptTitle        string  `orm:"appt_title"`
	CostPrice        float64 `orm:"cost_price"`       // 原价单价
	UnitPrice        float64 `orm:"unit_price"`       // 单价
	TotalCostPrice   float64 `orm:"total_cost_price"` // 全部原价
	TotalPrice       float64 `orm:"total_price"`      // 全部价格
	OrderId          int     `orm:"order_id"`
	Cuid             int     `orm:"cuid"`
	Status           int     `orm:"status"`    // 0 未付款 1付款(拼团属于还没有拼成)  5已确认（拼团拼成）  8已接单 10已服务 待处理
	WorkCuid         int     `orm:"work_cuid"` // 分配工人uid
	Flag             int     `orm:"flag"`      // 删除标识
	NotAppt          int     `orm:"not_appt"`  // 是否暂时不选择服务区间 0否 1是
	CreatedAt        string  `orm:"created_at"`
	UpdatedAt        string  `orm:"updated_at"`
	Price            float64 `orm:"price"`       // 实际付款价格
	Num              int     `orm:"num"`         // 数量
	GjPrice          float64 `orm:"gj_price"`    // 工匠金额
	IsGjPrice        int     `orm:"is_gj_price"` // 是否固定工匠金额
	SkuName          string  `orm:"sku_name"`    // sku 名称
	AddressId        int     `orm:"address_id"`
	ServiceNo        string  `orm:"service_no"`
	IsComment        int     `orm:"is_comment"`      // 是否评论
	IsWorkComment    int     `orm:"is_work_comment"` // 是否商家评论
	Qianzi           string  `orm:"qianzi"`
	IsGp             int     `orm:"is_gp"`               // 是否团购
	UsersServiceGpId int     `orm:"users_service_gp_id"` // 用户团购id
	ServiceGpId      int     `orm:"service_gp_id"`       // 团购id
	IsSeckill        int     `orm:"is_seckill"`          // 是否秒杀
	ServiceSeckillId int     `orm:"service_seckill_id"`
	Describe         string  `orm:"describe"`
	Pics             string  `orm:"pics"`
	IsNewWork        int     `orm:"is_new_work"` // 多工人订单-后结算
	IsJiesuan        int     `orm:"is_jiesuan"`
	TotalFrPrice     float64 `orm:"total_fr_price"`   // 总分润出去的金额
	TotalWorkPrice   float64 `orm:"total_work_price"` // 总工人金额
	AdminDesc        string  `orm:"admin_desc"`
	PartnerId        int     `orm:"partner_id"`    // 合伙人id
	UseShareLv1      int     `orm:"use_share_lv1"` // 强制一级分享
	UseShareLv2      int     `orm:"use_share_lv2"` // 强制二级分享
}

type Marketing_score_rule struct {
	Id          int    `orm:"id"`
	Name        string `orm:"name"`         // 积分模板名称
	ClientTitle string `orm:"client_title"` // 获取积分标题
	PlatformKey string `orm:"platform_key"` // 平台key
	Score       int    `orm:"score"`        // 获取积分
	Status      int    `orm:"status"`       // 是否开启
	IsOne       int    `orm:"is_one"`       // 是否只能获取一次
	Flag        int    `orm:"flag"`         // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	ScoreKey    string `orm:"score_key"` // 唯一标识key
}

type Sm_users_service_work struct {
	Id             int     `orm:"id"`
	Cuid           int     `orm:"cuid"`
	UsersServiceId int     `orm:"users_service_id"` // 用户服务表id
	Price          float64 `orm:"price"`            // 分得金额。默认等分。
	IsEnd          int     `orm:"is_end"`           // 是否结束，一般跟订单结算
	ServiceId      int     `orm:"service_id"`       // 服务id
	OrderId        int     `orm:"order_id"`         // 订单id
	Flag           int     `orm:"flag"`             // 删除标识
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	FrPrice        float64 `orm:"fr_price"` // 最终从这个用户分润出去的金额
}

type Ucenter_accounts struct {
	Id          int     `orm:"id"`
	Cuid        int     `orm:"cuid"`         // ucenter uid
	PlatformKey string  `orm:"platform_key"` // 牵扯的平台
	Type        int     `orm:"type"`         // 0直接收益 1分润收益 2师徒收益 10现金红包收益 50充值收益 100提现 101购买商品
	Level       int     `orm:"level"`        // 收益等级 比如 0一级分享收益 1二级分享收益
	Content     string  `orm:"content"`      // 详细内容
	Describe    string  `orm:"describe"`     // 描述  主要是显示这里
	ProjectId   int     `orm:"project_id"`   // 项目id
	OrderId     int     `orm:"order_id"`     // 订单id
	OrderNo     string  `orm:"order_no"`     // 订单编号
	Price       float64 `orm:"price"`        // 金额
	IsDz        int     `orm:"is_dz"`        // 是否到账 1是
	SourceCuid  int     `orm:"source_cuid"`  // 来源用户。比如是谁分享产生的给你费用
	ProjectName string  `orm:"project_name"` // 项目名称
	Title       string  `orm:"title"`        // 标题
	Flag        int     `orm:"flag"`         // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	AccountNo   string  `orm:"account_no"`   // 订单号
	IsOld       int     `orm:"is_old"`       // 是否是老订单
	ProjectType int     `orm:"project_type"` // 0上门服务 1招工
}

type Ucenter_address struct {
	Id             int     `orm:"id"`
	AdrName        string  `orm:"adr_name"`      // 选择地址名称
	AdrLatitude    float64 `orm:"adr_latitude"`  // 维度
	AdrLongitude   float64 `orm:"adr_longitude"` // 经度
	Address        string  `orm:"address"`       // 详细门牌号
	Cuid           int     `orm:"cuid"`
	IsDefault      int     `orm:"is_default"` // 是否默认
	Name           string  `orm:"name"`       // 姓名
	Mobile         string  `orm:"mobile"`     // 手机号
	Flag           int     `orm:"flag"`       // 删除标识
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	AreaLevel      int     `orm:"area_level"`       // 0省1市 2区 3小区
	AddreaaCheckId int     `orm:"addreaa_check_id"` // 对应的id city 或者housingid
	PlatformKey    string  `orm:"platform_key"`     // 平台key
}

type Face_check_log struct {
	Id        int    `orm:"id"`
	Mac       string `orm:"mac"`       // 设备mac
	DeviceId  int    `orm:"device_id"` // 设备id
	CloudUid  string `orm:"cloud_uid"`
	Photo     string `orm:"photo"`      // 图像
	CloudTime int    `orm:"cloud_time"` // 云端时间
	PhotoHash string `orm:"photo_hash"` // hash
	FaceUid   int    `orm:"face_uid"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Hksmart_database struct {
	Id          int     `orm:"id"`
	AreaId      int     `orm:"area_id"`     // 区域id 省市区最终点
	AreaLevel   int     `orm:"area_level"`  // 区域等级 国 省 市 区
	ProvinceId  int     `orm:"province_id"` // 省id
	CityId      int     `orm:"city_id"`     // 市id
	Name        string  `orm:"name"`        // 库名称
	Year        string  `orm:"year"`        // 年份
	TagNo       string  `orm:"tag_no"`      // 库编号
	Flag        int     `orm:"flag"`        // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	SId         int     `orm:"s_id"` // 专业id
	Desc        string  `orm:"desc"`
	MonthPrice  float64 `orm:"month_price"`  // 月价
	BuyoutPrice float64 `orm:"buyout_price"` // 买断价
}

type Manager_users struct {
	Id            int     `orm:"id"`
	Username      string  `orm:"username"`        // 用户名
	Nickname      string  `orm:"nickname"`        // 昵称
	Password      string  `orm:"password"`        // 密码
	Email         string  `orm:"email"`           // 邮箱地址
	Mobile        string  `orm:"mobile"`          // 手机号码
	Avatar        string  `orm:"avatar"`          // 头像
	Money         float64 `orm:"money"`           // 余额
	Score         int     `orm:"score"`           // 积分
	SignupIp      int     `orm:"signup_ip"`       // 注册ip
	LastLoginTime int     `orm:"last_login_time"` // 最后一次登录时间
	LastLoginIp   int     `orm:"last_login_ip"`   // 登录ip
	Sort          int     `orm:"sort"`            // 排序
	Status        int     `orm:"status"`          // 状态：0禁用，1启用
	InfoBind      int     `orm:"info_bind"`       // 是否绑定用户昵称和头像
	Gender        string  `orm:"gender"`          // 性别
	UserKey       string  `orm:"user_key"`        // 用户唯一编码
	SmallOpenid   string  `orm:"small_openid"`
	RegFlag       int     `orm:"reg_flag"`
	RegTime       int     `orm:"reg_time"`
	CreatedAt     string  `orm:"created_at"`
	UpdatedAt     string  `orm:"updated_at"`
	DeletedAt     string  `orm:"deleted_at"`
	Flag          int     `orm:"flag"`      // -1删除
	AuthType      int     `orm:"auth_type"` // 0无认证 1企业认证 2施工队 3施工人员
	IsAuth        int     `orm:"is_auth"`   // 是否认证
	Lv            int     `orm:"lv"`
	Pingfen       float64 `orm:"pingfen"` // 评分
	ManagerName   string  `orm:"manager_name"`
	Pid           int     `orm:"pid"`
	PidTwo        int     `orm:"pid_two"`
	Muid          int     `orm:"muid"`
	Sgzl          float64 `orm:"sgzl"`    // 施工质量
	Sgjd          float64 `orm:"sgjd"`    // 施工进度
	Fwtd          float64 `orm:"fwtd"`    // 服务态度
	Gdgl          float64 `orm:"gdgl"`    // 工地管理
	Sc            string  `orm:"sc"`      // 擅长标签
	Gz            string  `orm:"gz"`      // 工种
	SgCity        string  `orm:"sg_city"` // 施工区域
	Shifu         int     `orm:"shifu"`
	ShareLv1      int     `orm:"share_lv1"`
	ShareLv2      int     `orm:"share_lv2"`
	Shouyi        float64 `orm:"shouyi"`
	IsTopic       int     `orm:"is_topic"`
	StShouyi      float64 `orm:"st_shouyi"`
	ShareShouyi   float64 `orm:"share_shouyi"`
	Latitude      float64 `orm:"latitude"`  // 纬度
	Longitude     float64 `orm:"longitude"` // 经度
}

type Marketing_card_tpl struct {
	Id          int    `orm:"id"`
	PlatformKey string `orm:"platform_key"`
	EndTime     int    `orm:"end_time"`    // 到期时间-模板是天数
	CardLogo    string `orm:"card_logo"`   // 卡片的logo
	CardPic     string `orm:"card_pic"`    // 卡片图片
	Num         int    `orm:"num"`         // 预留最大生成数量
	AreaType    int    `orm:"area_type"`   // 区域类型 0 全国 1省 2 市 3区 4小区
	ProvinceId  int    `orm:"province_id"` // 省id
	CityId      int    `orm:"city_id"`     // 市id
	AreaId      int    `orm:"area_id"`     // 区id
	HousingId   int    `orm:"housing_id"`  // 小区id
	Flag        int    `orm:"flag"`        // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	Describe    string `orm:"describe"` // 描述
	Title       string `orm:"title"`    // 卡券标题
}

type Marketing_seckill_packet_link struct {
	Id              int     `orm:"id"`
	SeckillPacketId int     `orm:"seckill_packet_id"` // 抢红包id
	Price           float64 `orm:"price"`             // 红包金额
	Num             int     `orm:"num"`               // 红包数量
	UseNum          int     `orm:"use_num"`           // 使用数量
	IsEnd           int     `orm:"is_end"`            // 是否抢完
	Flag            int     `orm:"flag"`              // 删除标识
	CreatedAt       string  `orm:"created_at"`
	UpdatedAt       string  `orm:"updated_at"`
}

type Sm_service_seckill_sku struct {
	Id               int    `orm:"id"`
	ServiceId        int    `orm:"service_id"`
	ServiceSeckillId int    `orm:"service_seckill_id"`
	ServiceAreaId    int    `orm:"service_area_id"` // 区域id
	ServiceSkuId     int    `orm:"service_sku_id"`  // sku id
	Flag             int    `orm:"flag"`            // 删除标识
	CreatedAt        string `orm:"created_at"`
	UpdatedAt        string `orm:"updated_at"`
}

type Face_users struct {
	Id        int    `orm:"id"`
	Name      string `orm:"name"`      // 姓名
	Mobile    string `orm:"mobile"`    // 手机号
	Feature   string `orm:"feature"`   // 特征
	CloudUid  string `orm:"cloud_uid"` // 云端用于匹配的id
	CloudId   int    `orm:"cloud_id"`  // 云端id
	Flag      int    `orm:"flag"`      // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	AvatarUrl string `orm:"avatar_url"` // 头像url
}

type Sm_quotation struct {
	Id           int     `orm:"id"`
	TplId        int     `orm:"tpl_id"` // 模板id
	Name         string  `orm:"name"`   // 报价单名称
	Describe     string  `orm:"describe"`
	QuotationNo  string  `orm:"quotation_no"` // 报价单号
	TotalPrice   float64 `orm:"total_price"`  // 总价
	TotalNum     int     `orm:"total_num"`    // 项目总数
	Status       int     `orm:"status"`       // 0正常 1进行中 2签订
	UAddress     string  `orm:"u_address"`    // 用户地址
	UMobile      string  `orm:"u_mobile"`     // 用户手机号
	UName        string  `orm:"u_name"`       // 用户姓名
	PartnerId    int     `orm:"partner_id"`   // 合伙人id
	EndPrice     float64 `orm:"end_price"`    // 最终价格
	Flag         int     `orm:"flag"`         // 删除标识
	CreatedAt    string  `orm:"created_at"`
	UpdatedAt    string  `orm:"updated_at"`
	ManagerBl    int     `orm:"manager_bl"` // 管理费比例
	ManagerPrice float64 `orm:"manager_price"`
	FixedPrice   float64 `orm:"fixed_price"` // 一口价
	TaxBl        int     `orm:"tax_bl"`      // 税率
}

type Sm_quotation_unit struct {
	Id        int    `orm:"id"`
	UnitName  string `orm:"unit_name"`
	Describe  string `orm:"describe"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type Sm_users_address struct {
	Id             int     `orm:"id"`
	AdrName        string  `orm:"adr_name"`      // 选择地址名称
	AdrLatitude    float64 `orm:"adr_latitude"`  // 维度
	AdrLongitude   float64 `orm:"adr_longitude"` // 经度
	Address        string  `orm:"address"`       // 详细门牌号
	Cuid           int     `orm:"cuid"`
	IsDefault      int     `orm:"is_default"` // 是否默认
	Name           string  `orm:"name"`       // 姓名
	Mobile         string  `orm:"mobile"`     // 手机号
	Flag           int     `orm:"flag"`       // 删除标识
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	AreaLevel      int     `orm:"area_level"`       // 0省1市 2区 3小区
	AddreaaCheckId int     `orm:"addreaa_check_id"` // 对应的id city 或者housingid
}

type Xinxi_news struct {
	Id        int    `orm:"id"`
	Pic       string `orm:"pic"`     // 图片
	Desc      string `orm:"desc"`    // 描述
	Source    string `orm:"source"`  // 来源
	Content   string `orm:"content"` // 详情
	Flag      int    `orm:"flag"`    // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	CatId     int    `orm:"cat_id"`
	Title     string `orm:"title"` // 标题
	IsShow    int    `orm:"is_show"`
}

type Admin_hook struct {
	Id          int    `orm:"id"`
	Name        string `orm:"name"`        // 钩子名称
	Plugin      string `orm:"plugin"`      // 钩子来自哪个插件
	Description string `orm:"description"` // 钩子描述
	System      int    `orm:"system"`      // 是否为系统钩子
	CreateTime  int    `orm:"create_time"` // 创建时间
	UpdateTime  int    `orm:"update_time"` // 更新时间
	Status      int    `orm:"status"`      // 状态
}

type Marketing_packet_tpl struct {
	Id          int     `orm:"id"`
	Price       float64 `orm:"price"`
	Title       string  `orm:"title"`
	Describe    string  `orm:"describe"` // 描述
	Type        string  `orm:"type"`     // 来源渠道
	Flag        int     `orm:"flag"`     // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	PlatformKey string  `orm:"platform_key"`
}

type Partner_area struct {
	Id         int    `orm:"id"`
	PartnerId  int    `orm:"partner_id"`
	AreaType   int    `orm:"area_type"`   // 0 全国 1全省 2市 3区 4小区
	ProvinceId int    `orm:"province_id"` // 省id
	CityId     int    `orm:"city_id"`     // 市id
	AreaId     int    `orm:"area_id"`     // 区id
	HousingId  int    `orm:"housing_id"`  // 小区id
	Flag       int    `orm:"flag"`        // 删除标识
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
	IsDefault  int    `orm:"is_default"` // 是否默认
}
