package main

type Face_device struct {
	Id            int     `:"id" json:"id"`
	Name          string  `:"name" json:"name"`                     // 设备名字
	Describe      string  `:"describe" json:"describe"`             // 设备描述
	Address       string  `:"address" json:"address"`               // 详细地址
	Mac           string  `:"mac" json:"mac"`                       // 设备mac
	Sign          string  `:"sign" json:"sign"`                     // 设备sign
	Cmd           int     `:"cmd" json:"cmd"`                       // 荔枝默认14
	RelayOpenVol  int     `:"relay_open_vol" json:"relay_open_vol"` // 继电器极性
	ChkFacePose   int     `:"chk_face_pose" json:"chk_face_pose"`   // 正脸判断使能
	FaceThreshold int     `:"face_threshold" json:"face_threshold"` // 不传则表示不修改。识别阈值，范围1-100。建议80
	LedBright     float64 `:"led_bright" json:"led_bright"`         // 不传则表示不修改。补光灯亮度调节，如0.1， 0.01
	RelayOpenS    int     `:"relay_open_s" json:"relay_open_s"`     // 不传则表示不修改。继电器开光延时，可选1，2，3等
	Flag          int     `:"flag" json:"flag"`                     // 删除标识
	CreatedAt     string  `:"created_at" json:"created_at"`
	UpdatedAt     string  `:"updated_at" json:"updated_at"`
	IsDelete      int     `:"is_delete" json:"is_delete"`       // 是否远程删除
	NetworkType   int     `:"network_type" json:"network_type"` // 0 wifi 1网口 2 4G
	NetworkData   string  `:"network_data" json:"network_data"` // 联网配置参数存储
}

type Sm_users_cart struct {
	Id            int     `:"id" json:"id"`
	ServiceId     int     `:"service_id" json:"service_id"`         // 服务id
	ServiceSkuId  int     `:"service_sku_id" json:"service_sku_id"` // sku id
	ApptId        int     `:"appt_id" json:"appt_id"`               // 对应服务时间id
	Num           int     `:"num" json:"num"`                       // 购物车数量
	Cuid          int     `:"cuid" json:"cuid"`                     // 用户id
	UnitPrice     float64 `:"unit_price" json:"unit_price"`         // 售价
	CostPrice     float64 `:"cost_price" json:"cost_price"`         // 原价
	ServiceLogo   string  `:"service_logo" json:"service_logo"`
	ServiceTitle  string  `:"service_title" json:"service_title"` // 标题
	ApptTitle     string  `:"appt_title" json:"appt_title"`
	Flag          int     `:"flag" json:"flag"` // 删除标识
	CreatedAt     string  `:"created_at" json:"created_at"`
	UpdatedAt     string  `:"updated_at" json:"updated_at"`
	AddressId     string  `:"address_id" json:"address_id"`           // 用户下单的地址id
	ServiceAreaId int     `:"service_area_id" json:"service_area_id"` // 区域id
	NotAppt       int     `:"not_appt" json:"not_appt"`               // 是否暂时不选择服务区间 0否 1是
	GjPrice       float64 `:"gj_price" json:"gj_price"`               // 工匠金额
	IsGjPrice     int     `:"is_gj_price" json:"is_gj_price"`         // 是否固定工匠金额
}

type Ucenter_feedback struct {
	Id          int    `:"id" json:"id"`
	Uname       string `:"uname" json:"uname"`
	Mobile      string `:"mobile" json:"mobile"`
	Content     string `:"content" json:"content"`
	Cuid        int    `:"cuid" json:"cuid"`
	Flag        int    `:"flag" json:"flag"` // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	PlatformKey string `:"platform_key" json:"platform_key"`
	IsBlack     int    `:"is_black" json:"is_black"` // 是否是黑名单提交
}

type Admin_action struct {
	Id         int    `:"id" json:"id"`
	Module     string `:"module" json:"module"`           // 所属模块名
	Name       string `:"name" json:"name"`               // 行为唯一标识
	Title      string `:"title" json:"title"`             // 行为标题
	Remark     string `:"remark" json:"remark"`           // 行为描述
	Rule       string `:"rule" json:"rule"`               // 行为规则
	Log        string `:"log" json:"log"`                 // 日志规则
	Status     int    `:"status" json:"status"`           // 状态
	CreateTime int    `:"create_time" json:"create_time"` // 创建时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
}

type Message_queue struct {
	Id              int    `:"id" json:"id"`
	Uid             int    `:"uid" json:"uid"`
	UserType        int    `:"user_type" json:"user_type"` // 0用户端 1商家端
	UcUid           int    `:"uc_uid" json:"uc_uid"`
	MessageKey      string `:"message_key" json:"message_key"`
	Title           string `:"title" json:"title"`               // 标题
	Desc            string `:"desc" json:"desc"`                 // 描述
	Content         string `:"content" json:"content"`           // 内容
	MessageType     int    `:"message_type" json:"message_type"` // 消息类型
	PathType        string `:"path_type" json:"path_type"`       // 链接类型
	PathId          string `:"path_id" json:"path_id"`           // 链接id
	IsFormId        int    `:"is_form_id" json:"is_form_id"`     // 是否消息模板
	SmallTplId      string `:"small_tpl_id" json:"small_tpl_id"` // 消息模板id
	SmallTplContent string `:"small_tpl_content" json:"small_tpl_content"`
	SmallTplOpenid  string `:"small_tpl_openid" json:"small_tpl_openid"`
	IsSms           int    `:"is_sms" json:"is_sms"` // 是否发送短信
	Mobile          string `:"mobile" json:"mobile"`
	SmsContent      string `:"sms_content" json:"sms_content"` // 短信内容
	IsEmail         int    `:"is_email" json:"is_email"`       // 是否发邮件
	Email           string `:"email" json:"email"`
	EmailTitle      string `:"email_title" json:"email_title"`
	EmailContent    string `:"email_content" json:"email_content"`
	IsSend          int    `:"is_send" json:"is_send"`       // 是否发送
	MsgTplId        int    `:"msg_tpl_id" json:"msg_tpl_id"` // message tpl id
	CreatedAt       string `:"created_at" json:"created_at"`
	UpdatedAt       string `:"updated_at" json:"updated_at"`
	DeletedAt       string `:"deleted_at" json:"deleted_at"`
	Flag            int    `:"flag" json:"flag"`     // -1删除
	IsPop           int    `:"is_pop" json:"is_pop"` // 1已经出列
	PushData        string `:"push_data" json:"push_data"`
	PlatformKey     string `:"platform_key" json:"platform_key"`
	IsUcId          int    `:"is_uc_id" json:"is_uc_id"`
	SmallTplPath    string `:"small_tpl_path" json:"small_tpl_path"`
}

type Sm_service_cats struct {
	Id          int    `:"id" json:"id"`
	CatName     string `:"cat_name" json:"cat_name"`           // 分类名称
	CatLogo     string `:"cat_logo" json:"cat_logo"`           // 分类logo
	CatHomeLogo string `:"cat_home_logo" json:"cat_home_logo"` // 用于首页显示的logo
	CatTags     string `:"cat_tags" json:"cat_tags"`           // 分类标签
	Sort        int    `:"sort" json:"sort"`                   // 排序 倒序
	IsShow      int    `:"is_show" json:"is_show"`
	Flag        int    `:"flag" json:"flag"` // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	Pid         int    `:"pid" json:"pid"`
	Level       int    `:"level" json:"level"` // 层级
}

type Sm_users struct {
	Id              int     `:"id" json:"id"`
	Cuid            int     `:"cuid" json:"cuid"`
	Longitude       float64 `:"longitude" json:"longitude"`
	Latitude        float64 `:"latitude" json:"latitude"`
	IsSetArea       int     `:"is_set_area" json:"is_set_area"`
	AreaName        string  `:"area_name" json:"area_name"`               // 当前区域的名称
	AreaLevel       int     `:"area_level" json:"area_level"`             // 0省1市 2区 3小区
	AddreaaCheckId  int     `:"addreaa_check_id" json:"addreaa_check_id"` // 对应的id city 或者housingid
	Flag            int     `:"flag" json:"flag"`                         // 删除标识
	CreatedAt       string  `:"created_at" json:"created_at"`
	UpdatedAt       string  `:"updated_at" json:"updated_at"`
	IsNewUser       int     `:"is_new_user" json:"is_new_user"`               // 是否是当前端新用户
	IsGjAuth        int     `:"is_gj_auth" json:"is_gj_auth"`                 // 是否工匠认证 0否 1认证中 2已认证 -1驳回
	GjIdcardTop     string  `:"gj_idcard_top" json:"gj_idcard_top"`           // 身份证正面
	GjIdcardBom     string  `:"gj_idcard_bom" json:"gj_idcard_bom"`           // 身份证背面
	GjIdcardId      string  `:"gj_idcard_id" json:"gj_idcard_id"`             // 身份证号码
	GjIdcardEndTime string  `:"gj_idcard_end_time" json:"gj_idcard_end_time"` // 到期时间
	GjIdcardAddress string  `:"gj_idcard_address" json:"gj_idcard_address"`   // 身份证地址
	GjName          string  `:"gj_name" json:"gj_name"`                       // 认证姓名
	GjAvatar        string  `:"gj_avatar" json:"gj_avatar"`                   // 工匠头像
	GjProvinceId    int     `:"gj_province_id" json:"gj_province_id"`         // 省id
	GjCityId        int     `:"gj_city_id" json:"gj_city_id"`                 // 市id
	GjAreaId        int     `:"gj_area_id" json:"gj_area_id"`                 // 区id
	GjAreaLevel     int     `:"gj_area_level" json:"gj_area_level"`           // 区域等级 0省 1市 2区
	GjMobile        string  `:"gj_mobile" json:"gj_mobile"`                   // 工匠认证手机号
	GjDesc          string  `:"gj_desc" json:"gj_desc"`                       // 工匠介绍
	GjSfCuid        int     `:"gj_sf_cuid" json:"gj_sf_cuid"`                 // 师傅id
	IsSf            int     `:"is_sf" json:"is_sf"`                           // 是否师父
	GjErrMsg        string  `:"gj_err_msg" json:"gj_err_msg"`                 // 认证失败文案
}

