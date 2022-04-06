package std

import (
	"fmt"
	"testing"
	"time"
)


// TestTimeCheckNil 校验时间初始值
func TestTimeCheckNil(t *testing.T) {
	date, _ := time.ParseInLocation("2006-01-02", "", time.Local)
	t.Log(date)

	if date.IsZero() {
		t.Log(date.Format("2006-01-02 15:04:05"))
	}
}

// TestTimeAdd 添加时间
func TestTimeAdd(t *testing.T) {
	t.Log(time.Now().Format("2006-01-02"))
	t.Log(time.Now().AddDate(1,0,0).Format("2006-01-02"))
}

// TestTimeNano 输出当前时间
func TestTimeNano(t *testing.T) {
	t.Log(time.Now().UnixNano())
}

// TestTimeStrParse 根据字符串和格式转换成时间
func TestTimeStrParse(t *testing.T) {
	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm, _ := time.Parse("01/02/2006", "02/08/2015")

	fmt.Println(tm.Unix())
}
