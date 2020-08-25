package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type SmQuotation struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	TplId       int     `json:"tpl_id"` // 模板id
	Name        string  `json:"name"`   // 报价单名称
	Describe    string  `json:"describe"`
	QuotationNo string  `json:"quotation_no"` // 报价单号
	TotalPrice  float64 `json:"total_price"`  // 总价
	TotalNum    int     `json:"total_num"`    // 项目总数
	Status      int     `json:"status"`       // 0正常 1进行中 2签订
	UAddress    string  `json:"u_address"`    // 用户地址
	UMobile     string  `json:"u_mobile"`     // 用户手机号
	UName       string  `json:"u_name"`       // 用户姓名
	PartnerId   int     `json:"partner_id"`   // 合伙人id
	EndPrice    float64 `json:"end_price"`    // 最终价格
	FixedPrice  float64 `json:"fixed_price"`  //一口价
	ManagerBl   int     `json:"manager_bl"`   //管理费比例
	TaxBl       int     `json:"tax_bl"`       //税率
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SmQuotation))
}

// 用户服务订单
func GetSmQuotationById(id int64) (v *SmQuotation, err error) {
	o := orm.NewOrm()
	v = &SmQuotation{}
	if err = o.QueryTable(new(SmQuotation)).Filter("Id", id).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
