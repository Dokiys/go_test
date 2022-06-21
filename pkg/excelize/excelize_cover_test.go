package excelize1

import (
	"github.com/xuri/excelize/v2"
	"go_test/pkg"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

// TestExcelizeCoverCell 测试单元格覆盖
func TestExcelizeCoverCell(t *testing.T) {
	f, err := excelize.OpenFile("BookCover.xlsx")
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

	f.SaveAs(pkg.PathPrefix + "BookCover_out.xlsx")
}
