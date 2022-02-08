package std

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// TestIO
func TestIO(t *testing.T) {
	{
		//reader
		reader := bytes.NewReader([]byte("1234567890"))

		for {
			b := make([]byte, 6)
			n, err := reader.Read(b)
			if err != nil {
				if err == io.EOF {
					t.Log("Read Finished")
					break
				}
				t.Fatal("Read错误：", err)
			}

			t.Log(string(b[:n]))
		}
	}

	{
		// writer
		//b := []byte{}
		//buf := bytes.NewBuffer(b)
		//_, err := buf.Write([]byte("1234567890a"))
		//assert.NoError(t, err)
		//t.Log(b)
		//t.Log(buf)

		file, err := os.Create("write_test.txt")
		assert.NoError(t, err)
		writer := bufio.NewWriter(file)
		_, err = writer.Write([]byte("123456"))
		assert.NoError(t, err)
		writer.Flush()
	}
}

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