package excelize

import (
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"testing"
)

type Alignment struct {
	Horizontal      string `json:"horizontal"`
	Indent          int    `json:"indent"`
	JustifyLastLine bool   `json:"justify_last_line"`
	ReadingOrder    uint64 `json:"reading_order"`
	RelativeIndent  int    `json:"relative_indent"`
	ShrinkToFit     bool   `json:"shrink_to_fit"`
	TextRotation    int    `json:"text_rotation"`
	Vertical        string `json:"vertical"`
	WrapText        bool   `json:"wrap_text"`
}

type Border struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}

type Font struct {
	Bold      bool    `json:"bold"`
	Italic    bool    `json:"italic"`
	Underline string  `json:"underline"`
	Family    string  `json:"family"`
	Size      float64 `json:"size"`
	Strike    bool    `json:"strike"`
	Color     string  `json:"color"`
}

type Fill struct {
	Type    string   `json:"type"`
	Pattern int      `json:"pattern"`
	Color   []string `json:"color"`
	Shading int      `json:"shading"`
}

type Protection struct {
	Hidden bool `json:"hidden"`
	Locked bool `json:"locked"`
}

type Style struct {
	Border        []Border    `json:"border"`
	Fill          Fill        `json:"fill"`
	Font          *Font       `json:"font"`
	Alignment     *Alignment  `json:"alignment"`
	Protection    *Protection `json:"protection"`
	NumFmt        int         `json:"number_format"`
	DecimalPlaces int         `json:"decimal_places"`
	CustomNumFmt  *string     `json:"custom_number_format"`
	Lang          string      `json:"lang"`
	NegRed        bool        `json:"negred"`
}

// TestExcelizeStyle 测试excelize样式
func TestExcelizeStyle(t *testing.T) {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	//style, _ := f.NewStyle(`{"border":[{"type":"left","color":"0000FF","style":3},{"type":"top","color":"00FF00","style":4},{"type":"bottom","color":"FFFF00","style":5},{"type":"right","color":"FF0000","style":6},{"type":"diagonalDown","color":"A020F0","style":7},{"type":"diagonalUp","color":"A020F0","style":8}]}`)
	//style, _ := f.NewStyle(`{"border":[{"type":"","color":"","style":0}],"fill":{"type":"pattern","pattern":1,"color":["#FF0000"],"shading":0},"font":{"bold":false,"italic":false,"underline":"","family":"Comic Sans MS","size":0,"strike":false,"color":""},"alignment":{"horizontal":"","indent":0,"justify_last_line":false,"reading_order":0,"relative_indent":0,"shrink_to_fit":false,"text_rotation":0,"vertical":"top","wrap_text":false},"protection":{"hidden":false,"locked":false},"number_format":0,"decimal_places":0,"custom_number_format":"","lang":"","negred":false}`)
	style, _ := f.NewStyle(&excelize.Style{Border: []excelize.Border{
		{
			Type:  "top",
			Color: "000000",
			Style: 5,
		},
	}})
	err := f.MergeCell("Sheet1", "B2", "C4")
	if err != nil {
		t.Fatal(err)
	}
	_ = f.SetCellStyle("Sheet1", "B2", "B2", style)

	// 设置单元格的值
	f.SetCellValue("Sheet1", "C2", "lalala")
	f.SetCellValue("Sheet1", "D4", "lalala")
	//f.SetColStyle()	// 当前版本v2.4.1不支持

	f.SaveAs("../assert/BookStyle_out.xlsx")
}

