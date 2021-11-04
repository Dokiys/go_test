package std

import (
	"os"
	"testing"
)

// TestOsWd 获取当前路径
func TestOsWd(t *testing.T) {
	dir, _ := os.Getwd()
	t.Log(dir)
}

// TestOsCreateFile 创建文件
func TestOsCreateFile(t *testing.T) {
	_, err := os.Create("./test.txt")
	if err != nil {
		t.Error(err)
	}
}