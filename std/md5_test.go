package std

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	file, err := os.Open("../asset/IOReadFrom.txt")
	assert.NoError(t, err)

	all, err := ioutil.ReadAll(file)
	assert.NoError(t, err)
	t.Log("1:/n")
	t.Log(all)

	all, err = ioutil.ReadAll(file)
	assert.NoError(t, err)
	t.Log("2:/n")
	t.Log(all)

	md := md5.New()
	_, err = io.Copy(md, file)
	assert.NoError(t, err)
	sum := md.Sum(nil)
	t.Log(hex.EncodeToString(sum))
}

func TestLo(t *testing.T) {
	name := "test_" + fmt.Sprintf("%d_%d", time.Now().UnixNano(), 123)
	t.Log(name)
}
