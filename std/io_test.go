package std

import (
	"bufio"
	"os"
	"testing"
)

// TestIoReader 测试Reader
func TestIoReader(t *testing.T) {
	file, err := os.Open("../assert/io_reader_file.txt")
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

	t.Logf("read file size: %d",count)
	t.Logf("read file content: %s", b[:count])
}
