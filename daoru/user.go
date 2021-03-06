package daoru

type ClientUsers struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	Nickname         string `json:"nickname"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Mobile           string `json:"mobile"`
	Avatar           string `json:"avatar"`
	Money            string `json:"money"`
	Score            string `json:"score"`
	SignupIP         string `json:"signup_ip"`
	LastLoginTime    string `json:"last_login_time"`
	LastLoginIP      string `json:"last_login_ip"`
	Sort             string `json:"sort"`
	Status           string `json:"status"`
	InfoBind         string `json:"info_bind"`
	Gender           string `json:"gender"`
	UserKey          string `json:"user_key"`
	SmallOpenid      string `json:"small_openid"`
	RegFlag          string `json:"reg_flag"`
	RegTime          string `json:"reg_time"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
	Flag             string `json:"flag"`
	IsEntAuth        string `json:"is_ent_auth"`
	EntName          string `json:"ent_name"`
	EntCode          string `json:"ent_code"`
	EntPic           string `json:"ent_pic"`
	EntLogo          string `json:"ent_logo"`
	EntMobile        string `json:"ent_mobile"`
	EntAddr          string `json:"ent_addr"`
	Lv               string `json:"lv"`
	Pingfen          string `json:"pingfen"`
	Caogao           string `json:"caogao"`
	Sgzl             string `json:"sgzl"`
	Sgjd             string `json:"sgjd"`
	Fwtd             string `json:"fwtd"`
	Gdgl             string `json:"gdgl"`
	ShareLv1         string `json:"share_lv1"`
	ShareLv2         string `json:"share_lv2"`
	Shouyi           string `json:"shouyi"`
	ShareShouyi      string `json:"share_shouyi"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	MyLatitude       string `json:"my_latitude"`
	MyLongitude      string `json:"my_longitude"`
	TplAddrName      string `json:"tpl_addr_name"`
	TplAddrLatitude  string `json:"tpl_addr_latitude"`
	TplAddrLongitude string `json:"tpl_addr_longitude"`
	TplAddr          string `json:"tpl_addr"`
}
