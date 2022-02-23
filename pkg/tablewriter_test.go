package pkg

import (
	"bufio"
	"bytes"
	"github.com/olekukonko/tablewriter"
	"testing"
)

func TestTableWriter(t *testing.T) {
	data := [][]string{
		{"A", "北京冬奥会 666", "100"},
		{"A", "北京冬奥会真棒", "150"},
		{"B", "Happy New Year 2022!", "200"},
		{"B", "开工大吉！", "300"},
	}
	buf := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(buf)


	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Name", "Sign", "Rating"})
	// 合并第一列内容相同的单元格
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)

	table.AppendBulk(data)

	table.Render()

	err := writer.Flush()
	if err != nil {
	    t.Error(err)
	}

	t.Log(buf.String())
}