type Sm_users_service_gp struct {
	Id             int    `:"id" json:"id"`
	ServiceId      int    `:"service_id" json:"service_id"`             // 服务id
	ServiceGpId    int    `:"service_gp_id" json:"service_gp_id"`       // 服务拼团id
	ServiceSkuId   int    `:"service_sku_id" json:"service_sku_id"`     // 服务sku
	OrderId        int    `:"order_id" json:"order_id"`                 // 订单id
	UsersServiceId int    `:"users_service_id" json:"users_service_id"` // 用户服务表id
	Cuid           int    `:"cuid" json:"cuid"`
	Status         int    `:"status" json:"status"` // -1结束 0未支付 1支付 2完成
	Flag           int    `:"flag" json:"flag"`     // 删除标识
	CreatedAt      string `:"created_at" json:"created_at"`
	UpdatedAt      string `:"updated_at" json:"updated_at"`
	Num            int    `:"num" json:"num"`
}

type System_banner struct {
	Id          int    `:"id" json:"id"`
	Title       string `:"title" json:"title"`       // 标题
	Describe    string `:"describe" json:"describe"` // 描述
	Banner      string `:"banner" json:"banner"`
	GotoType    int    `:"goto_type" json:"goto_type"` // 0不转跳 1单页 2path
	Path        string `:"path" json:"path"`           // 路径
	Param       string `:"param" json:"param"`         // 页面参数 或者id
	PlatformKey string `:"platform_key" json:"platform_key"`
	IsShow      int    `:"is_show" json:"is_show"`
	Sort        int    `:"sort" json:"sort"`
	Flag        int    `:"flag" json:"flag"` // 删除标识
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	CreatedAt   string `:"created_at" json:"created_at"`
	BannerKey   string `:"banner_key" json:"banner_key"`
	VideoUrl    string `:"video_url" json:"video_url"`
}

type Log struct {
	Id        int    `:"id" json:"id"`
	Content   string `:"content" json:"content"`
	Flag      int    `:"flag" json:"flag"` // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
}

type Marketing_coupon struct {
	Id            int     `:"id" json:"id"`
	CouponTplId   int     `:"coupon_tpl_id" json:"coupon_tpl_id"`     // 模板id
	CouponQueueId int     `:"coupon_queue_id" json:"coupon_queue_id"` // 批量发送id 如果有
	SourceType    int     `:"source_type" json:"source_type"`         // 0后台发送 1新用户注册 2事件触发
	Cuid          int     `:"cuid" json:"cuid"`
	Logo          string  `:"logo" json:"logo"`
	Title         string  `:"title" json:"title"`               // 优惠券名称
	PlatformKey   string  `:"platform_key" json:"platform_key"` // 平台key
	Price         float64 `:"price" json:"price"`               // 优惠金额/最大优惠金额
	FullPrice     float64 `:"full_price" json:"full_price"`     // 满金额条件
	Type          int     `:"type" json:"type"`                 // 0满减 1全局折扣 2满折
	Zkb           int     `:"zkb" json:"zkb"`                   // 折扣比 0 100
	ProjectId     int     `:"project_id" json:"project_id"`     // 对应项目id （上门端是skuid 或者serviceid）
	ProjectType   int     `:"project_type" json:"project_type"` // 对应项目类型。每个平台不同  （上门端0 sku 1 service）
	EndTime       int     `:"end_time" json:"end_time"`         // 到期时间。被转换的时间戳
	Describe      string  `:"describe" json:"describe"`         // 描述
	IsUse         int     `:"is_use" json:"is_use"`             // 是否使用
	OrderId       int     `:"order_id" json:"order_id"`         // 使用后订单id
	Flag          int     `:"flag" json:"flag"`                 // 删除标识
	UpdatedAt     string  `:"updated_at" json:"updated_at"`
	CreatedAt     string  `:"created_at" json:"created_at"`
	CouponKey     string  `:"coupon_key" json:"coupon_key"` // 优惠券唯一key
}

type Marketing_coupon_tpl struct {
	Id          int     `:"id" json:"id"`
	Logo        string  `:"logo" json:"logo"`
	Title       string  `:"title" json:"title"`               // 优惠券名称
	PlatformKey string  `:"platform_key" json:"platform_key"` // 平台key
	Price       float64 `:"price" json:"price"`               // 优惠金额/最大优惠金额
	FullPrice   float64 `:"full_price" json:"full_price"`     // 满金额条件
	Type        int     `:"type" json:"type"`                 // 0满减 1全局折扣 2满折
	Zkb         int     `:"zkb" json:"zkb"`                   // 折扣比 0 100
	ProjectId   int     `:"project_id" json:"project_id"`     // 对应项目id （上门端是skuid 或者serviceid）
	ProjectType int     `:"project_type" json:"project_type"` // 对应项目类型。每个平台不同  （上门端0 sku 1 service）
	EndTime     int     `:"end_time" json:"end_time"`         // 到期时间 天
	Flag        int     `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string  `:"created_at" json:"created_at"`
	UpdatedAt   string  `:"updated_at" json:"updated_at"`
	Describe    string  `:"describe" json:"describe"` // 描述
}

type Partner_users struct {
	Id        int     `:"id" json:"id"`
	Username  string  `:"username" json:"username"`   // 账号
	Password  string  `:"password" json:"password"`   // 密码
	Mobile    string  `:"mobile" json:"mobile"`       // 手机号
	Nickname  string  `:"nickname" json:"nickname"`   // 昵称
	Email     string  `:"email" json:"email"`         // 邮箱
	Gender    string  `:"gender" json:"gender"`       // 性别
	Status    int     `:"status" json:"status"`       // 状态 0停用 1启用
	RoleType  int     `:"role_type" json:"role_type"` // 0无权限 1小区合伙人 2区县级合伙人 3城市合伙人 4省合伙人
	Flag      int     `:"flag" json:"flag"`           // 删除标识
	CreatedAt string  `:"created_at" json:"created_at"`
	UpdatedAt string  `:"updated_at" json:"updated_at"`
	Describe  string  `:"describe" json:"describe"`
	Money     float64 `:"money" json:"money"` // 余额
}

type Sm_users_comment struct {
	Id             int    `:"id" json:"id"`
	ServiceId      int    `:"service_id" json:"service_id"`             // 服务id
	ApptId         int    `:"appt_id" json:"appt_id"`                   // 服务区间
	ServiceSkuId   int    `:"service_sku_id" json:"service_sku_id"`     // 服务sku
	OrderId        int    `:"order_id" json:"order_id"`                 // 订单id
	UsersServiceId int    `:"users_service_id" json:"users_service_id"` // 用户服务表id
	ServiceAreaId  int    `:"service_area_id" json:"service_area_id"`   // 区域id
	Cuid           int    `:"cuid" json:"cuid"`
	Type           int    `:"type" json:"type"` // 0用户 1商家
	Content        string `:"content" json:"content"`
	Pics           string `:"pics" json:"pics"`
	VideoUrl       string `:"video_url" json:"video_url"`
	Flag           int    `:"flag" json:"flag"` // 删除标识
	CreatedAt      string `:"created_at" json:"created_at"`
	UpdatedAt      string `:"updated_at" json:"updated_at"`
	Avatar         string `:"avatar" json:"avatar"`     // 备份头像
	Nickname       string `:"nickname" json:"nickname"` // 备份昵称
	Star           int    `:"star" json:"star"`         // 星 1非常不满意 2不满意 3一般 4满意 5 非常满意
	Tags           string `:"tags" json:"tags"`
}

type System_config struct {
	Id          int    `:"id" json:"id"`
	K           string `:"k" json:"k"`                       // key
	V           string `:"v" json:"v"`                       // value
	Title       string `:"title" json:"title"`               // 标题
	Describe    string `:"describe" json:"describe"`         // 描述
	PlatformKey string `:"platform_key" json:"platform_key"` // platform_key  如果全局写ALL
	Flag        int    `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
}

type Address_housing struct {
	Id         int     `:"id" json:"id"`
	Name       string  `:"name" json:"name"`               // 小区名称
	ProvinceId int     `:"province_id" json:"province_id"` // 省id
	CityId     int     `:"city_id" json:"city_id"`         // 市id
	AreaId     int     `:"area_id" json:"area_id"`         // 区id
	AddressId  int     `:"address_id" json:"address_id"`   // 最后一级id
	Describe   string  `:"describe" json:"describe"`       // 描述
	IsShow     int     `:"is_show" json:"is_show"`         // 是否显示
	Flag       int     `:"flag" json:"flag"`               // 删除标识
	Longitude  float64 `:"longitude" json:"longitude"`     // 经度
	Latitude   float64 `:"latitude" json:"latitude"`       // 维度
	CreatedAt  string  `:"created_at" json:"created_at"`
	UpdatedAt  string  `:"updated_at" json:"updated_at"`
	Logo       string  `:"logo" json:"logo"`
	Points     string  `:"points" json:"points"` // 区域坐标系
	Scale      int     `:"scale" json:"scale"`   // 分润
}

