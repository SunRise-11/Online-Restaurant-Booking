package helpers

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func CreateExcelReport(RestoName, RestoAddress, dateReport string, nos, nof, noc, total, nor []int) error {
	styles := [][]interface{}{
		{"A6", "D6", 1, 1, 1, 1},
		{"A8", "A11", 1, 1, 1, 1},
		{"B7", "B11", 1, 1, 1, 1},
		{"C7", "C11", 1, 1, 1, 1},
		{"D7", "D11", 1, 1, 1, 1},
		{"B13", "C13", 1, 1, 1, 1},
	}

	f := excelize.NewFile()

	for _, style := range styles {
		styleBold, _ := f.NewStyle(`{"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1},{"type":"right","color":"#000000","style":1}],"font":{"bold":true},"alignment":{"horizontal":"center"}}`)
		f.SetColWidth("Sheet1", "A", "A", 20)
		f.SetColWidth("Sheet1", "B", "B", 15)
		f.SetColWidth("Sheet1", "C", "C", 15)
		f.SetColWidth("Sheet1", "D", "D", 15)
		f.SetCellValue("Sheet1", "A1", "Restaurant Name")
		f.SetCellValue("Sheet1", "B1", fmt.Sprint(": ", RestoName))
		f.SetCellValue("Sheet1", "A2", "Restaurant Address")
		f.SetCellValue("Sheet1", "B2", fmt.Sprint(": ", RestoAddress))
		f.SetCellValue("Sheet1", "A3", "Date")
		f.SetCellValue("Sheet1", "B3", fmt.Sprint(": ", dateReport))
		f.SetCellValue("Sheet1", "A6", "Summary")
		f.MergeCell("Sheet1", "A6", "D6")
		f.SetCellStyle("Sheet1", "A7", "A7", styleBold)
		f.SetCellValue("Sheet1", "A7", "Number Of Accepted")
		f.SetCellValue("Sheet1", "B7", "Orders")
		f.SetCellValue("Sheet1", "C7", "Seats")
		f.SetCellValue("Sheet1", "D7", "Total")
		f.SetCellValue("Sheet1", "A8", "Number Of Success")
		f.SetCellValue("Sheet1", "B8", fmt.Sprint(nos[0], " Orders"))
		f.SetCellValue("Sheet1", "C8", fmt.Sprint(nos[1], " Seats"))
		f.SetCellValue("Sheet1", "D8", fmt.Sprint("Rp.", nos[2]))
		f.SetCellValue("Sheet1", "A9", "Number Of Fail")
		f.SetCellValue("Sheet1", "B9", fmt.Sprint(nof[0], " Orders"))
		f.SetCellValue("Sheet1", "C9", fmt.Sprint(nof[1], " Seats"))
		f.SetCellValue("Sheet1", "D9", fmt.Sprint("Rp.", nof[2]))
		f.SetCellValue("Sheet1", "A10", "Number Of Cancel")
		f.SetCellValue("Sheet1", "B10", fmt.Sprint(noc[0], " Orders"))
		f.SetCellValue("Sheet1", "C10", fmt.Sprint(noc[1], " Seats"))
		f.SetCellValue("Sheet1", "D10", fmt.Sprint("Rp.", noc[2]))
		f.SetCellValue("Sheet1", "A11", "Total")
		f.SetCellValue("Sheet1", "B11", fmt.Sprint(total[0], " Orders"))
		f.SetCellValue("Sheet1", "C11", fmt.Sprint(total[1], " Seats"))
		f.SetCellValue("Sheet1", "D11", fmt.Sprint("Rp.", total[2]))
		f.SetCellValue("Sheet1", "A13", "Number Of Rejected")
		f.SetCellValue("Sheet1", "B13", fmt.Sprint(nor[0], " Orders"))
		f.SetCellValue("Sheet1", "C13", fmt.Sprint(nor[1], " Seats"))
		f.SetCellStyle("Sheet1", "A13", "A13", styleBold)
		s, err := f.NewStyle(&excelize.Style{
			Border: []excelize.Border{
				{
					Type:  "left",
					Color: "#000000",
					Style: style[2].(int),
				}, {
					Type:  "top",
					Color: "#000000",
					Style: style[3].(int),
				}, {
					Type:  "bottom",
					Color: "#000000",
					Style: style[4].(int),
				}, {
					Type:  "right",
					Color: "#000000",
					Style: style[5].(int),
				},
			},
			Alignment: &excelize.Alignment{
				Horizontal: "center",
			},
		})
		if err != nil {
			fmt.Println(err)
		}
		f.SetCellStyle("Sheet1", style[0].(string), style[1].(string), s)
	}

	if err := f.SaveAs(fmt.Sprintf("./EXPORT/EXCEL/%v.xlsx", RestoName)); err != nil {
		return err
	}
	return nil
}
