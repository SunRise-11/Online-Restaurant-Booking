package helpers

import (
	"fmt"
	"log"

	"github.com/unidoc/unipdf/v3/creator"
)

func CreatePDFReport(RestoName, RestoAddress, dateReport string) error {

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)
	c.NewPage()

	// report := PdfResponseFormat(c, "IMAGES/Restobook.png", RestoName, dateReport)

	restoName := c.NewParagraph(fmt.Sprintf("Restaurant Name : %v", RestoName))
	restoName.SetFontSize(14)
	restoName.SetMargins(35, 0, 50, 0)
	restoName.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(restoName)

	restoAddress := c.NewParagraph(fmt.Sprintf("Restaurant Address : %v", RestoAddress))
	restoAddress.SetFontSize(14)
	restoAddress.SetMargins(35, 0, 10, 0)
	restoAddress.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(restoAddress)

	restoDate := c.NewParagraph(fmt.Sprintf("Date \t:\t %v", dateReport))
	restoDate.SetFontSize(14)
	restoDate.SetMargins(35, 0, 10, 0)
	restoDate.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(restoDate)

	restoTable := c.NewChapter("")
	restoTableIn := restoTable.NewSubchapter("")

	summaryTable := c.NewTable(1)
	summaryTable.SetMargins(50, 0, 30, 0)

	tab1 := c.NewParagraph("Summary")
	tab1.SetFontSize(14)
	cell := summaryTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell.SetContent(tab1)

	restoTableIn.Add(summaryTable)
	c.Draw(restoTable)

	// Write to output file.
	if err := c.WriteToFile(fmt.Sprintf("%v.pdf", RestoName)); err != nil {
		log.Fatal(err)
	}

	return nil
}

// func PdfResponseFormat(c *creator.Creator, logoPath, RestoName, dateReport string) *creator.Paragraph {

// 	report := c.NewParagraph(RestoName)
// 	report.SetFontSize(30)
// 	report.SetMargins(85, 0, 150, 0)
// 	report.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
// 	c.Draw(report)

// 	report = c.NewParagraph(dateReport)
// 	report.SetFontSize(20)
// 	report.SetMargins(80, 0, 0, 0)
// 	report.SetColor(creator.ColorBlue)
// 	c.Draw(report)

// 	return report

// }
