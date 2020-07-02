package models

import (
	"errors"
	"fmt"
	"luakit/utils"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UcenterUsers struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Username       string  `json:"username"`         // 账号
	Password       string  `json:"password"`         // 密码
	Mobile         string  `json:"mobile"`           // 手机号
	Nickname       string  `json:"nickname"`         // 昵称
	Email          string  `json:"email"`            // 邮箱
	Avatar         string  `json:"avatar"`           // 头像
	Gender         string  `json:"gender"`           // 性别
	Status         int     `json:"status"`           // 状态 0停用 1启用
	RoleType       int     `json:"role_type"`        // 0无权限 1小区合伙人 2区县级合伙人 3城市合伙人 4省合伙人
	Score          int     `json:"score"`            // 积分
	Money          float64 `json:"money"`            // 余额
	OkMoney        float64 `json:"ok_money"`         // 可提现余额
	NoMoney        float64 `json:"no_money"`         // 不可提现金额
	LastLoginIp    string  `json:"last_login_ip"`    // 最后一次登录ip
	LastLoginTime  int     `json:"last_login_time"`  // 最后一次登录时间戳
	LastLongitude  float64 `json:"last_longitude"`   // 最后一次经度
	LastLatitude   float64 `json:"last_latitude"`    // 最后一次维度
	IsAuth         int     `json:"is_auth"`          // 是否实名认证 0 否 1审核 2通过 -1拒绝
	IdcardTop      string  `json:"idcard_top"`       // 身份证正面
	IdcardBom      string  `json:"idcard_bom"`       // 身份证背面
	IdcardId       string  `json:"idcard_id"`        // 身份证号
	ShareOne       int     `json:"share_one"`        // 一级分享
	ShareTwo       int     `json:"share_two"`        // 二级分享
	StOne          int     `json:"st_one"`           // 一级师徒
	StTwo          int     `json:"st_two"`           // 二级师徒
	UserKey        string  `json:"user_key"`         // 用户注册唯一key
	WechatUnionid  string  `json:"wechat_unionid"`   // 微信相关unionid
	RegType        int     `json:"reg_type"`         // 注册类型 0手机号验证码 1账号
	RegSource      string  `json:"reg_source"`       // 注册来源 例如 手机 微信 小程序
	RegPlatformKey string  `json:"reg_platform_key"` // 从哪个平台注册的
	BindUserinfo   int     `json:"bind_userinfo"`    // 是否绑定用户信息
	IsVip          int     `json:"is_vip"`           // 是否是vip
	VipEndTime     string  `json:"vip_end_time"`     // vip到期时间
	IsPayPassword  int     `json:"is_pay_password"`  // 是否填写支付密码
	PayPassword    string  `json:"pay_password"`     // 支付密码
	OldManagerCuid int     `json:"ld_manager_cuid"`
}

// 用于给Api展示的结构体
type UcenterUsersMini struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Username  string    `json:"username"` // 账号
	Mobile    string    `json:"mobile"`   // 手机号
	Nickname  string    `json:"nickname"` // 昵称
	Email     string    `json:"email"`    // 邮箱
	Avatar    string    `json:"avatar"`   // 头像
	Gender    string    `json:"gender"`   // 性别
	Status    int       `json:"status"`   // 状态 0停用 1启用

	Score         int     `json:"score"`           // 积分
	Money         float64 `json:"money"`           // 余额
	OkMoney       float64 `json:"ok_money"`        // 可提现余额
	NoMoney       float64 `json:"no_money"`        // 不可提现金额
	LastLoginIp   string  `json:"last_login_ip"`   // 最后一次登录ip
	LastLoginTime int     `json:"last_login_time"` // 最后一次登录时间戳
	LastLongitude float64 `json:"last_longitude"`  // 最后一次经度
	LastLatitude  float64 `json:"last_latitude"`   // 最后一次维度

	IsAuth    int    `json:"is_auth"`    // 是否实名认证 0 否 1审核 2通过 -1拒绝
	IdcardTop string `json:"idcard_top"` // 身份证正面
	IdcardBom string `json:"idcard_bom"` // 身份证背面
	IdcardId  string `json:"idcard_id"`  // 身份证号
	ShareOne  int    `json:"share_one"`  // 一级分享
	ShareTwo  int    `json:"share_two"`  // 二级分享
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(UcenterUsers))
}

// AddPost insert a new Post into database and returns
// last inserted Id on success.
func AddUcenterUsers(m *UcenterUsers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPostById retrieves Post by Id. Returns error if
// Id doesn't exist
func GetUcenterUsersById(id int64) (v *UcenterUsers, err error) {
	o := orm.NewOrm()
	v = &UcenterUsers{}

	if err = o.QueryTable(new(UcenterUsers)).Filter("Id", id).Filter("flag", 1).OrderBy("-Id").RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUcenterUsersByOldId(id int64) (v *UcenterUsers, err error) {
	o := orm.NewOrm()
	v = &UcenterUsers{}

	if err = o.QueryTable(new(UcenterUsers)).Filter("OldManagerCuid", id).Filter("flag", 1).OrderBy("-Id").RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUcenterUsersByMobile(mobile string) (v *UcenterUsers, err error) {
	o := orm.NewOrm()
	v = &UcenterUsers{}

	if err = o.QueryTable(new(UcenterUsers)).Filter("Mobile", mobile).Filter("flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllUcenterUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	query["flag"] = "1"
	o := orm.NewOrm()
	qs := o.QueryTable(new(UcenterUsers))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []UcenterUsers
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePost updates Post by Id and returns error if
// the record to be updated doesn't exist
func UpdateUcenterUsersById(m *UcenterUsers) (err error) {
	o := orm.NewOrm()
	v := UcenterUsers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePost deletes Post by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUcenterUsers(id int64) (err error) {
	o := orm.NewOrm()
	v := UcenterUsers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UcenterUsers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
