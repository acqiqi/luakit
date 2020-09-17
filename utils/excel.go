package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

type ExcelUtils struct {
	Title     string                   `json:"title"`      //标题
	Header    Header                   `json:"header"`     //表头
	Keys      []ExcelKey               `json:"keys"`       //导出字段
	DataBase  []map[string]interface{} `json:"data_base"`  //数据源
	SaveType  string                   `json:"save_type"`  //保存类型 File Base64 UploaderUrl
	FilePath  string                   `json:"file_path"`  //文件保存地址
	FileName  string                   `json:"file_name"`  //表格名称
	SheetName string                   `json:"sheet_name"` //Sheet名称
	F         *excelize.File           `json:"f"`          //三方库实体File
}

type Header struct {
	Title  string  `json:"title"`
	Style  string  `json:"style"`  //样式
	Width  float64 `json:"width"`  //宽度
	Height float64 `json:"height"` //高度
}

type ExcelKey struct {
	Name string `json:"name"` //字段名
	Key  string `json:"key"`  //字段key
}

var abc_key = []string{
	"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

// 初始化表格系统
func (this *ExcelUtils) InitExcel() {
	this.F = excelize.NewFile()
	//+strconv.Itoa(tableRow)
}

//导出
func (this *ExcelUtils) Export() {
	this.InitExcel()
	//初始化 Sheet
	index := this.F.NewSheet(this.SheetName)
	this.F.SetActiveSheet(index)
	//创建表头
	keysLen := len(this.Keys)
	//检测Sheet名称
	if this.SheetName == "" {
		this.SheetName = "Sheet1"
	}
	//合并表头和列数量一样
	this.F.MergeCell(this.SheetName, "A1", abc_key[keysLen]+"1")
	this.F.SetCellValue(this.SheetName, "A1", this.Header.Title)
	// 处理样式
	if this.Header.Style != "" {
		style, err := this.F.NewStyle(this.Header.Style)
		if err == nil {
			this.F.SetCellStyle(this.SheetName, "A1", abc_key[keysLen]+"1", style)
		}
	}
	// 处理表头高度
	if this.Header.Height > 0 {
		this.F.SetRowHeight("Sheet1", 1, this.Header.Height)
	}

	//生成数据
	for i, v := range this.DataBase {
		for key, val := range this.Keys {
			//插入
			s := v[val.Key]
		}
	}

}

//func (this *ExcelUtils)  ()  {
//
//}