type Ucenter_bank struct {
	Id          int    `:"id" json:"id"`
	Uname       string `:"uname" json:"uname"`               // 姓名
	BankId      string `:"bank_id" json:"bank_id"`           // 银行卡号
	BankName    string `:"bank_name" json:"bank_name"`       // 银行名称
	BankAddress string `:"bank_address" json:"bank_address"` // 开户行
	Mobile      string `:"mobile" json:"mobile"`             // 预留手机号
	IsDefault   int    `:"is_default" json:"is_default"`     // 是否默认
	Cuid        int    `:"cuid" json:"cuid"`
	Flag        int    `:"flag" json:"flag"` // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
}

type Ucenter_platform struct {
	Id               int    `:"id" json:"id"`
	PlatformName     string `:"platform_name" json:"platform_name"`         // 平台名称
	PlatformUsername string `:"platform_username" json:"platform_username"` // 平台账号
	PlatformPassword string `:"platform_password" json:"platform_password"` // 平台密码
	PlatformAppType  int    `:"platform_app_type" json:"platform_app_type"` // 平台类型 0 api 1 web 5 wechat 6 alipay
	PlatformKey      string `:"platform_key" json:"platform_key"`           // 平台key 完全标识
	Status           int    `:"status" json:"status"`                       // 状态 1运行 0暂停维护 -1禁用
	Flag             int    `:"flag" json:"flag"`                           // 删除标识
	CreatedAt        string `:"created_at" json:"created_at"`
	UpdatedAt        string `:"updated_at" json:"updated_at"`
	Ak               string `:"ak" json:"ak"`
	Sk               string `:"sk" json:"sk"`
	PayName          string `:"pay_name" json:"pay_name"`
	PayAk            string `:"pay_ak" json:"pay_ak"`
	PaySk            string `:"pay_sk" json:"pay_sk"`
	PayNotifyUrl     string `:"pay_notify_url" json:"pay_notify_url"` // 支付回调
	PayNotifyFunc    string `:"pay_notify_func" json:"pay_notify_func"`
}

type Admin_log struct {
	Id         int    `:"id" json:"id"`                   // 主键
	ActionId   int    `:"action_id" json:"action_id"`     // 行为id
	UserId     int    `:"user_id" json:"user_id"`         // 执行用户id
	ActionIp   int    `:"action_ip" json:"action_ip"`     // 执行行为者ip
	Model      string `:"model" json:"model"`             // 触发行为的表
	RecordId   int    `:"record_id" json:"record_id"`     // 触发行为的数据id
	Remark     string `:"remark" json:"remark"`           // 日志备注
	Status     int    `:"status" json:"status"`           // 状态
	CreateTime int    `:"create_time" json:"create_time"` // 执行行为的时间
}

type Admin_message struct {
	Id         int    `:"id" json:"id"`
	UidReceive int    `:"uid_receive" json:"uid_receive"` // 接收消息的用户id
	UidSend    int    `:"uid_send" json:"uid_send"`       // 发送消息的用户id
	Type       string `:"type" json:"type"`               // 消息分类
	Content    string `:"content" json:"content"`         // 消息内容
	Status     int    `:"status" json:"status"`           // 状态
	CreateTime int    `:"create_time" json:"create_time"` // 创建时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
	ReadTime   int    `:"read_time" json:"read_time"`     // 阅读时间
}

type Marketing_packet struct {
	Id          int     `:"id" json:"id"`
	Cuid        int     `:"cuid" json:"cuid"`
	Price       float64 `:"price" json:"price"`
	Title       string  `:"title" json:"title"`
	Describe    string  `:"describe" json:"describe"` // 描述
	Status      int     `:"status" json:"status"`     // 0正常 1使用 -1过期或者禁用
	Type        string  `:"type" json:"type"`         // 来源渠道
	Flag        int     `:"flag" json:"flag"`         // 删除标识
	CreatedAt   string  `:"created_at" json:"created_at"`
	UpdatedAt   string  `:"updated_at" json:"updated_at"`
	PlatformKey string  `:"platform_key" json:"platform_key"`
	PacketNo    string  `:"packet_no" json:"packet_no"`
}

type Marketing_score struct {
	Id          int    `:"id" json:"id"`
	Title       string `:"title" json:"title"`               // 获取积分标题
	PlatformKey string `:"platform_key" json:"platform_key"` // 平台key
	Score       int    `:"score" json:"score"`               // 获取积分
	Flag        int    `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	Type        int    `:"type" json:"type"` // 0获取 1支出
	Cuid        int    `:"cuid" json:"cuid"`
	ScoreRuleId int    `:"score_rule_id" json:"score_rule_id"` // 规则id
	OldScore    int    `:"old_score" json:"old_score"`         // 获取之前的积分
}

type Ucenter_accounts struct {
	Id          int     `:"id" json:"id"`
	Cuid        int     `:"cuid" json:"cuid"`                 // ucenter uid
	PlatformKey string  `:"platform_key" json:"platform_key"` // 牵扯的平台
	Type        int     `:"type" json:"type"`                 // 0直接收益 1分润收益 2师徒收益 10现金红包收益 50充值收益 100提现 101购买商品
	Level       int     `:"level" json:"level"`               // 收益等级 比如 0一级分享收益 1二级分享收益
	Content     string  `:"content" json:"content"`           // 详细内容
	Describe    string  `:"describe" json:"describe"`         // 描述  主要是显示这里
	ProjectId   int     `:"project_id" json:"project_id"`     // 项目id
	OrderId     int     `:"order_id" json:"order_id"`         // 订单id
	OrderNo     int     `:"order_no" json:"order_no"`         // 订单编号
	Price       float64 `:"price" json:"price"`               // 金额
	IsDz        int     `:"is_dz" json:"is_dz"`               // 是否到账 1是
	SourceCuid  int     `:"source_cuid" json:"source_cuid"`   // 来源用户。比如是谁分享产生的给你费用
	ProjectName string  `:"project_name" json:"project_name"` // 项目名称
	Title       string  `:"title" json:"title"`               // 标题
	Flag        int     `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string  `:"created_at" json:"created_at"`
	UpdatedAt   string  `:"updated_at" json:"updated_at"`
	AccountNo   string  `:"account_no" json:"account_no"` // 订单号
}

