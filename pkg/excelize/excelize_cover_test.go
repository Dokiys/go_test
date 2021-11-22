package excelize

import (
	"github.com/xuri/excelize/v2"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

// TestExcelizeCoverCell 测试单元格覆盖
func TestExcelizeCoverCell(t *testing.T) {
	f, err := excelize.OpenFile("../assert/BookCover.xlsx")
	if err != nil {
		t.Fatal(err)
		return
	}

	cell, err := f.GetCellValue("Sheet1", "A1")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(cell)

	err = f.SetCellValue("Sheet1", "A1", "hahahah")
	if err != nil {
		t.Fatal(err)
		return
	}

	f.SaveAs("../assert/BookCover_out.xlsx")
}
