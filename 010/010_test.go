package _10

import (
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestS(t *testing.T) {
	j := `{"border":[{"type":"top","color":"#FF00FF","style":5},{"type":"bottom","color":"#FF00FF","style":2},{"type":"left","color":"#FF00FF","style":1},{"type":"right","color":"#FF00FF","style":4}],"fill":{"type":"pattern","pattern":1,"color":["#AA00A0"],"shading":0},"font":{"bold":true,"italic":true,"underline":"single","family":"YaHei","size":16,"strike":true,"color":"#0000FF"},"alignment":{"horizontal":"center","indent":0,"justify_last_line":false,"reading_order":0,"relative_indent":0,"shrink_to_fit":false,"text_rotation":0,"vertical":"center","wrap_text":false},"protection":{"hidden":false,"locked":false},"number_format":0,"decimal_places":0,"lang":"","negred":false}`
	f := excelize.NewFile()
	styleIndex, _ := f.NewStyle(j)
	t.Log(styleIndex)
}
