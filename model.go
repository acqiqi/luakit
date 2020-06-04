package main

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
	IsAuth         int     `orm:"is_auth"`         // 是否实名认证 0 否 1审核 2通过 -1拒绝
	IdcardTop      string  `orm:"idcard_top"`      // 身份证正面
	IdcardBom      string  `orm:"idcard_bom"`      // 身份证背面
	IdcardId       string  `orm:"idcard_id"`       // 身份证号
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
	BindUserinfo   int     `orm:"bind_userinfo"`    // 是否绑定用户信息
	IsVip          int     `orm:"is_vip"`           // 是否是vip
	VipEndTime     string  `orm:"vip_end_time"`     // vip到期时间
	IsPayPassword  int     `orm:"is_pay_password"`  // 是否填写支付密码
	PayPassword    string  `orm:"pay_password"`     // 支付密码
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

type Ucenter_orders struct {
	Id            int     `orm:"id"`
	OrderNo       string  `orm:"order_no"`     // 订单编号
	PlatformKey   string  `orm:"platform_key"` // 平台key
	Cuid          int     `orm:"cuid"`         // ucenter id
	CouponKey     string  `orm:"coupon_key"`   // 优惠券key
	CouponId      int     `orm:"coupon_id"`    // 优惠券id
	CouponPrice   float64 `orm:"coupon_price"` // 优惠券抵扣金额
	CostPrice     float64 `orm:"cost_price"`   // 原价
	UnitPrice     float64 `orm:"unit_price"`   // 单价
	Price         float64 `orm:"price"`        // 现价 支付总价
	GoodsNum      float64 `orm:"goods_num"`    // 商品总数量
	PayType       int     `orm:"pay_type"`     // 支付类型 0线上 1线下
	PayPlatform   string  `orm:"pay_platform"` // 支付平台 wechat alipay ...
	Status        int     `orm:"status"`       // 状态 0下单 2审核通过 7已接单 8已发货 9已结算或者已收货
	IsPay         int     `orm:"is_pay"`       // 是否支付 0否 1线下提交 2已支付
	IsUComment    int     `orm:"is_u_comment"` // 是否用户评论 0否 1是
	IsMComment    int     `orm:"is_m_comment"` // 是否商家评论 0否1是
	Flag          int     `orm:"flag"`         // 删除标识
	CreatedAt     string  `orm:"created_at"`
	UpdatedAt     string  `orm:"updated_at"`
	PayTime       int     `orm:"pay_time"`       // 支付时间
	ServicePrice  float64 `orm:"service_price"`  // 服务费
	SharePrice    float64 `orm:"share_price"`    // 分享出去多少钱
	PlatformPrice float64 `orm:"platform_price"` // 平台受益
	ShareLv1      float64 `orm:"share_lv1"`      // 一级分享
	ShareLv2      float64 `orm:"share_lv2"`      // 二级分享
	OptLv1        float64 `orm:"opt_lv1"`        // 一级其他分享
	OptLv2        float64 `orm:"opt_lv2"`        // 二级其他分享
	ShareLv1Cuid  int     `orm:"share_lv1_cuid"` // 一级分享用户id
	ShareLv2Cuid  int     `orm:"share_lv2_cuid"`
	OptLv1Cuid    int     `orm:"opt_lv1_cuid"`
	OptLv2Cuid    int     `orm:"opt_lv2_cuid"`
	OrderType     int     `orm:"order_type"` // 订单类型0常规购买订单  1团购 10VIP
	PayNo         string  `orm:"pay_no"`     // 线上付款订单
	Describe      string  `orm:"describe"`   // 描述
	ProjectId     int     `orm:"project_id"`
}

