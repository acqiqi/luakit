package common

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"luakit/models"
	"luakit/utils"
	"strconv"
)

type ExcelUtils struct {
	FileName       string           `json:"file_name"`
	QuotationTable []QuotationTable `json:"quotation_table"`
}

type QuotationTable struct {
	Content  string  `json:"content"`
	Unit     string  `json:"unit"`
	Num      float64 `json:"num"`
	Zl       float64 `json:"zl"`
	Fl       float64 `json:"fl"`
	Work     float64 `json:"work"`
	Total    float64 `json:"total"`
	Describe string  `json:"describe"`
}

func (this *ExcelUtils) Test() error {
	log.Println("????")
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// 创建表头
	f.MergeCell("Sheet1", "A1", "K1")
	f.SetCellValue("Sheet1", "A1", "工程报价单")
	f.SetRowHeight("Sheet1", 1, 20)
	style, err := f.NewStyle(`{"font":{"family":"Berlin Sans FB Demi","size":12,"color":"#333333"}}`)
	if err != nil {
		return err
	}
	f.SetCellStyle("Sheet1", "A1", "K1", style)
	// header
	f.SetCellValue("Sheet1", "A2", "序号")
	f.SetCellValue("Sheet1", "B2", "工作内容")
	f.SetColWidth("Sheet1", "B", "B", 20)
	f.SetCellValue("Sheet1", "C2", "单位")
	f.SetCellValue("Sheet1", "D2", "数量")
	f.SetCellValue("Sheet1", "E2", "主料单价")
	f.SetCellValue("Sheet1", "F2", "辅料单价")
	f.SetCellValue("Sheet1", "G2", "用工单价")
	f.SetCellValue("Sheet1", "H2", "综合单价")
	f.SetCellValue("Sheet1", "I2", "合价")
	f.SetCellValue("Sheet1", "J2", "区域")
	f.SetCellValue("Sheet1", "K2", "备注")
	f.SetColWidth("Sheet1", "K", "K", 50)
	f.SetRowHeight("Sheet1", 2, 20)
	style, err = f.NewStyle(`{"font":{"family":"Berlin Sans FB Demi","size":12,"color":"#333333"},"alignment":{"horizontal":"center"}}`)
	if err != nil {
		return err
	}
	f.SetCellStyle("Sheet1", "A2", "K2", style)

	//测试业务
	q, err := models.GetSmQuotationById(2)
	if err != nil {
		return errors.New("报价单不存在")
	}
	q_count, q_list, err := models.GetSmQuotationLinkListByQId(2)
	if err != nil || q_count == 0 {
		return errors.New("没有项目")
	}

	keyList := []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十"}

	for i, v := range *q_list {
		log.Println(i, v)
		//创建子表头
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(3+i), keyList[i])
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(3+i), v.CatName)
		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(3+i), "K"+strconv.Itoa(3+i), style)
		table_list := []QuotationTable{}
		err := utils.JsonDecode(v.TableJson, &table_list)
		if err != nil {
			return errors.New("Table数据错误")
		}
		for ti, tv := range table_list {
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(4+ti), ti+1)
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(4+ti), tv.Content)
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(4+ti), tv.Unit)
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(4+ti), tv.Num)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(4+ti), tv.Zl)
			f.SetCellValue("Sheet1", "F"+strconv.Itoa(4+ti), tv.Fl)
			f.SetCellValue("Sheet1", "G"+strconv.Itoa(4+ti), tv.Work)
			f.SetCellValue("Sheet1", "H"+strconv.Itoa(4+ti), tv.Zl+tv.Fl+tv.Work)
			f.SetCellValue("Sheet1", "I"+strconv.Itoa(4+ti), tv.Total)
			f.SetCellValue("Sheet1", "J"+strconv.Itoa(4+ti), "区域")
			f.SetCellValue("Sheet1", "K"+strconv.Itoa(4+ti), tv.Describe)
			f.SetCellStyle("Sheet1", "A"+strconv.Itoa(4+ti), "K"+strconv.Itoa(4+ti), style)
		}

	}

	//string := strconv.Itoa(int)

	log.Println(q)

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("./Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	log.Print("ojbk")
	return nil
}
