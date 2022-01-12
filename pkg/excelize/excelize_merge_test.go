package excelize1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"testing"
)
// TestExcelizeDup 测试插入复制行
func TestMergeCellRectRow(t *testing.T) {
	f := excelize.NewFile()
	sheet := "Sheet1"

	assert.NoError(t, f.DuplicateRowTo(sheet, 4, 5))
	f.MergeCell(sheet,"A4","A5")

	assert.NoError(t, f.DuplicateRowTo(sheet, 2, 3))
	assert.NoError(t, f.DuplicateRowTo(sheet, 2, 4))
	f.MergeCell(sheet,"A2","A3")
	f.MergeCell(sheet,"A3","A4")

	f.SaveAs(fmt.Sprintf("./1.xlsx"))
}

func TestMergeCellRectCol(t *testing.T) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "B1", 1)
	f.SetCellValue(sheet, "C1", 0)
	f.SetCellValue(sheet, "D1", 1)

	f.InsertCol(sheet, "E")
	f.MergeCell(sheet,"D1","E1")

	f.InsertCol(sheet, "C")
	f.InsertCol(sheet, "D")
	f.MergeCell(sheet,"B1","C1")
	f.MergeCell(sheet,"C1","D1")

	f.SaveAs(fmt.Sprintf("./2.xlsx"))
}

func TestMergeCellDup(t *testing.T) {
	// Test merge cell after duplicate null row
	{
		f := excelize.NewFile()
		sheet := "Sheet1"
		assert.NoError(t, f.MergeCell(sheet,"A4","A5"))

		assert.NoError(t, f.DuplicateRowTo(sheet, 2, 3))
		assert.NoError(t, f.DuplicateRowTo(sheet, 2, 4))
		assert.NoError(t, f.MergeCell(sheet,"A2","A3"))
		assert.NoError(t, f.MergeCell(sheet,"A3","A4"))

		mergeCells, err := f.GetMergeCells(sheet)
		assert.NoError(t, err)
		assert.ElementsMatch(t, mergeCells, []excelize.MergeCell{{"A6:A7",""},{"A2:A4",""}})
	}

	// Test merge cell after duplicate null col
	{

		f := excelize.NewFile()
		sheet := "Sheet1"

		assert.NoError(t, f.MergeCell(sheet,"D1","E1"))

		assert.NoError(t, f.InsertCol(sheet, "C"))
		assert.NoError(t, f.InsertCol(sheet, "D"))
		assert.NoError(t, f.MergeCell(sheet,"B1","C1"))
		assert.NoError(t, f.MergeCell(sheet,"C1","D1"))

		mergeCells, err := f.GetMergeCells(sheet)
		assert.NoError(t, err)
		assert.ElementsMatch(t, mergeCells, []excelize.MergeCell{{"B1:D1",""},{"F1:G1",""}})
	}
}