type Admin_access struct {
	Module string `orm:"module"` // 模型名称
	Group  string `orm:"group"`  // 权限分组标识
	Uid    int    `orm:"uid"`    // 用户id
	Nid    string `orm:"nid"`    // 授权节点id
	Tag    string `orm:"tag"`    // 分组标签
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

type Admin_icon_list struct {
	Id     int    `orm:"id"`
	IconId int    `orm:"icon_id"` // 所属图标id
	Title  string `orm:"title"`   // 图标标题
	Class  string `orm:"class"`   // 图标类名
	Code   string `orm:"code"`    // 图标关键词
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

type Sm_service_appt_tpl struct {
	Id        int    `orm:"id"`
	TplName   string `orm:"tpl_name"` // 模板名称
	Describe  string `orm:"describe"` // 描述
	Flag      int    `orm:"flag"`     // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
}

type System_log struct {
	Id        int    `orm:"id"`
	Content   string `orm:"content"`
	Flag      int    `orm:"flag"` // 删除标识
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

type Sm_service struct {
	Id           int     `orm:"id"`
	Title        string  `orm:"title"`          // 标题
	Logo         string  `orm:"logo"`           // 列表logo
	Icon         string  `orm:"icon"`           // 首页icon
	Banner       string  `orm:"banner"`         // banner列表 json
	CatsId       int     `orm:"cats_id"`        // 分类id
	Describe     string  `orm:"describe"`       // 描述
	Content      string  `orm:"content"`        // 富文本内容
	IsVideo      int     `orm:"is_video"`       // 是否显示视频banner
	VideoUrl     string  `orm:"video_url"`      // 视频url
	ShowNum      int     `orm:"show_num"`       // 显示数量
	PayNum       int     `orm:"pay_num"`        // 销售数量
	CollectNum   int     `orm:"collect_num"`    // 收藏数量
	ShareNum     int     `orm:"share_num"`      // 分享数量
	MinPrice     float64 `orm:"min_price"`      // 最小售价（用于限定最低单品购买以及显示）
	MinCostPrice float64 `orm:"min_cost_price"` // 最小销售原价
	ProjectType  int     `orm:"project_type"`   // 0正常 1团购
	IsShow       int     `orm:"is_show"`        // 是否显示
	Status       int     `orm:"status"`         // 0停售 1正常销售  10预售
	Flag         int     `orm:"flag"`           // 删除标识
	CreatedAt    string  `orm:"created_at"`
	UpdatedAt    string  `orm:"updated_at"`
	IsDelete     int     `orm:"is_delete"`   // 是否远程删除
	IsTopic      int     `orm:"is_topic"`    // 是否推荐
	IsNew        int     `orm:"is_new"`      // 是否new
	Sort         int     `orm:"sort"`        // 排序
	AreaType     int     `orm:"area_type"`   // 限定区域范围等级 0全国 1省 2市 3区 4小区
	ProvinceId   int     `orm:"province_id"` // 省id
	CityId       int     `orm:"city_id"`     // 市id
	AreaId       int     `orm:"area_id"`     // 区id
	HousingId    int     `orm:"housing_id"`  // 小区id
}

type Sm_service_gp struct {
	Id        int     `orm:"id"`
	ServiceId int     `orm:"service_id"`  // 服务id
	Title     string  `orm:"title"`       // 团购标题
	Content   string  `orm:"content"`     // 富文本
	AllNum    int     `orm:"all_num"`     // 满足多少成大团
	ShareNum  int     `orm:"share_num"`   // 满足多少成小团
	Type      int     `orm:"type"`        // 0满大团 1满小团 2满大小团 10不管团多少都tm成
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

type Log struct {
	Id        int    `orm:"id"`
	Content   string `orm:"content"`
	Flag      int    `orm:"flag"` // 删除标识
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
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
	OrderNo     int     `orm:"order_no"`     // 订单编号
	Price       float64 `orm:"price"`        // 金额
	IsDz        int     `orm:"is_dz"`        // 是否到账 1是
	SourceCuid  int     `orm:"source_cuid"`  // 来源用户。比如是谁分享产生的给你费用
	ProjectName string  `orm:"project_name"` // 项目名称
	Title       string  `orm:"title"`        // 标题
	Flag        int     `orm:"flag"`         // 删除标识
	CreatedAt   string  `orm:"created_at"`
	UpdatedAt   string  `orm:"updated_at"`
	AccountNo   string  `orm:"account_no"` // 订单号
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

type Admin_hook_plugin struct {
	Id         int    `orm:"id"`
	Hook       string `orm:"hook"`        // 钩子id
	Plugin     string `orm:"plugin"`      // 插件标识
	CreateTime int    `orm:"create_time"` // 添加时间
	UpdateTime int    `orm:"update_time"` // 更新时间
	Sort       int    `orm:"sort"`        // 排序
	Status     int    `orm:"status"`      // 状态
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

type Partner_area struct {
	Id         int    `orm:"id"`
	PartenrId  int    `orm:"partenr_id"`
	AreaType   int    `orm:"area_type"`   // 0 全国 1全省 2市 3区 4小区
	ProvinceId int    `orm:"province_id"` // 省id
	CityId     int    `orm:"city_id"`     // 市id
	AreaId     int    `orm:"area_id"`     // 区id
	HousingId  int    `orm:"housing_id"`  // 小区id
	Flag       int    `orm:"flag"`        // 删除标识
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
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