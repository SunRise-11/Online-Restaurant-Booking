package helpers

import (
	"fmt"
	"log"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func CreatePDFReport(RestoName, RestoAddress, dateReport string, nos, nof, noc, total, nor []int) error {

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)
	c.NewPage()
	fontBold, _ := model.NewStandard14Font("Helvetica-Bold")

	// report := PdfResponseFormat(c, "IMAGES/Restobook.png", RestoName, dateReport)

	restoName := c.NewParagraph(fmt.Sprintf("Restaurant Name     : %v", RestoName))
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

	summaryTable := c.NewTable(1)
	summaryTable.SetMargins(35, 0, 30, 0)
	tab1 := c.NewParagraph("Summary")
	tab1.SetFontSize(12)
	cell := summaryTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell.SetContent(tab1)
	c.Draw(summaryTable)

	noaTable := c.NewTable(4)
	noaTable.SetMargins(35, 0, 0, 0)
	tab2 := c.NewParagraph("Number of Accepted")
	tab2.SetFont(fontBold)
	tab2.SetFontSize(11)
	cell2 := noaTable.NewCell()
	cell2.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell2.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell2.SetContent(tab2)

	tab2 = c.NewParagraph("Orders")
	tab2.SetFontSize(11)
	cell2 = noaTable.NewCell()
	cell2.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell2.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell2.SetContent(tab2)

	tab2 = c.NewParagraph("Seats")
	tab2.SetFontSize(11)
	cell2 = noaTable.NewCell()
	cell2.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell2.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell2.SetContent(tab2)

	tab2 = c.NewParagraph("Total")
	tab2.SetFontSize(11)
	cell2 = noaTable.NewCell()
	cell2.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell2.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell2.SetContent(tab2)
	c.Draw(noaTable)

	nosTable := c.NewTable(4)
	nosTable.SetMargins(35, 0, 0, 0)
	tab3 := c.NewParagraph("Number of Success")
	tab3.SetFontSize(11)
	cell3 := nosTable.NewCell()
	cell3.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell3.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell3.SetContent(tab3)

	tab3 = c.NewParagraph(fmt.Sprintf("%v Orders", nos[0]))
	tab3.SetFontSize(11)
	cell3 = nosTable.NewCell()
	cell3.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell3.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell3.SetContent(tab3)

	tab3 = c.NewParagraph(fmt.Sprintf("%v Seats", nos[1]))
	tab3.SetFontSize(11)
	cell3 = nosTable.NewCell()
	cell3.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell3.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell3.SetContent(tab3)

	tab3 = c.NewParagraph(fmt.Sprintf("Rp.%v", nos[2]))
	tab3.SetFontSize(11)
	cell3 = nosTable.NewCell()
	cell3.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell3.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell3.SetContent(tab3)
	c.Draw(nosTable)

	nofTable := c.NewTable(4)
	nofTable.SetMargins(35, 0, 0, 0)
	tab4 := c.NewParagraph("Number of Fail")
	tab4.SetFontSize(11)
	cell4 := nofTable.NewCell()
	cell4.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell4.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell4.SetContent(tab4)

	tab4 = c.NewParagraph(fmt.Sprintf("%v Orders", nof[0]))
	tab4.SetFontSize(11)
	cell4 = nofTable.NewCell()
	cell4.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell4.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell4.SetContent(tab4)

	tab4 = c.NewParagraph(fmt.Sprintf("%v Seats", nof[1]))
	tab4.SetFontSize(11)
	cell4 = nofTable.NewCell()
	cell4.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell4.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell4.SetContent(tab4)

	tab4 = c.NewParagraph(fmt.Sprintf("Rp.%v", nof[2]))
	tab4.SetFontSize(11)
	cell4 = nofTable.NewCell()
	cell4.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell4.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell4.SetContent(tab4)
	c.Draw(nofTable)

	nocTable := c.NewTable(4)
	nocTable.SetMargins(35, 0, 0, 0)
	tab5 := c.NewParagraph("Number of Cancel")
	tab5.SetFontSize(11)
	cell5 := nocTable.NewCell()
	cell5.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell5.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell5.SetContent(tab5)

	tab5 = c.NewParagraph(fmt.Sprintf("%v Orders", noc[0]))
	tab5.SetFontSize(11)
	cell5 = nocTable.NewCell()
	cell5.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell5.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell5.SetContent(tab5)

	tab5 = c.NewParagraph(fmt.Sprintf("%v Seats", noc[1]))
	tab5.SetFontSize(11)
	cell5 = nocTable.NewCell()
	cell5.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell5.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell5.SetContent(tab5)

	tab5 = c.NewParagraph(fmt.Sprintf("Rp.%v", noc[2]))
	tab5.SetFontSize(12)
	cell5 = nocTable.NewCell()
	cell5.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell5.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell5.SetContent(tab5)
	c.Draw(nocTable)

	totalTable := c.NewTable(4)
	totalTable.SetMargins(35, 0, 0, 0)
	tab6 := c.NewParagraph("Total")
	tab6.SetFontSize(11)
	cell6 := totalTable.NewCell()
	cell6.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell6.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell6.SetContent(tab6)

	tab6 = c.NewParagraph(fmt.Sprintf("%v Orders", total[0]))
	tab6.SetFontSize(11)
	cell6 = totalTable.NewCell()
	cell6.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell6.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell6.SetContent(tab6)

	tab6 = c.NewParagraph(fmt.Sprintf("%v Seats", total[1]))
	tab6.SetFontSize(11)
	cell6 = totalTable.NewCell()
	cell6.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell6.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell6.SetContent(tab6)

	tab6 = c.NewParagraph(fmt.Sprintf("Rp.%v", total[2]))
	tab6.SetFontSize(11)
	cell6 = totalTable.NewCell()
	cell6.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell6.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell6.SetContent(tab6)
	c.Draw(totalTable)

	norTable := c.NewTable(4)
	norTable.SetMargins(35, 0, 30, 0)
	tab7 := c.NewParagraph("Number of Rejected")
	tab7.SetFont(fontBold)
	tab7.SetFontSize(11)
	cell7 := norTable.NewCell()
	cell7.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell7.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell7.SetContent(tab7)

	tab7 = c.NewParagraph(fmt.Sprintf("%v Orders", nor[0]))
	tab7.SetFontSize(11)
	cell7 = norTable.NewCell()
	cell7.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell7.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell7.SetContent(tab7)

	tab7 = c.NewParagraph(fmt.Sprintf("%v Seats", nor[1]))
	tab7.SetFontSize(11)
	cell7 = norTable.NewCell()
	cell7.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell7.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
	cell7.SetContent(tab7)

	c.Draw(norTable)

	// Write to output file.
	if err := c.WriteToFile(fmt.Sprintf("./EXPORTPDF/%v.pdf", RestoName)); err != nil {
		log.Fatal(err)
	}

	return nil
}