// TestExcelizeStyleJsonStruct 默认值生成excel
func TestExcelizeStyleJsonStruct(t *testing.T) {
	f := excelize.NewFile()
	str := "[$-380A]dddd\\\\,\\\\ dd\\\" de \\\"mmmm\\\" de \\\"yyyy;@"

	style := Style{
		Border: []Border{{
			Type:  "",
			Color: "",
			Style: 0,
		}},
		Fill: Fill{
			Type:    "",
			Pattern: 0,
			Color:   []string{},
			Shading: 0,
		},
		Font: &Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "",
			Size:      0,
			Strike:    false,
			Color:     "",
		},
		Alignment: &Alignment{
			Horizontal:      "",
			Indent:          0,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "",
			WrapText:        false,
		},
		Protection: &Protection{
			Hidden: false,
			Locked: false,
		},
		NumFmt:        0,
		DecimalPlaces: 0,
		CustomNumFmt:  &str,
		Lang:          "",
		NegRed:        false,
	}

	j, _ := json.Marshal(style)
	t.Logf("style json: %s", string(j))

	styleIndex, _ := f.NewStyle(string(j))
	sheet := "Sheet1"
	f.MergeCell(sheet, "B2", "C4")
	f.SetCellStyle(sheet, "B2", "B2", styleIndex)

	// 设置单元格的值
	f.SetCellValue(sheet, "C2", "lalala")
	f.SetCellValue(sheet, "D4", "lalala")
	//f.SetColStyle()	// 当前版本v2.4.1不支持

	f.SaveAs("../assert/BookDefaultStyle_out.xlsx")
}

// TestExcelizeMergeCellStyle 测试设置合并单元格样式
func TestExcelizeMergeCellStyle(t *testing.T) {
	f := excelize.NewFile()

	styleIndex, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#FF0000"},
		Shading: 0,
	}})
	f.MergeCell("Sheet1", "A1", "C4")
	f.SetCellStyle("Sheet1", "A1", "A1", styleIndex)

	f.SaveAs("../assert/BookMergeCellStyle_out.xlsx")
}

// TestExcelizeRepeatedStyle 测试New重复的Style index是否复用
func TestExcelizeRepeatedStyle(t *testing.T) {
	f := excelize.NewFile()

	styleIndex1, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#FF0000"},
		Shading: 0,
	}})
	styleIndex2, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#FF0000"},
		Shading: 0,
	}})

	t.Logf("styleIndex1: %v", styleIndex1)
	t.Logf("styleIndex2: %v", styleIndex2)
}

// TestExcelizeBorderStyle 测试边框样式实际展示效果
func TestExcelizeBorderStyle(t *testing.T) {
	f := excelize.NewFile()

	styleIndex, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#FF00FF", Style: 5},
			{Type: "bottom", Color: "#FF00FF", Style: 2},
			{Type: "left", Color: "#FF00FF", Style: 3},
			{Type: "right", Color: "#FF00FF", Style: 4},
		},
	})
	f.MergeCell("Sheet1", "B2", "C4")
	f.SetCellStyle("Sheet1", "B2", "C4", styleIndex)
	f.SetCellValue("Sheet1", "B2", "B2")

	f.SaveAs("../assert/BookBorderStyle_out.xlsx")
}

// TestExcelizeNilStyle 测试空style
func TestExcelizeNilStyle(t *testing.T) {
	f := excelize.NewFile()

	style := Style{
		Border: []Border{
			{Type: "top", Color: "#FF00FF", Style: 5},
			{Type: "bottom", Color: "#FF00FF", Style: 2},
			{Type: "left", Color: "#FF00FF", Style: 3},
			{Type: "right", Color: "#FF00FF", Style: 4},
		},
		Fill: Fill{
			Type:    "",
			Pattern: 0,
			Color:   []string{},
			Shading: 0,
		},
		Font: &Font{
			Bold:      false,
			Italic:    false,
			Underline: "",
			Family:    "",
			Size:      0,
			Strike:    false,
			Color:     "",
		},
		Alignment: &Alignment{
			Horizontal:      "",
			Vertical:        "",
		},
	}

	j, _ := json.Marshal(style)

	styleIndex, _ := f.NewStyle(string(j))
	sheet := "Sheet1"
	f.MergeCell(sheet, "B2", "C4")
	f.SetCellStyle(sheet, "B2", "C4", styleIndex)

	// 设置单元格的值
	f.SetCellValue(sheet, "C2", "lalala")
	f.SetCellValue(sheet, "D4", "lalala")
	//f.SetColStyle()	// 当前版本v2.4.1不支持

	f.SaveAs("../assert/BookNilStyle_out.xlsx")
}