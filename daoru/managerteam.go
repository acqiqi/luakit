package daoru

type ManagerTeam struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Uname        string `json:"uname"`
	UcardID      string `json:"ucard_id"`
	UcardPic     string `json:"ucard_pic"`
	UcardPicBack string `json:"ucard_pic_back"`
	EntRen       string `json:"ent_ren"`
	EntMobile    string `json:"ent_mobile"`
	EntEmail     string `json:"ent_email"`
	EntAddr      string `json:"ent_addr"`
	EntQq        string `json:"ent_qq"`
	GrNum        string `json:"gr_num"`
	Sc           string `json:"sc"`
	Gz           string `json:"gz"`
	Cjnl         string `json:"cjnl"`
	Dznl         string `json:"dznl"`
	SgCity       string `json:"sg_city"`
	About        string `json:"about"`
	Gwal         string `json:"gwal"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	Flag         string `json:"flag"`
	Status       string `json:"status"`
	Muid         string `json:"muid"`
	EntLogo      string `json:"ent_logo"`
}
