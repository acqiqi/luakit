package daoru

type Accounts struct {
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