type Ucenter_tixian struct {
	Id          int     `:"id" json:"id"`
	PlatformKey string  `:"platform_key" json:"platform_key"` // 平台key
	Cuid        int     `:"cuid" json:"cuid"`                 // ucenter id
	Price       float64 `:"price" json:"price"`               // 提现金额
	FirstMoney  float64 `:"first_money" json:"first_money"`   // 提现前金额
	Status      int     `:"status" json:"status"`             // 状态0提现中 1成功 -1 失败
	TixianFlag  string  `:"tixian_flag" json:"tixian_flag"`   // 空或者bank 银行 wechat 微信钱包 alipay 支付宝钱包
	BankName    string  `:"bank_name" json:"bank_name"`       // 银行卡银行
	BankId      string  `:"bank_id" json:"bank_id"`           // 银行卡号
	BankUser    string  `:"bank_user" json:"bank_user"`       // 银行卡姓名
	BankAddress string  `:"bank_address" json:"bank_address"` // 开户行详细地址
	Flag        int     `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string  `:"created_at" json:"created_at"`
	UpdatedAt   string  `:"updated_at" json:"updated_at"`
	TixianNo    string  `:"tixian_no" json:"tixian_no"`
}

type Admin_packet struct {
	Id         int    `:"id" json:"id"`
	Name       string `:"name" json:"name"`             // 数据包名
	Title      string `:"title" json:"title"`           // 数据包标题
	Author     string `:"author" json:"author"`         // 作者
	AuthorUrl  string `:"author_url" json:"author_url"` // 作者url
	Version    string `:"version" json:"version"`
	Tables     string `:"tables" json:"tables"`           // 数据表名
	CreateTime int    `:"create_time" json:"create_time"` // 创建时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
	Status     int    `:"status" json:"status"`           // 状态
}

type Admin_user struct {
	Id            int     `:"id" json:"id"`
	Username      string  `:"username" json:"username"`               // 用户名
	Nickname      string  `:"nickname" json:"nickname"`               // 昵称
	Password      string  `:"password" json:"password"`               // 密码
	Email         string  `:"email" json:"email"`                     // 邮箱地址
	EmailBind     int     `:"email_bind" json:"email_bind"`           // 是否绑定邮箱地址
	Mobile        string  `:"mobile" json:"mobile"`                   // 手机号码
	MobileBind    int     `:"mobile_bind" json:"mobile_bind"`         // 是否绑定手机号码
	Avatar        int     `:"avatar" json:"avatar"`                   // 头像
	Money         float64 `:"money" json:"money"`                     // 余额
	Score         int     `:"score" json:"score"`                     // 积分
	Role          int     `:"role" json:"role"`                       // 角色ID
	Group         int     `:"group" json:"group"`                     // 部门id
	SignupIp      int     `:"signup_ip" json:"signup_ip"`             // 注册ip
	CreateTime    int     `:"create_time" json:"create_time"`         // 创建时间
	UpdateTime    int     `:"update_time" json:"update_time"`         // 更新时间
	LastLoginTime int     `:"last_login_time" json:"last_login_time"` // 最后一次登录时间
	LastLoginIp   int     `:"last_login_ip" json:"last_login_ip"`     // 登录ip
	Sort          int     `:"sort" json:"sort"`                       // 排序
	Status        int     `:"status" json:"status"`                   // 状态：0禁用，1启用
}

type Partner_area struct {
	Id         int    `:"id" json:"id"`
	PartenrId  int    `:"partenr_id" json:"partenr_id"`
	AreaType   int    `:"area_type" json:"area_type"`     // 0 全国 1全省 2市 3区 4小区
	ProvinceId int    `:"province_id" json:"province_id"` // 省id
	CityId     int    `:"city_id" json:"city_id"`         // 市id
	AreaId     int    `:"area_id" json:"area_id"`         // 区id
	HousingId  int    `:"housing_id" json:"housing_id"`   // 小区id
	Flag       int    `:"flag" json:"flag"`               // 删除标识
	CreatedAt  string `:"created_at" json:"created_at"`
	UpdatedAt  string `:"updated_at" json:"updated_at"`
}

type Sm_service_log struct {
	Id             int    `:"id" json:"id"`
	ServiceId      int    `:"service_id" json:"service_id"`
	ServiceSkuId   int    `:"service_sku_id" json:"service_sku_id"`     // sku id
	ServiceAreaId  int    `:"service_area_id" json:"service_area_id"`   // 区域id
	UsersServiceId int    `:"users_service_id" json:"users_service_id"` // 用户服务表id
	Cuid           int    `:"cuid" json:"cuid"`
	Content        string `:"content" json:"content"`
	Pics           string `:"pics" json:"pics"`
	Avatar         string `:"avatar" json:"avatar"`     // 备份头像
	Nickname       string `:"nickname" json:"nickname"` // 备份昵称
	Flag           int    `:"flag" json:"flag"`         // 删除标识
	CreatedAt      string `:"created_at" json:"created_at"`
	UpdatedAt      string `:"updated_at" json:"updated_at"`
}

type Sm_users_service struct {
	Id               int     `:"id" json:"id"`
	ServiceId        int     `:"service_id" json:"service_id"`           // 服务id
	ServiceSkuId     int     `:"service_sku_id" json:"service_sku_id"`   // sku id
	ServiceAreaId    int     `:"service_area_id" json:"service_area_id"` // 区域id
	ApptId           int     `:"appt_id" json:"appt_id"`
	CartId           int     `:"cart_id" json:"cart_id"` // 购物车对应id
	ServiceLogo      string  `:"service_logo" json:"service_logo"`
	ServiceTitle     string  `:"service_title" json:"service_title"` // 标题
	ApptTitle        string  `:"appt_title" json:"appt_title"`
	CostPrice        float64 `:"cost_price" json:"cost_price"`             // 原价单价
	UnitPrice        float64 `:"unit_price" json:"unit_price"`             // 单价
	TotalCostPrice   float64 `:"total_cost_price" json:"total_cost_price"` // 全部原价
	TotalPrice       float64 `:"total_price" json:"total_price"`           // 全部价格
	OrderId          int     `:"order_id" json:"order_id"`
	Cuid             int     `:"cuid" json:"cuid"`
	Status           int     `:"status" json:"status"`       // 0 未付款 1付款(拼团属于还没有拼成)  5已确认（拼团拼成）  8已接单 10已服务 待处理
	WorkCuid         int     `:"work_cuid" json:"work_cuid"` // 分配工人uid
	Flag             int     `:"flag" json:"flag"`           // 删除标识
	NotAppt          int     `:"not_appt" json:"not_appt"`   // 是否暂时不选择服务区间 0否 1是
	CreatedAt        string  `:"created_at" json:"created_at"`
	UpdatedAt        string  `:"updated_at" json:"updated_at"`
	Price            float64 `:"price" json:"price"`             // 实际付款价格
	Num              int     `:"num" json:"num"`                 // 数量
	GjPrice          float64 `:"gj_price" json:"gj_price"`       // 工匠金额
	IsGjPrice        int     `:"is_gj_price" json:"is_gj_price"` // 是否固定工匠金额
	SkuName          string  `:"sku_name" json:"sku_name"`       // sku 名称
	AddressId        int     `:"address_id" json:"address_id"`
	ServiceNo        string  `:"service_no" json:"service_no"`
	IsComment        int     `:"is_comment" json:"is_comment"`           // 是否评论
	IsWorkComment    int     `:"is_work_comment" json:"is_work_comment"` // 是否商家评论
	Qianzi           string  `:"qianzi" json:"qianzi"`
	IsGp             int     `:"is_gp" json:"is_gp"`                             // 是否团购
	UsersServiceGpId int     `:"users_service_gp_id" json:"users_service_gp_id"` // 用户团购id
	ServiceGpId      int     `:"service_gp_id" json:"service_gp_id"`             // 团购id
}

type Ucenter_users struct {
	Id             int     `:"id" json:"id"`
	Username       string  `:"username" json:"username"`               // 账号
	Password       string  `:"password" json:"password"`               // 密码
	Mobile         string  `:"mobile" json:"mobile"`                   // 手机号
	Nickname       string  `:"nickname" json:"nickname"`               // 昵称
	Email          string  `:"email" json:"email"`                     // 邮箱
	Avatar         string  `:"avatar" json:"avatar"`                   // 头像
	Gender         string  `:"gender" json:"gender"`                   // 性别
	Status         int     `:"status" json:"status"`                   // 状态 0停用 1启用
	RoleType       int     `:"role_type" json:"role_type"`             // 0无权限 1小区合伙人 2区县级合伙人 3城市合伙人 4省合伙人
	Score          int     `:"score" json:"score"`                     // 积分
	Money          float64 `:"money" json:"money"`                     // 余额
	OkMoney        float64 `:"ok_money" json:"ok_money"`               // 可提现余额
	NoMoney        float64 `:"no_money" json:"no_money"`               // 不可提现金额
	LastLoginIp    string  `:"last_login_ip" json:"last_login_ip"`     // 最后一次登录ip
	LastLoginTime  int     `:"last_login_time" json:"last_login_time"` // 最后一次登录时间戳
	LastLongitude  float64 `:"last_longitude" json:"last_longitude"`   // 最后一次经度
	LastLatitude   float64 `:"last_latitude" json:"last_latitude"`     // 最后一次维度
	IsAuth         int     `:"is_auth" json:"is_auth"`                 // 是否实名认证 0 否 1审核 2通过 -1拒绝
	IdcardTop      string  `:"idcard_top" json:"idcard_top"`           // 身份证正面
	IdcardBom      string  `:"idcard_bom" json:"idcard_bom"`           // 身份证背面
	IdcardId       string  `:"idcard_id" json:"idcard_id"`             // 身份证号
	ShareOne       int     `:"share_one" json:"share_one"`             // 一级分享
	ShareTwo       int     `:"share_two" json:"share_two"`             // 二级分享
	StOne          int     `:"st_one" json:"st_one"`                   // 一级师徒
	StTwo          int     `:"st_two" json:"st_two"`                   // 二级师徒
	UserKey        string  `:"user_key" json:"user_key"`               // 用户注册唯一key
	WechatUnionid  string  `:"wechat_unionid" json:"wechat_unionid"`   // 微信相关unionid
	RegType        int     `:"reg_type" json:"reg_type"`               // 注册类型 0手机号验证码 1账号
	RegSource      string  `:"reg_source" json:"reg_source"`           // 注册来源 例如 手机 微信 小程序
	Flag           int     `:"flag" json:"flag"`                       // 删除标识
	CreatedAt      string  `:"created_at" json:"created_at"`
	UpdatedAt      string  `:"updated_at" json:"updated_at"`
	RegPlatformKey string  `:"reg_platform_key" json:"reg_platform_key"` // 从哪个平台注册的
	BindUserinfo   int     `:"bind_userinfo" json:"bind_userinfo"`       // 是否绑定用户信息
	IsVip          int     `:"is_vip" json:"is_vip"`                     // 是否是vip
	VipEndTime     string  `:"vip_end_time" json:"vip_end_time"`         // vip到期时间
	IsPayPassword  int     `:"is_pay_password" json:"is_pay_password"`   // 是否填写支付密码
	PayPassword    string  `:"pay_password" json:"pay_password"`         // 支付密码
}

type Admin_attachment struct {
	Id         int    `:"id" json:"id"`
	Uid        int    `:"uid" json:"uid"`                 // 用户id
	Name       string `:"name" json:"name"`               // 文件名
	Module     string `:"module" json:"module"`           // 模块名，由哪个模块上传的
	Path       string `:"path" json:"path"`               // 文件路径
	Thumb      string `:"thumb" json:"thumb"`             // 缩略图路径
	Url        string `:"url" json:"url"`                 // 文件链接
	Mime       string `:"mime" json:"mime"`               // 文件mime类型
	Ext        string `:"ext" json:"ext"`                 // 文件类型
	Size       int    `:"size" json:"size"`               // 文件大小
	Md5        string `:"md5" json:"md5"`                 // 文件md5
	Sha1       string `:"sha1" json:"sha1"`               // sha1 散列值
	Driver     string `:"driver" json:"driver"`           // 上传驱动
	Download   int    `:"download" json:"download"`       // 下载次数
	CreateTime int    `:"create_time" json:"create_time"` // 上传时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
	Sort       int    `:"sort" json:"sort"`               // 排序
	Status     int    `:"status" json:"status"`           // 状态
	Width      int    `:"width" json:"width"`             // 图片宽度
	Height     int    `:"height" json:"height"`           // 图片高度
}

type Admin_menu struct {
	Id         int    `:"id" json:"id"`
	Pid        int    `:"pid" json:"pid"`                 // 上级菜单id
	Module     string `:"module" json:"module"`           // 模块名称
	Title      string `:"title" json:"title"`             // 菜单标题
	Icon       string `:"icon" json:"icon"`               // 菜单图标
	UrlType    string `:"url_type" json:"url_type"`       // 链接类型（link：外链，module：模块）
	UrlValue   string `:"url_value" json:"url_value"`     // 链接地址
	UrlTarget  string `:"url_target" json:"url_target"`   // 链接打开方式：_blank,_self
	OnlineHide int    `:"online_hide" json:"online_hide"` // 网站上线后是否隐藏
	CreateTime int    `:"create_time" json:"create_time"` // 创建时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
	Sort       int    `:"sort" json:"sort"`               // 排序
	SystemMenu int    `:"system_menu" json:"system_menu"` // 是否为系统菜单，系统菜单不可删除
	Status     int    `:"status" json:"status"`           // 状态
	Params     string `:"params" json:"params"`           // 参数
}

type Face_users struct {
	Id        int    `:"id" json:"id"`
	Name      string `:"name" json:"name"`           // 姓名
	Mobile    string `:"mobile" json:"mobile"`       // 手机号
	Feature   string `:"feature" json:"feature"`     // 特征
	CloudUid  string `:"cloud_uid" json:"cloud_uid"` // 云端用于匹配的id
	CloudId   int    `:"cloud_id" json:"cloud_id"`   // 云端id
	Flag      int    `:"flag" json:"flag"`           // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
	AvatarUrl string `:"avatar_url" json:"avatar_url"` // 头像url
}

type Sm_service_appt struct {
	Id            int    `:"id" json:"id"`
	ApptItemTplId int    `:"appt_item_tpl_id" json:"appt_item_tpl_id"` // 对应服务时间名
	BeginTime     string `:"begin_time" json:"begin_time"`             // 开始时间
	EndTime       string `:"end_time" json:"end_time"`                 // 结束时间
	Title         string `:"title" json:"title"`                       // 用于前台展示
	Sort          int    `:"sort" json:"sort"`                         // 排序
	IsShow        int    `:"is_show" json:"is_show"`                   // 是否显示
	Num           int    `:"num" json:"num"`                           // 可预约次数
	Status        int    `:"status" json:"status"`                     // 0待预约 1已约满  9暂停预约 -1禁止预约
	UseNum        int    `:"use_num" json:"use_num"`                   // 已预约人数
	AreaType      int    `:"area_type" json:"area_type"`               // 限定区域范围等级 0全国 1省 2市 3区 4小区
	ProvinceId    int    `:"province_id" json:"province_id"`           // 省id
	CityId        int    `:"city_id" json:"city_id"`                   // 市id
	AreaId        int    `:"area_id" json:"area_id"`                   // 区id
	HousingId     int    `:"housing_id" json:"housing_id"`             // 小区id
	CreatedAt     string `:"created_at" json:"created_at"`
	ServiceId     int    `:"service_id" json:"service_id"`         // 服务id
	ServiceSkuId  int    `:"service_sku_id" json:"service_sku_id"` // sku id
	UpdatedAt     string `:"updated_at" json:"updated_at"`
	Flag          int    `:"flag" json:"flag"`               // 删除标识
	ApptDate      string `:"appt_date" json:"appt_date"`     // 预约日期
	ApptTplId     int    `:"appt_tpl_id" json:"appt_tpl_id"` // 预生成模板id
	ServiceAreaId int    `:"service_area_id" json:"service_area_id"`
}

type Sm_service_area struct {
	Id         int    `:"id" json:"id"`
	Title      string `:"title" json:"title"`
	ServiceId  int    `:"service_id" json:"service_id"`   // 服务id
	AreaType   int    `:"area_type" json:"area_type"`     // 0 全国 1全省 2市 3区 4小区
	ProvinceId int    `:"province_id" json:"province_id"` // 省id
	CityId     int    `:"city_id" json:"city_id"`         // 市id
	AreaId     int    `:"area_id" json:"area_id"`         // 区id
	HousingId  int    `:"housing_id" json:"housing_id"`   // 小区id
	Flag       int    `:"flag" json:"flag"`               // 删除标识
	CreatedAt  string `:"created_at" json:"created_at"`
	UpdatedAt  string `:"updated_at" json:"updated_at"`
}

type Ucenter_commission struct {
	Id            int    `:"id" json:"id"`
	PlatformKey   string `:"platform_key" json:"platform_key"`       // 平台key
	AllShareScale int    `:"all_share_scale" json:"all_share_scale"` // 分享最大比例
	AllOptScale   int    `:"all_opt_scale" json:"all_opt_scale"`     // 其他收益最大比例
	ShareLv1Scale int    `:"share_lv1_scale" json:"share_lv1_scale"` // 分享一级
	ShareLv2Scale int    `:"share_lv2_scale" json:"share_lv2_scale"` // 分享二级
	OptLv1Scale   int    `:"opt_lv1_scale" json:"opt_lv1_scale"`     // 其他收益一级
	OptLv2Scale   int    `:"opt_lv2_scale" json:"opt_lv2_scale"`     // 其他收益二级
	Status        int    `:"status" json:"status"`                   // 0停用 1启用
	Flag          int    `:"flag" json:"flag"`                       // 删除标识
	CreatedAt     string `:"created_at" json:"created_at"`
	UpdatedAt     string `:"updated_at" json:"updated_at"`
}

type Marketing_score_rule struct {
	Id          int    `:"id" json:"id"`
	Name        string `:"name" json:"name"`                 // 积分模板名称
	ClientTitle string `:"client_title" json:"client_title"` // 获取积分标题
	PlatformKey string `:"platform_key" json:"platform_key"` // 平台key
	Score       int    `:"score" json:"score"`               // 获取积分
	Status      int    `:"status" json:"status"`             // 是否开启
	IsOne       int    `:"is_one" json:"is_one"`             // 是否只能获取一次
	Flag        int    `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	ScoreKey    string `:"score_key" json:"score_key"` // 唯一标识key
}

