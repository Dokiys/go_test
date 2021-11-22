package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

// TestLogrus 测试Logrus
func TestLogrus(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logger := log.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "192.168.32.15",
	})

	logrus.Infof("mytest: %s", "Hello Work!")
	logger.Infof("mytest: %s", "Hello Work!")
}

// TestLogrusOut logrus输出到文件
func TestLogrusOut(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	//log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Error("Failed to log to file, using default stderr")
	}

	log.Info("Hello First File Log")		// 输出到logrus.log文件
}
