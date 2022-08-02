package std

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

// TestOsMkdir 创建文件夹
func TestOsMkdir(t *testing.T) {
	if err := os.MkdirAll("./a/b/c", os.ModePerm); err != nil {
		t.Error(err)
	}
}

// TestOsTempFile 生成临时文件
func TestOsTempFile(t *testing.T) {
	t.Logf("os.TempDir():\t%s", os.TempDir())

	mkdirTemp, err := os.MkdirTemp("", "HHH")
	assert.NoError(t, err)
	t.Logf("os.MkdirTemp():\t%s", mkdirTemp)

	tempFile, err := ioutil.TempFile("", "HHH*.xlsx")
	assert.NoError(t, err)
	t.Logf("ioutil.TempFile():\t%s", tempFile.Name())
}

// TestOsHostname get hostname
func TestOsHostname(t *testing.T) {
	hostname, err := os.Hostname()
	assert.NoError(t, err)
	t.Log(hostname)
}

// TestOsPid get Pid
func TestOsPid(t *testing.T) {
	t.Log(os.Getpid())
}
