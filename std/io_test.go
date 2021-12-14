package std

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

// TestIoReader 测试Reader
func TestIoReader(t *testing.T) {
	file, err := os.Open("../assert/IOReadFrom.txt")
	defer file.Close()
	if err != nil {
		t.Fatal(err)
	}

	b := make([]byte, 1000)
	//count, err := file.Read(b)

	buf := bufio.NewReader(file)
	count, err := buf.Read(b)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("read file size: %d", count)
	t.Logf("read file content: %s", b[:count])
}

// TestIoReadFrom 测试从io.Reader 写入文件
func TestIoReadFrom(t *testing.T) {
	file, err := os.OpenFile("../assert/IOReadFrom.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	assert.NoError(t, err)
	_, err = file.ReadFrom(bytes.NewReader([]byte("lalalalla")))
	assert.NoError(t, err)
}

// TestIoWriteToNewFile 测试ioutil 创建文件并写入数据
func TestIoWriteToNewFile(t *testing.T) {
	assert.NoError(t, ioutil.WriteFile("../assert/IOWriteFile_out.txt",[]byte("lalalalla"),0777))
}

func TestIo(t *testing.T) {
	file, err := os.Create("../assert/IOCreate_out.txt")
	assert.NoError(t, err)
	_, err = file.Write([]byte("lalalalla"))
	assert.NoError(t, err)
}