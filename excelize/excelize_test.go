package excelize

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

// TestExcelizeHello 测试Excelize
func TestExcelizeHello(t *testing.T) {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet2")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello work!")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("BookHello_out.xlsx"); err != nil {
		fmt.Println(err)
	}
}

// TestExcelizeReadExcel 测试读取文件
func TestExcelizeReadExcel(t *testing.T) {
	f, err := excelize.OpenFile("BookRead.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	f.SaveAs("BookRead_out.xlsx")
}

// TestExcelizeChart 添加图表
func TestExcelizeChart(t *testing.T) {
	f := excelize.NewFile()
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}

	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet1!$A$2",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$2:$D$2"
        },
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return
	}
	// 根据指定路径保存文件
	if err := f.SaveAs("BookChart_out.xlsx"); err != nil {
		fmt.Println(err)
	}
}

// TestExcelizeImage 插入图片
func TestExcelizeImage(t *testing.T) {
	f := excelize.NewFile()
	// 插入图片
	if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的缩放比例
	if err := f.AddPicture("Sheet1", "E2", "image.jpeg", `{
        "x_scale": 0.5,
        "y_scale": 0.5
    }`); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的打印属性
	if err := f.AddPicture("Sheet1", "G2", "image.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
		fmt.Println(err)
	}
	// 保存文件
	if err := f.SaveAs("BookImage_out.xlsx"); err != nil {
		fmt.Println(err)
	}
}
