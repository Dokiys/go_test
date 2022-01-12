package std

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

// TestIoRead 从文件读数据
func TestIoRead(t *testing.T) {
	{
		file, err := os.Open("../assert/IOReadFrom.txt")
		defer file.Close()
		assert.NoError(t, err)
		b := make([]byte, 1000)

		buf := bufio.NewReader(file)
		_, err = buf.Read(b)
		//count, err := file.Read(b)
		assert.NoError(t, err)
	}

	{
		file, err := os.OpenFile("../assert/IOReadFrom.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		assert.NoError(t, err)
		_, err = file.ReadFrom(bytes.NewReader([]byte("lalalalla")))
		assert.NoError(t, err)
	}
}

// TestIoWrite 向文件写入数据
func TestIoWrite(t *testing.T) {
	{
		assert.NoError(t, ioutil.WriteFile("../assert/IOWriteFile_out.txt", []byte("lalalalla"), 0777))
	}

	{
		file, err := os.Create("../assert/IOCreate_out.txt")
		assert.NoError(t, err)
		_, err = file.Write([]byte("lalalalla"))
		assert.NoError(t, err)
	}
}