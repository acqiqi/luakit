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
	Content   string  `json:"content"`
	Unit      string  `json:"unit"`
	Num       float64 `json:"num"`
	Zl        float64 `json:"zl"`
	Fl        float64 `json:"fl"`
	Work      float64 `json:"work"`
	UnitTotal float64 `json:"unit_total"`
	Total     float64 `json:"total"`
	Describe  string  `json:"describe"`
}

func (this *ExcelUtils) Test(id int64) (err error, quotation_no string) {
	//测试业务
	q, err := models.GetSmQuotationById(id)
	if err != nil {
		return errors.New("报价单不存在"), ""
	}
	q_count, q_list, err := models.GetSmQuotationLinkListByQId(id)
	if err != nil || q_count == 0 {
		return errors.New("没有项目"), ""
	}

	log.Println("????")
	f := excelize.NewFile()

	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	// 创建表头
	f.MergeCell("Sheet1", "A1", "D1")
	f.MergeCell("Sheet1", "E1", "J1")
	f.SetCellValue("Sheet1", "A1", q.Name)
	f.SetCellValue("Sheet1", "F1", "唯一编码："+q.QuotationNo)

	f.SetRowHeight("Sheet1", 1, 30)
	style, err := f.NewStyle(`{"font":{"family":"Berlin Sans FB Demi","size":12,"color":"#333333"}}`)
	if err != nil {
		return err, ""
	}
	f.SetCellStyle("Sheet1", "A1", "J1", style)
	// header
	f.SetCellValue("Sheet1", "A2", "序号")
	f.SetColWidth("Sheet1", "A", "A", 5)
	f.SetCellValue("Sheet1", "B2", "工作内容")
	f.SetColWidth("Sheet1", "B", "B", 20)
	f.SetCellValue("Sheet1", "C2", "单位")
	f.SetColWidth("Sheet1", "C", "C", 5)
	f.SetCellValue("Sheet1", "D2", "数量")
	f.SetColWidth("Sheet1", "D", "D", 6)
	f.SetCellValue("Sheet1", "E2", "主料单价")
	f.SetColWidth("Sheet1", "E", "E", 10)
	f.SetCellValue("Sheet1", "F2", "辅料单价")
	f.SetColWidth("Sheet1", "F", "F", 10)
	f.SetCellValue("Sheet1", "G2", "用工单价")
	f.SetColWidth("Sheet1", "G", "G", 10)
	f.SetCellValue("Sheet1", "H2", "综合单价")
	f.SetCellValue("Sheet1", "I2", "合价")
	f.SetCellValue("Sheet1", "J2", "备注")
	//f.SetCellValue("Sheet1", "K2", "备注")
	f.SetColWidth("Sheet1", "J", "J", 20)
	f.SetRowHeight("Sheet1", 2, 30)
	style, err = f.NewStyle(`{"font":{"family":"Berlin Sans FB Demi","size":12,"color":"#333333"},"alignment":{"horizontal":"center"}}`)
	if err != nil {
		return err, ""
	}
	f.SetCellStyle("Sheet1", "A2", "J2", style)

	keyList := []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十"}

	tableRow := 2
	totalAll := 0.00
	for i, v := range *q_list {
		totalLink := 0.00
		tableRow++
		log.Println(i, v)
		//创建子表头
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(tableRow), keyList[i])
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), v.CatName)
		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow), style)
		f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow))
		table_list := []QuotationTable{}
		log.Println(v.TableJson)
		log.Println("-")
		err := utils.JsonDecode(v.TableJson, &table_list)
		if err != nil {
			return errors.New("Table数据错误"), ""
		}
		for ti, tv := range table_list {
			tableRow++
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(tableRow), ti+1)
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), tv.Content)
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(tableRow), tv.Unit)
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(tableRow), tv.Num)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(tableRow), tv.Zl)
			f.SetCellValue("Sheet1", "F"+strconv.Itoa(tableRow), tv.Fl)
			f.SetCellValue("Sheet1", "G"+strconv.Itoa(tableRow), tv.Work)
			f.SetCellValue("Sheet1", "H"+strconv.Itoa(tableRow), tv.UnitTotal)
			f.SetCellValue("Sheet1", "I"+strconv.Itoa(tableRow), tv.Total)
			f.SetCellValue("Sheet1", "J"+strconv.Itoa(tableRow), tv.Describe)
			//f.SetCellValue("Sheet1", "K"+strconv.Itoa(tableRow), tv.Describe)
			f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow), style)
			totalLink += tv.Total
		}
		//插入合计
		tableRow++
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), "小计")
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(tableRow), totalLink)
		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(3+i), style)
		f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "H"+strconv.Itoa(tableRow))
		totalAll += totalLink
	}

	//其他
	tableRow++

	//创建子表头
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(tableRow), keyList[len(*q_list)])
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), "其他")
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow), style)
	f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow))
	//创建管理费
	tableRow++
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(tableRow), 1)
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), "管理费")

	f.SetCellValue("Sheet1", "I"+strconv.Itoa(tableRow), utils.Decimal(totalAll*(float64(q.ManagerBl)/100)))
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(tableRow), "总价的"+strconv.Itoa(q.ManagerBl)+"%")
	//f.SetCellValue("Sheet1", "K"+strconv.Itoa(tableRow), "")
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow), style)
	f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "H"+strconv.Itoa(tableRow))
	//创建税费
	tableRow++
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(tableRow), 2)
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), "税费")

	f.SetCellValue("Sheet1", "I"+strconv.Itoa(tableRow), utils.Decimal(totalAll*(float64(q.TaxBl)/100)))
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(tableRow), "总价的"+strconv.Itoa(q.TaxBl)+"%")
	f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "H"+strconv.Itoa(tableRow))
	//f.SetCellValue("Sheet1", "K"+strconv.Itoa(tableRow), "")
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "H"+strconv.Itoa(tableRow), style)

	tableRow++
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), "总价")
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(tableRow), totalAll+utils.Decimal(totalAll*(float64(q.ManagerBl)/100))+utils.Decimal(totalAll*(float64(q.TaxBl)/100)))
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow), style)
	f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "H"+strconv.Itoa(tableRow))

	tableRow++
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), "优惠")
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow), style)
	f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "H"+strconv.Itoa(tableRow))

	tableRow++
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(tableRow), "一口价")
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(tableRow), q.FixedPrice)
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(tableRow), "J"+strconv.Itoa(tableRow), style)
	f.MergeCell("Sheet1", "C"+strconv.Itoa(tableRow), "H"+strconv.Itoa(tableRow))

	log.Println(totalAll)
	// 绘制边框

	b_style, err := f.NewStyle(`{"border":[{"type":"left","color":"333333","style":1},{"type":"top","color":"333333","style":1},{"type":"bottom","color":"333333","style":1},{"type":"right","color":"333333","style":1}],"font":{"family":"Berlin Sans FB Demi","size":12,"color":"#333333"},"alignment":{"horizontal":"center","vertical":"center"}}`)
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(1), "J"+strconv.Itoa(tableRow), b_style)
	b_style, err = f.NewStyle(`{"border":[{"type":"left","color":"333333","style":1},{"type":"top","color":"333333","style":1},{"type":"bottom","color":"333333","style":1},{"type":"right","color":"333333","style":1}],"font":{"family":"Berlin Sans FB Demi","size":12,"color":"#333333"},"alignment":{"horizontal":"left","vertical":"center"}}`)
	f.SetCellStyle("Sheet1", "A1", "J1", b_style)

	//设定高度

	for hi := 0; hi < tableRow; hi++ {
		f.SetRowHeight("Sheet1", hi+1, 30)
	}

	//string := strconv.Itoa(int)

	log.Println(q)

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("./excelout/" + q.QuotationNo + ".xlsx"); err != nil {
		fmt.Println(err)
		fmt.Println("?????????????")
	}
	log.Print("ojbk")
	return nil, q.QuotationNo
}
