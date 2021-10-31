package inner

import (
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
