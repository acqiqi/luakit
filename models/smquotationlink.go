package models

import (
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type SmQuotationLink struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	QuotationId int     `json:"quotation_id"` // 报价id
	CatId       int     `json:"cat_id"`       // 分类id
	CatName     string  `json:"cat_name"`     // 分类名称
	Describe    string  `json:"describe"`     // 描述
	TotalPrice  float64 `json:"total_price"`
	TotalNum    int     `json:"total_num"`  // 总数
	TableJson   string  `json:"table_json"` // 数据集合
	LinkNo      string  `json:"link_no"`    // 编号
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(SmQuotationLink))
}

func GetSmQuotationLinkListByQId(qid int64) (row int64, v *[]SmQuotationLink, err error) {
	o := orm.NewOrm()
	v = &[]SmQuotationLink{}
	if row, err = o.QueryTable(new(SmQuotationLink)).Filter("QuotationId", qid).Filter("flag", 1).RelatedSel().All(v); err == nil {
		return row, v, nil
	}
	return row, nil, err
}