type System_log struct {
	Id        int    `:"id" json:"id"`
	Content   string `:"content" json:"content"`
	Flag      int    `:"flag" json:"flag"` // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
}

type Ucenter_vip struct {
	Id        int     `:"id" json:"id"`
	Title     string  `:"title" json:"title"`           // 标题
	VipDay    int     `:"vip_day" json:"vip_day"`       // 时长
	CostPrice float64 `:"cost_price" json:"cost_price"` // 原价
	Price     float64 `:"price" json:"price"`           // 现价 支付价格
	Describe  string  `:"describe" json:"describe"`     // 描述
	Sort      int     `:"sort" json:"sort"`             // 排序
	Flag      int     `:"flag" json:"flag"`             // 删除标识
	CreatedAt string  `:"created_at" json:"created_at"`
	UpdatedAt string  `:"updated_at" json:"updated_at"`
}

type Helper struct {
	Id          int    `:"id" json:"id"`
	CatId       int    `:"cat_id" json:"cat_id"`
	Title       string `:"title" json:"title"`     // 标题
	Content     string `:"content" json:"content"` // 富文本
	Logo        string `:"logo" json:"logo"`       // logo
	Sort        int    `:"sort" json:"sort"`       // 排序
	Flag        int    `:"flag" json:"flag"`       // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	PlatformKey string `:"platform_key" json:"platform_key"`
	IsTopic     int    `:"is_topic" json:"is_topic"`
	IsShow      int    `:"is_show" json:"is_show"` // 是否显示
}

type Marketing_coupon_queue struct {
	Id          int     `:"id" json:"id"`
	CouponTplId int     `:"coupon_tpl_id" json:"coupon_tpl_id"`
	Logo        string  `:"logo" json:"logo"`
	Title       string  `:"title" json:"title"`               // 优惠券名称
	PlatformKey string  `:"platform_key" json:"platform_key"` // 平台key
	Price       float64 `:"price" json:"price"`               // 优惠金额/最大优惠金额
	FullPrice   float64 `:"full_price" json:"full_price"`     // 满金额条件
	Type        int     `:"type" json:"type"`                 // 0满减 1全局折扣 2满折
	Zkb         int     `:"zkb" json:"zkb"`                   // 折扣比 0 100
	ProjectId   int     `:"project_id" json:"project_id"`     // 对应项目id （上门端是skuid 或者serviceid）
	ProjectType int     `:"project_type" json:"project_type"` // 对应项目类型。每个平台不同  （上门端0 sku 1 service）
	EndTime     int     `:"end_time" json:"end_time"`         // 到期时间。被转换的时间戳
	Describe    string  `:"describe" json:"describe"`         // 描述
	Num         int     `:"num" json:"num"`                   // 最大发放数量
	QueueType   int     `:"queue_type" json:"queue_type"`     // 0用户列表 1省 2市 3区 10全国
	IsSend      int     `:"is_send" json:"is_send"`           // 是否发放成功
	ProvinceId  int     `:"province_id" json:"province_id"`   // 省id
	CityId      int     `:"city_id" json:"city_id"`           // 市id
	AreaId      int     `:"area_id" json:"area_id"`           // 区id
	Cuids       string  `:"cuids" json:"cuids"`               // 用户数组，最大不可超过2万人
	Flag        int     `:"flag" json:"flag"`                 // 删除标识
	UpdatedAt   string  `:"updated_at" json:"updated_at"`
	CreatedAt   string  `:"created_at" json:"created_at"`
	QueueName   string  `:"queue_name" json:"queue_name"` // 队列名称
	UseNum      int     `:"use_num" json:"use_num"`       // 已经发放数量
	IsRun       int     `:"is_run" json:"is_run"`         // 是否在队列运行中
	ErrMsg      string  `:"err_msg" json:"err_msg"`
}

