package excelize

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strconv"
	"testing"
)
// TestExcelizeDup 测试插入复制行
func TestExcelizeDup(t *testing.T) {
	f, err := excelize.OpenFile("BookDup.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.DuplicateRowTo("Sheet1", 1, 2)
	if err != nil {
		return
	}

	f.SaveAs("BookDup_out.xlsx")
}

// TestExcelizeCopy 测试替换复制行
func TestExcelizeCopy(t *testing.T) {
	f, _ := excelize.OpenFile("BookCopy.xlsx")

	data := []int{1, 2, 3}
	// 校验指定单元格高度是否为数据条数的倍数
	cells := []string{"A1", "C9"}
	a1, err := strconv.Atoi(string(cells[0][1]))
	if err != nil {
		t.Error(err)
	}
	a2, err := strconv.Atoi(string(cells[1][1]))
	if err != nil {
		t.Error(err)
	}
	if (a2-a1+1)%len(data) != 0 {
		t.Error("渲染数据与行数不匹配")
	}
	height := 2
	start := 0
	sheet := "Sheet1"
	// 复制块数据
	for i := 1; i < len(data); i++ {
		for j := 0; j < height; j++ {
			err := f.DuplicateRowTo(sheet, start+j+1, start+height+j+height*(i-1)+1)
			if err != nil {
				t.Error(err)

			}
			err = f.RemoveRow(sheet, start+height+j+height*(i-1)+2)
			if err != nil {
				t.Error(err)

			}
		}
	}
	f.SaveAs("BookCopy_out.xlsx")
}