type Message_tpl struct {
	Id              int    `:"id" json:"id"`
	MessageKey      string `:"message_key" json:"message_key"` // 唯一标识
	Title           string `:"title" json:"title"`
	Desc            string `:"desc" json:"desc"`
	Content         string `:"content" json:"content"`
	AppType         int    `:"app_type" json:"app_type"` // 0用户端 1商家端
	IsMsg           int    `:"is_msg" json:"is_msg"`     // 是否发送消息
	MessageType     string `:"message_type" json:"message_type"`
	PathType        string `:"path_type" json:"path_type"`                 // 路径类型
	PathId          string `:"path_id" json:"path_id"`                     // 路径id 或者路径
	IsFormId        int    `:"is_form_id" json:"is_form_id"`               // 是否使用小程序模板id推送
	SmallTplId      string `:"small_tpl_id" json:"small_tpl_id"`           // 小程序模板id
	SmallTplContent string `:"small_tpl_content" json:"small_tpl_content"` // 小程序模板内容 json
	IsSms           int    `:"is_sms" json:"is_sms"`                       // 是否发送短信
	SmsContent      string `:"sms_content" json:"sms_content"`             // 短信内容
	IsEmail         int    `:"is_email" json:"is_email"`                   // 是否发送短信
	EmailTitle      string `:"email_title" json:"email_title"`
	EmailContent    string `:"email_content" json:"email_content"`
	CreatedAt       string `:"created_at" json:"created_at"`
	UpdatedAt       string `:"updated_at" json:"updated_at"`
	DeletedAt       string `:"deleted_at" json:"deleted_at"`
	Flag            int    `:"flag" json:"flag"`         // -1删除
	IsUcId          int    `:"is_uc_id" json:"is_uc_id"` // 是否使用用户平台
	PlatformKey     string `:"platform_key" json:"platform_key"`
	SmallTplPath    string `:"small_tpl_path" json:"small_tpl_path"`
}

type Partner_accounts struct {
	Id         int     `:"id" json:"id"`
	PartnerId  int     `:"partner_id" json:"partner_id"`   // 合伙人id
	OrderId    int     `:"order_id" json:"order_id"`       // 订单id
	Price      float64 `:"price" json:"price"`             // 到账金额
	AreaType   int     `:"area_type" json:"area_type"`     // 0 全国 1全省 2市 3区 4小区
	HousingId  int     `:"housing_id" json:"housing_id"`   // 小区id
	ProvinceId int     `:"province_id" json:"province_id"` // 省id
	CityId     int     `:"city_id" json:"city_id"`         // 市id
	AreaId     int     `:"area_id" json:"area_id"`         // 区id
	Flag       int     `:"flag" json:"flag"`               // 删除标识
	CreatedAt  string  `:"created_at" json:"created_at"`
	UpdatedAt  string  `:"updated_at" json:"updated_at"`
}

type Sm_service_appt_item_tpl struct {
	Id        int    `:"id" json:"id"`
	ApptTplId int    `:"appt_tpl_id" json:"appt_tpl_id"` // 模板tplid
	BeginTime string `:"begin_time" json:"begin_time"`   // 开始时间
	EndTime   string `:"end_time" json:"end_time"`       // 结束时间
	Title     string `:"title" json:"title"`             // 用于前台展示
	Sort      int    `:"sort" json:"sort"`               // 排序
	IsShow    int    `:"is_show" json:"is_show"`         // 是否显示
	Num       int    `:"num" json:"num"`                 // 可预约次数
	Flag      int    `:"flag" json:"flag"`               // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
}

type Ucenter_action_log struct {
	Id        int    `:"id" json:"id"`
	Cuid      int    `:"cuid" json:"cuid"`
	Action    string `:"action" json:"action"`   // 行为
	Content   string `:"content" json:"content"` // 内容
	Title     string `:"title" json:"title"`
	Flag      int    `:"flag" json:"flag"` // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
}

type Admin_icon struct {
	Id         int    `:"id" json:"id"`
	Name       string `:"name" json:"name"`               // 图标名称
	Url        string `:"url" json:"url"`                 // 图标css地址
	Prefix     string `:"prefix" json:"prefix"`           // 图标前缀
	FontFamily string `:"font_family" json:"font_family"` // 字体名
	CreateTime int    `:"create_time" json:"create_time"` // 创建时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
	Status     int    `:"status" json:"status"`           // 状态
}

type Helper_cats struct {
	Id          int    `:"id" json:"id"`
	PlatformKey string `:"platform_key" json:"platform_key"`
	CatName     string `:"cat_name" json:"cat_name"` // 分类名称
	CatLogo     string `:"cat_logo" json:"cat_logo"` // 分类logo
	CatDesc     string `:"cat_desc" json:"cat_desc"` // 分类描述
	IsShow      int    `:"is_show" json:"is_show"`   // 是否显示
	Sort        int    `:"sort" json:"sort"`         // 排序
	Flag        int    `:"flag" json:"flag"`         // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
}

type Sm_service_gp struct {
	Id        int     `:"id" json:"id"`
	ServiceId int     `:"service_id" json:"service_id"`   // 服务id
	Title     string  `:"title" json:"title"`             // 团购标题
	Content   string  `:"content" json:"content"`         // 富文本
	AllNum    int     `:"all_num" json:"all_num"`         // 满足多少成大团
	ShareNum  int     `:"share_num" json:"share_num"`     // 满足多少成小团
	Type      int     `:"type" json:"type"`               // 0满大团 1满小团 2满大小团 10不管团多少都tm成
	Price     float64 `:"price" json:"price"`             // 价格
	CostPrice float64 `:"cost_price" json:"cost_price"`   // 原价
	BeginTime string  `:"begin_time" json:"begin_time"`   // 开始时间
	EndTime   string  `:"end_time" json:"end_time"`       // 到期时间
	LimitType int     `:"limit_type" json:"limit_type"`   // 拼团区域限制 0不限制sku相等下单  1限制sku相等下单
	UseMaxNum int     `:"use_max_num" json:"use_max_num"` // 当前拼团一个用户最大可拼次数
	Status    int     `:"status" json:"status"`           // 0未开始  1已开始 10已结束
	Flag      int     `:"flag" json:"flag"`               // 删除标识
	CreatedAt string  `:"created_at" json:"created_at"`
	UpdatedAt string  `:"updated_at" json:"updated_at"`
	IsShow    int     `:"is_show" json:"is_show"`         // 是否显示
	Logo      string  `:"logo" json:"logo"`               // 团购Logo
	Banner    string  `:"banner" json:"banner"`           // 团购Banner
	UseAllNum int     `:"use_all_num" json:"use_all_num"` // 大团购数量
	IsTopic   int     `:"is_topic" json:"is_topic"`
	SkuName   string  `:"sku_name" json:"sku_name"` // sku name
}

type Sm_service_gp_sku struct {
	Id            int    `:"id" json:"id"`
	ServiceId     int    `:"service_id" json:"service_id"`
	ServiceGpId   int    `:"service_gp_id" json:"service_gp_id"`
	ServiceAreaId int    `:"service_area_id" json:"service_area_id"` // 区域id
	ServiceSkuId  int    `:"service_sku_id" json:"service_sku_id"`   // sku id
	Flag          int    `:"flag" json:"flag"`                       // 删除标识
	CreatedAt     string `:"created_at" json:"created_at"`
	UpdatedAt     string `:"updated_at" json:"updated_at"`
}

type Admin_module struct {
	Id           int    `:"id" json:"id"`
	Name         string `:"name" json:"name"`                   // 模块名称（标识）
	Title        string `:"title" json:"title"`                 // 模块标题
	Icon         string `:"icon" json:"icon"`                   // 图标
	Description  string `:"description" json:"description"`     // 描述
	Author       string `:"author" json:"author"`               // 作者
	AuthorUrl    string `:"author_url" json:"author_url"`       // 作者主页
	Config       string `:"config" json:"config"`               // 配置信息
	Access       string `:"access" json:"access"`               // 授权配置
	Version      string `:"version" json:"version"`             // 版本号
	Identifier   string `:"identifier" json:"identifier"`       // 模块唯一标识符
	SystemModule int    `:"system_module" json:"system_module"` // 是否为系统模块
	CreateTime   int    `:"create_time" json:"create_time"`     // 创建时间
	UpdateTime   int    `:"update_time" json:"update_time"`     // 更新时间
	Sort         int    `:"sort" json:"sort"`                   // 排序
	Status       int    `:"status" json:"status"`               // 状态
}

type Admin_users struct {
	Id        int    `:"id" json:"id"`
	Username  string `:"username" json:"username"` // 账号
	Password  string `:"password" json:"password"` // 密码
	Mobile    string `:"mobile" json:"mobile"`     // 手机号
	Nickname  string `:"nickname" json:"nickname"` // 昵称
	Email     string `:"email" json:"email"`       // 邮箱
	Gender    string `:"gender" json:"gender"`     // 性别
	Status    int    `:"status" json:"status"`     // 状态 0停用 1启用
	Flag      int    `:"flag" json:"flag"`         // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
}

type Sm_service struct {
	Id           int     `:"id" json:"id"`
	Title        string  `:"title" json:"title"`                   // 标题
	Logo         string  `:"logo" json:"logo"`                     // 列表logo
	Icon         string  `:"icon" json:"icon"`                     // 首页icon
	Banner       string  `:"banner" json:"banner"`                 // banner列表 json
	CatsId       int     `:"cats_id" json:"cats_id"`               // 分类id
	Describe     string  `:"describe" json:"describe"`             // 描述
	Content      string  `:"content" json:"content"`               // 富文本内容
	IsVideo      int     `:"is_video" json:"is_video"`             // 是否显示视频banner
	VideoUrl     string  `:"video_url" json:"video_url"`           // 视频url
	ShowNum      int     `:"show_num" json:"show_num"`             // 显示数量
	PayNum       int     `:"pay_num" json:"pay_num"`               // 销售数量
	CollectNum   int     `:"collect_num" json:"collect_num"`       // 收藏数量
	ShareNum     int     `:"share_num" json:"share_num"`           // 分享数量
	MinPrice     float64 `:"min_price" json:"min_price"`           // 最小售价（用于限定最低单品购买以及显示）
	MinCostPrice float64 `:"min_cost_price" json:"min_cost_price"` // 最小销售原价
	ProjectType  int     `:"project_type" json:"project_type"`     // 0正常 1团购
	IsShow       int     `:"is_show" json:"is_show"`               // 是否显示
	Status       int     `:"status" json:"status"`                 // 0停售 1正常销售  10预售
	Flag         int     `:"flag" json:"flag"`                     // 删除标识
	CreatedAt    string  `:"created_at" json:"created_at"`
	UpdatedAt    string  `:"updated_at" json:"updated_at"`
	IsDelete     int     `:"is_delete" json:"is_delete"`     // 是否远程删除
	IsTopic      int     `:"is_topic" json:"is_topic"`       // 是否推荐
	IsNew        int     `:"is_new" json:"is_new"`           // 是否new
	Sort         int     `:"sort" json:"sort"`               // 排序
	AreaType     int     `:"area_type" json:"area_type"`     // 限定区域范围等级 0全国 1省 2市 3区 4小区
	ProvinceId   int     `:"province_id" json:"province_id"` // 省id
	CityId       int     `:"city_id" json:"city_id"`         // 市id
	AreaId       int     `:"area_id" json:"area_id"`         // 区id
	HousingId    int     `:"housing_id" json:"housing_id"`   // 小区id
}

type Sm_service_appt_tpl struct {
	Id        int    `:"id" json:"id"`
	TplName   string `:"tpl_name" json:"tpl_name"` // 模板名称
	Describe  string `:"describe" json:"describe"` // 描述
	Flag      int    `:"flag" json:"flag"`         // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
}

type Ucenter_collect struct {
	Id          int    `:"id" json:"id"`
	ProjectType string `:"project_type" json:"project_type"` // 项目类型
	PlatformKey string `:"platform_key" json:"platform_key"` // 平台key
	Cuid        int    `:"cuid" json:"cuid"`                 // 用户id
	ProjectId   int    `:"project_id" json:"project_id"`     // 项目id
	Flag        int    `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
}

type Ucenter_orders struct {
	Id            int     `:"id" json:"id"`
	OrderNo       string  `:"order_no" json:"order_no"`         // 订单编号
	PlatformKey   string  `:"platform_key" json:"platform_key"` // 平台key
	Cuid          int     `:"cuid" json:"cuid"`                 // ucenter id
	CouponKey     string  `:"coupon_key" json:"coupon_key"`     // 优惠券key
	CouponId      int     `:"coupon_id" json:"coupon_id"`       // 优惠券id
	CouponPrice   float64 `:"coupon_price" json:"coupon_price"` // 优惠券抵扣金额
	CostPrice     float64 `:"cost_price" json:"cost_price"`     // 原价
	UnitPrice     float64 `:"unit_price" json:"unit_price"`     // 单价
	Price         float64 `:"price" json:"price"`               // 现价 支付总价
	GoodsNum      float64 `:"goods_num" json:"goods_num"`       // 商品总数量
	PayType       int     `:"pay_type" json:"pay_type"`         // 支付类型 0线上 1线下
	PayPlatform   string  `:"pay_platform" json:"pay_platform"` // 支付平台 wechat alipay ...
	Status        int     `:"status" json:"status"`             // 状态 0下单 2审核通过 7已接单 8已发货 9已结算或者已收货
	IsPay         int     `:"is_pay" json:"is_pay"`             // 是否支付 0否 1线下提交 2已支付
	IsUComment    int     `:"is_u_comment" json:"is_u_comment"` // 是否用户评论 0否 1是
	IsMComment    int     `:"is_m_comment" json:"is_m_comment"` // 是否商家评论 0否1是
	Flag          int     `:"flag" json:"flag"`                 // 删除标识
	CreatedAt     string  `:"created_at" json:"created_at"`
	UpdatedAt     string  `:"updated_at" json:"updated_at"`
	PayTime       int     `:"pay_time" json:"pay_time"`             // 支付时间
	ServicePrice  float64 `:"service_price" json:"service_price"`   // 服务费
	SharePrice    float64 `:"share_price" json:"share_price"`       // 分享出去多少钱
	PlatformPrice float64 `:"platform_price" json:"platform_price"` // 平台受益
	ShareLv1      float64 `:"share_lv1" json:"share_lv1"`           // 一级分享
	ShareLv2      float64 `:"share_lv2" json:"share_lv2"`           // 二级分享
	OptLv1        float64 `:"opt_lv1" json:"opt_lv1"`               // 一级其他分享
	OptLv2        float64 `:"opt_lv2" json:"opt_lv2"`               // 二级其他分享
	ShareLv1Cuid  int     `:"share_lv1_cuid" json:"share_lv1_cuid"` // 一级分享用户id
	ShareLv2Cuid  int     `:"share_lv2_cuid" json:"share_lv2_cuid"`
	OptLv1Cuid    int     `:"opt_lv1_cuid" json:"opt_lv1_cuid"`
	OptLv2Cuid    int     `:"opt_lv2_cuid" json:"opt_lv2_cuid"`
	OrderType     int     `:"order_type" json:"order_type"` // 订单类型0常规购买订单  1团购 10VIP
	PayNo         string  `:"pay_no" json:"pay_no"`         // 线上付款订单
	Describe      string  `:"describe" json:"describe"`     // 描述
	ProjectId     int     `:"project_id" json:"project_id"`
}

type Address_city struct {
	Id        int     `:"id" json:"id"`
	Name      string  `:"name" json:"name"`           // 名称
	Pid       int     `:"pid" json:"pid"`             // 父类
	Level     int     `:"level" json:"level"`         // 0 省级 1城市 2区县级
	IsShow    int     `:"is_show" json:"is_show"`     // 是否显示
	Flag      int     `:"flag" json:"flag"`           // 删除标识
	Longitude float64 `:"longitude" json:"longitude"` // 经度
	Latitude  float64 `:"latitude" json:"latitude"`   // 维度
	Scale     int     `:"scale" json:"scale"`         // 分润
	CreatedAt string  `:"created_at" json:"created_at"`
	UpdatedAt string  `:"updated_at" json:"updated_at"`
}

type Admin_hook struct {
	Id          int    `:"id" json:"id"`
	Name        string `:"name" json:"name"`               // 钩子名称
	Plugin      string `:"plugin" json:"plugin"`           // 钩子来自哪个插件
	Description string `:"description" json:"description"` // 钩子描述
	System      int    `:"system" json:"system"`           // 是否为系统钩子
	CreateTime  int    `:"create_time" json:"create_time"` // 创建时间
	UpdateTime  int    `:"update_time" json:"update_time"` // 更新时间
	Status      int    `:"status" json:"status"`           // 状态
}

type Admin_plugin struct {
	Id          int    `:"id" json:"id"`
	Name        string `:"name" json:"name"`               // 插件名称
	Title       string `:"title" json:"title"`             // 插件标题
	Icon        string `:"icon" json:"icon"`               // 图标
	Description string `:"description" json:"description"` // 插件描述
	Author      string `:"author" json:"author"`           // 作者
	AuthorUrl   string `:"author_url" json:"author_url"`   // 作者主页
	Config      string `:"config" json:"config"`           // 配置信息
	Version     string `:"version" json:"version"`         // 版本号
	Identifier  string `:"identifier" json:"identifier"`   // 插件唯一标识符
	Admin       int    `:"admin" json:"admin"`             // 是否有后台管理
	CreateTime  int    `:"create_time" json:"create_time"` // 安装时间
	UpdateTime  int    `:"update_time" json:"update_time"` // 更新时间
	Sort        int    `:"sort" json:"sort"`               // 排序
	Status      int    `:"status" json:"status"`           // 状态
}

type Ucenter_openid struct {
	Id          int    `:"id" json:"id"`
	Cuid        int    `:"cuid" json:"cuid"`                 // ucenter id
	PlatformKey string `:"platform_key" json:"platform_key"` // 平台key
	Type        string `:"type" json:"type"`                 // 类型 wechat ali app
	Flag        int    `:"flag" json:"flag"`                 // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	Openid      string `:"openid" json:"openid"`
}

type Ucenter_optshare struct {
	Id           int    `:"id" json:"id"`
	Cuid         int    `:"cuid" json:"cuid"`                     // 用户id
	PlatformKey  string `:"platform_key" json:"platform_key"`     // 平台key
	ShareLv1Cuid int    `:"share_lv1_cuid" json:"share_lv1_cuid"` // 上级
	ShareLv2Cuid int    `:"share_lv2_cuid" json:"share_lv2_cuid"` // 上级二级
	Level        int    `:"level" json:"level"`                   // 分享等级 0一级 1二级
	Status       int    `:"status" json:"status"`                 // 0不产生分润 1产生分润
	Flag         int    `:"flag" json:"flag"`                     // 删除标识
	CreatedAt    string `:"created_at" json:"created_at"`
	UpdatedAt    string `:"updated_at" json:"updated_at"`
}

type Admin_config struct {
	Id         int    `:"id" json:"id"`
	Name       string `:"name" json:"name"`               // 名称
	Title      string `:"title" json:"title"`             // 标题
	Group      string `:"group" json:"group"`             // 配置分组
	Type       string `:"type" json:"type"`               // 类型
	Value      string `:"value" json:"value"`             // 配置值
	Options    string `:"options" json:"options"`         // 配置项
	Tips       string `:"tips" json:"tips"`               // 配置提示
	AjaxUrl    string `:"ajax_url" json:"ajax_url"`       // 联动下拉框ajax地址
	NextItems  string `:"next_items" json:"next_items"`   // 联动下拉框的下级下拉框名，多个以逗号隔开
	Param      string `:"param" json:"param"`             // 联动下拉框请求参数名
	Format     string `:"format" json:"format"`           // 格式，用于格式文本
	Table      string `:"table" json:"table"`             // 表名，只用于快速联动类型
	Level      int    `:"level" json:"level"`             // 联动级别，只用于快速联动类型
	Key        string `:"key" json:"key"`                 // 键字段，只用于快速联动类型
	Option     string `:"option" json:"option"`           // 值字段，只用于快速联动类型
	Pid        string `:"pid" json:"pid"`                 // 父级id字段，只用于快速联动类型
	Ak         string `:"ak" json:"ak"`                   // 百度地图appkey
	CreateTime int    `:"create_time" json:"create_time"` // 创建时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
	Sort       int    `:"sort" json:"sort"`               // 排序
	Status     int    `:"status" json:"status"`           // 状态：0禁用，1启用
}

type Admin_hook_plugin struct {
	Id         int    `:"id" json:"id"`
	Hook       string `:"hook" json:"hook"`               // 钩子id
	Plugin     string `:"plugin" json:"plugin"`           // 插件标识
	CreateTime int    `:"create_time" json:"create_time"` // 添加时间
	UpdateTime int    `:"update_time" json:"update_time"` // 更新时间
	Sort       int    `:"sort" json:"sort"`               // 排序
	Status     int    `:"status" json:"status"`           // 状态
}

type System_pages struct {
	Id          int    `:"id" json:"id"`
	Title       string `:"title" json:"title"`       // 标题
	Describe    string `:"describe" json:"describe"` // 描述
	Content     string `:"content" json:"content"`   // 富文本
	PlatformKey string `:"platform_key" json:"platform_key"`
	IsShow      int    `:"is_show" json:"is_show"`
	Flag        int    `:"flag" json:"flag"` // 删除标识
	CreatedAt   string `:"created_at" json:"created_at"`
	UpdatedAt   string `:"updated_at" json:"updated_at"`
	Sort        int    `:"sort" json:"sort"` // 排序
	Logo        string `:"logo" json:"logo"`
	PageKey     string `:"page_key" json:"page_key"` // 页面key
}

type Admin_access struct {
	Module string `:"module" json:"module"` // 模型名称
	Group  string `:"group" json:"group"`   // 权限分组标识
	Uid    int    `:"uid" json:"uid"`       // 用户id
	Nid    string `:"nid" json:"nid"`       // 授权节点id
	Tag    string `:"tag" json:"tag"`       // 分组标签
}

type Admin_icon_list struct {
	Id     int    `:"id" json:"id"`
	IconId int    `:"icon_id" json:"icon_id"` // 所属图标id
	Title  string `:"title" json:"title"`     // 图标标题
	Class  string `:"class" json:"class"`     // 图标类名
	Code   string `:"code" json:"code"`       // 图标关键词
}

type Admin_role struct {
	Id            int    `:"id" json:"id"`                         // 角色id
	Pid           int    `:"pid" json:"pid"`                       // 上级角色
	Name          string `:"name" json:"name"`                     // 角色名称
	Description   string `:"description" json:"description"`       // 角色描述
	MenuAuth      string `:"menu_auth" json:"menu_auth"`           // 菜单权限
	Sort          int    `:"sort" json:"sort"`                     // 排序
	CreateTime    int    `:"create_time" json:"create_time"`       // 创建时间
	UpdateTime    int    `:"update_time" json:"update_time"`       // 更新时间
	Status        int    `:"status" json:"status"`                 // 状态
	Access        int    `:"access" json:"access"`                 // 是否可登录后台
	DefaultModule int    `:"default_module" json:"default_module"` // 默认访问模块
}

type Face_check_log struct {
	Id        int    `:"id" json:"id"`
	Mac       string `:"mac" json:"mac"`             // 设备mac
	DeviceId  int    `:"device_id" json:"device_id"` // 设备id
	CloudUid  string `:"cloud_uid" json:"cloud_uid"`
	Photo     string `:"photo" json:"photo"`           // 图像
	CloudTime int    `:"cloud_time" json:"cloud_time"` // 云端时间
	PhotoHash string `:"photo_hash" json:"photo_hash"` // hash
	FaceUid   int    `:"face_uid" json:"face_uid"`
	Flag      int    `:"flag" json:"flag"` // 删除标识
	CreatedAt string `:"created_at" json:"created_at"`
	UpdatedAt string `:"updated_at" json:"updated_at"`
}

type Sm_service_sku struct {
	Id            int     `:"id" json:"id"`
	ServiceId     int     `:"service_id" json:"service_id"` // 服务id
	SkuName       string  `:"sku_name" json:"sku_name"`     // sku名称
	Describe      string  `:"describe" json:"describe"`     // 描述
	SkuLogo       string  `:"sku_logo" json:"sku_logo"`     // logo 小图标
	Price         float64 `:"price" json:"price"`           // 售价
	CostPrice     float64 `:"cost_price" json:"cost_price"` // 原价
	SpNum         int     `:"sp_num" json:"sp_num"`         // 单个人可以购买促销产品数量
	SpType        int     `:"sp_type" json:"sp_type"`       // 促销类型 0非促销 1新用户首单限定 2当前sku限定 3当前商品限定  10捆绑销售
	Sort          int     `:"sort" json:"sort"`             // 排序
	Flag          int     `:"flag" json:"flag"`             // 删除标识
	CreatedAt     string  `:"created_at" json:"created_at"`
	UpdatedAt     string  `:"updated_at" json:"updated_at"`
	Stock         int     `:"stock" json:"stock"`                     // 库存
	ApptTplId     int     `:"appt_tpl_id" json:"appt_tpl_id"`         // 预生成模板id
	ServiceAreaId int     `:"service_area_id" json:"service_area_id"` // 区域id
	GjPrice       float64 `:"gj_price" json:"gj_price"`               // 工匠金额
	IsGjPrice     int     `:"is_gj_price" json:"is_gj_price"`         // 是否固定工匠金额
}

type Sm_users_address struct {
	Id             int     `:"id" json:"id"`
	AdrName        string  `:"adr_name" json:"adr_name"`           // 选择地址名称
	AdrLatitude    float64 `:"adr_latitude" json:"adr_latitude"`   // 维度
	AdrLongitude   float64 `:"adr_longitude" json:"adr_longitude"` // 经度
	Address        string  `:"address" json:"address"`             // 详细门牌号
	Cuid           int     `:"cuid" json:"cuid"`
	IsDefault      int     `:"is_default" json:"is_default"` // 是否默认
	Name           string  `:"name" json:"name"`             // 姓名
	Mobile         string  `:"mobile" json:"mobile"`         // 手机号
	Flag           int     `:"flag" json:"flag"`             // 删除标识
	CreatedAt      string  `:"created_at" json:"created_at"`
	UpdatedAt      string  `:"updated_at" json:"updated_at"`
	AreaLevel      int     `:"area_level" json:"area_level"`             // 0省1市 2区 3小区
	AddreaaCheckId int     `:"addreaa_check_id" json:"addreaa_check_id"` // 对应的id city 或者housingid
}
